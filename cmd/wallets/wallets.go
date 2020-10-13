package wallets

import (

	// pprof imports
	_ "net/http/pprof"

	"github.com/brave-intl/bat-go/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	walletsCmd = &cobra.Command{
		Use:   "wallet",
		Short: "provides wallets micro-service entrypoint",
	}

	walletsRestCmd = &cobra.Command{
		Use:   "rest",
		Short: "provides REST api services",
		Run:   WalletRestRun,
	}
	db                  string
	walletsFeatureFlag  bool
	enableLinkDrainFlag bool
	roDB                string
)

func init() {
	// add grpc and rest commands
	walletsCmd.AddCommand(walletsRestCmd)

	// add this command as a serve subcommand
	cmd.ServeCmd.AddCommand(walletsCmd)

	// setup the flags
	// datastore - the writable datastore
	walletsCmd.PersistentFlags().StringVarP(&db, "datastore", "", "",
		"the datastore for the wallet system")
	cmd.Must(viper.BindPFlag("datastore", walletsCmd.PersistentFlags().Lookup("datastore")))
	cmd.Must(viper.BindEnv("datastore", "DATABASE_URL"))

	// walletsFeatureFlag - enable the wallet endpoints through this feature flag
	walletsCmd.PersistentFlags().BoolVarP(&walletsFeatureFlag, "wallets-feature-flag", "", false,
		"the feature flag enabling the wallets feature")
	cmd.Must(viper.BindPFlag("wallets-feature-flag", walletsCmd.PersistentFlags().Lookup("wallets-feature-flag")))
	cmd.Must(viper.BindEnv("wallets-feature-flag", "FEATURE_WALLET"))

	// ENABLE_LINKING_DRAINING - enable ability to link wallets and drain wallets
	walletsCmd.PersistentFlags().BoolVarP(&enableLinkDrainFlag, "enable-link-drain-flag", "", false,
		"the in-migration flag disabling the wallets link feature")
	cmd.Must(viper.BindPFlag("enable-link-drain-flag", walletsCmd.PersistentFlags().Lookup("enable-link-drain-flag")))
	cmd.Must(viper.BindEnv("enable-link-drain-flag", "ENABLE_LINKING_DRAINING"))

	// ro-datastore - the writable datastore
	walletsCmd.PersistentFlags().StringVarP(&roDB, "ro-datastore", "", "",
		"the read only datastore for the wallet system")
	cmd.Must(viper.BindPFlag("ro-datastore", walletsCmd.PersistentFlags().Lookup("ro-datastore")))
	cmd.Must(viper.BindEnv("ro-datastore", "RO_DATABASE_URL"))
}