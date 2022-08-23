package promotion

// Code generated by gowrap. DO NOT EDIT.
// template: ../../.prom-gowrap.tmpl
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/brave-intl/bat-go/services/promotion -i ReadOnlyDatastore -t ../../.prom-gowrap.tmpl -o instrumented_read_only_datastore.go -l ""

import (
	"context"
	"time"

	walletutils "github.com/brave-intl/bat-go/libs/wallet"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// ReadOnlyDatastoreWithPrometheus implements ReadOnlyDatastore interface with all methods wrapped
// with Prometheus metrics
type ReadOnlyDatastoreWithPrometheus struct {
	base         ReadOnlyDatastore
	instanceName string
}

var readonlydatastoreDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "promotion_readonly_datastore_duration_seconds",
		Help:       "readonlydatastore runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewReadOnlyDatastoreWithPrometheus returns an instance of the ReadOnlyDatastore decorated with prometheus summary metric
func NewReadOnlyDatastoreWithPrometheus(base ReadOnlyDatastore, instanceName string) ReadOnlyDatastoreWithPrometheus {
	return ReadOnlyDatastoreWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// GetAvailablePromotions implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetAvailablePromotions(platform string) (pa1 []Promotion, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetAvailablePromotions", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetAvailablePromotions(platform)
}

// GetAvailablePromotionsForWallet implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetAvailablePromotionsForWallet(wallet *walletutils.Info, platform string) (pa1 []Promotion, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetAvailablePromotionsForWallet", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetAvailablePromotionsForWallet(wallet, platform)
}

// GetClaimByWalletAndPromotion implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetClaimByWalletAndPromotion(wallet *walletutils.Info, promotionID *Promotion) (cp1 *Claim, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetClaimByWalletAndPromotion", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetClaimByWalletAndPromotion(wallet, promotionID)
}

// GetClaimCreds implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetClaimCreds(claimID uuid.UUID) (cp1 *ClaimCreds, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetClaimCreds", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetClaimCreds(claimID)
}

// GetClaimSummary implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetClaimSummary(walletID uuid.UUID, grantType string) (cp1 *ClaimSummary, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetClaimSummary", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetClaimSummary(walletID, grantType)
}

// GetCustodianDrainInfo implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetCustodianDrainInfo(paymentID *uuid.UUID) (ca1 []CustodianDrain, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetCustodianDrainInfo", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetCustodianDrainInfo(paymentID)
}

// GetDrainPoll implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetDrainPoll(drainID *uuid.UUID) (dp1 *DrainPoll, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetDrainPoll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetDrainPoll(drainID)
}

// GetDrainsByBatchID implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetDrainsByBatchID(ctx context.Context, batchID *uuid.UUID) (da1 []DrainTransfer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetDrainsByBatchID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetDrainsByBatchID(ctx, batchID)
}

// GetIssuer implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetIssuer(promotionID uuid.UUID, cohort string) (ip1 *Issuer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetIssuer", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetIssuer(promotionID, cohort)
}

// GetIssuerByPublicKey implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetIssuerByPublicKey(publicKey string) (ip1 *Issuer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetIssuerByPublicKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetIssuerByPublicKey(publicKey)
}

// GetPreClaim implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetPreClaim(promotionID uuid.UUID, walletID string) (cp1 *Claim, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetPreClaim", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetPreClaim(promotionID, walletID)
}

// GetPromotion implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetPromotion(promotionID uuid.UUID) (pp1 *Promotion, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetPromotion", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetPromotion(promotionID)
}

// GetPromotionsMissingIssuer implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetPromotionsMissingIssuer(limit int) (ua1 []uuid.UUID, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetPromotionsMissingIssuer", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetPromotionsMissingIssuer(limit)
}

// GetWithdrawalsAssociated implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) GetWithdrawalsAssociated(walletID *uuid.UUID, claimID *uuid.UUID) (up1 *uuid.UUID, d1 decimal.Decimal, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetWithdrawalsAssociated", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetWithdrawalsAssociated(walletID, claimID)
}

// Migrate implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) Migrate(p1 ...uint) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "Migrate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Migrate(p1...)
}

// NewMigrate implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) NewMigrate() (mp1 *migrate.Migrate, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "NewMigrate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.NewMigrate()
}

// RawDB implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) RawDB() (dp1 *sqlx.DB) {
	_since := time.Now()
	defer func() {
		result := "ok"
		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RawDB", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawDB()
}

// RollbackTx implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) RollbackTx(tx *sqlx.Tx) {
	_since := time.Now()
	defer func() {
		result := "ok"
		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RollbackTx", result).Observe(time.Since(_since).Seconds())
	}()
	_d.base.RollbackTx(tx)
	return
}

// RollbackTxAndHandle implements ReadOnlyDatastore
func (_d ReadOnlyDatastoreWithPrometheus) RollbackTxAndHandle(tx *sqlx.Tx) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		readonlydatastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RollbackTxAndHandle", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RollbackTxAndHandle(tx)
}