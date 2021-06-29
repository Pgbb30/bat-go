package payment

import (
	"context"
	"fmt"

	appctx "github.com/brave-intl/bat-go/utils/context"
)

// List of all the allowed and whitelisted brave skus

const (
	prodUserWalletVote    = "AgEJYnJhdmUuY29tAiNicmF2ZSB1c2VyLXdhbGxldC12b3RlIHNrdSB0b2tlbiB2MQACFHNrdT11c2VyLXdhbGxldC12b3RlAAIKcHJpY2U9MC4yNQACDGN1cnJlbmN5PUJBVAACDGRlc2NyaXB0aW9uPQACGmNyZWRlbnRpYWxfdHlwZT1zaW5nbGUtdXNlAAAGIOaNAUCBMKm0IaLqxefhvxOtAKB0OfoiPn0NPVfI602J"
	prodAnonCardVote      = "AgEJYnJhdmUuY29tAiFicmF2ZSBhbm9uLWNhcmQtdm90ZSBza3UgdG9rZW4gdjEAAhJza3U9YW5vbi1jYXJkLXZvdGUAAgpwcmljZT0wLjI1AAIMY3VycmVuY3k9QkFUAAIMZGVzY3JpcHRpb249AAIaY3JlZGVudGlhbF90eXBlPXNpbmdsZS11c2UAAAYgrMZm85YYwnmjPXcegy5pBM5C+ZLfrySZfYiSe13yp8o="
	prodBraveTogetherPaid = "MDAyMGxvY2F0aW9uIHRvZ2V0aGVyLmJyYXZlLmNvbQowMDMwaWRlbnRpZmllciBicmF2ZS10b2dldGhlci1wYWlkIHNrdSB0b2tlbiB2MQowMDIwY2lkIHNrdT1icmF2ZS10b2dldGhlci1wYWlkCjAwMTBjaWQgcHJpY2U9NQowMDE1Y2lkIGN1cnJlbmN5PVVTRAowMDQzY2lkIGRlc2NyaXB0aW9uPU9uZSBtb250aCBwYWlkIHN1YnNjcmlwdGlvbiBmb3IgQnJhdmUgVG9nZXRoZXIKMDAyNWNpZCBjcmVkZW50aWFsX3R5cGU9dGltZS1saW1pdGVkCjAwMjZjaWQgY3JlZGVudGlhbF92YWxpZF9kdXJhdGlvbj1QMU0KMDAyZnNpZ25hdHVyZSAl/eGfP93lrklACcFClNPvkP3Go0HCtfYVQMs5n/NJpgo="

	stagingUserWalletVote    = "AgEJYnJhdmUuY29tAiNicmF2ZSB1c2VyLXdhbGxldC12b3RlIHNrdSB0b2tlbiB2MQACFHNrdT11c2VyLXdhbGxldC12b3RlAAIKcHJpY2U9MC4yNQACDGN1cnJlbmN5PUJBVAACDGRlc2NyaXB0aW9uPQACGmNyZWRlbnRpYWxfdHlwZT1zaW5nbGUtdXNlAAAGIOH4Li+rduCtFOfV8Lfa2o8h4SQjN5CuIwxmeQFjOk4W"
	stagingAnonCardVote      = "AgEJYnJhdmUuY29tAiFicmF2ZSBhbm9uLWNhcmQtdm90ZSBza3UgdG9rZW4gdjEAAhJza3U9YW5vbi1jYXJkLXZvdGUAAgpwcmljZT0wLjI1AAIMY3VycmVuY3k9QkFUAAIMZGVzY3JpcHRpb249AAIaY3JlZGVudGlhbF90eXBlPXNpbmdsZS11c2UAAAYgPV/WYY5pXhodMPvsilnrLzNH6MA8nFXwyg0qSWX477M="
	stagingBraveTogetherPaid = "MDAyNWxvY2F0aW9uIHRvZ2V0aGVyLmJyYXZlLnNvZnR3YXJlCjAwMzBpZGVudGlmaWVyIGJyYXZlLXRvZ2V0aGVyLXBhaWQgc2t1IHRva2VuIHYxCjAwMjBjaWQgc2t1PWJyYXZlLXRvZ2V0aGVyLXBhaWQKMDAxMGNpZCBwcmljZT01CjAwMTVjaWQgY3VycmVuY3k9VVNECjAwNDNjaWQgZGVzY3JpcHRpb249T25lIG1vbnRoIHBhaWQgc3Vic2NyaXB0aW9uIGZvciBCcmF2ZSBUb2dldGhlcgowMDI1Y2lkIGNyZWRlbnRpYWxfdHlwZT10aW1lLWxpbWl0ZWQKMDAyNmNpZCBjcmVkZW50aWFsX3ZhbGlkX2R1cmF0aW9uPVAxTQowMDJmc2lnbmF0dXJlIBBaYgRlOpoFKqpcnEzOJFKbLzul3DzLEbQbiJCxd9x3Cg=="
	stagingWebtestPJSKUDemo  = "AgEYd2VidGVzdC1wai5oZXJva3VhcHAuY29tAih3ZWJ0ZXN0LXBqLmhlcm9rdWFwcC5jb20gYnJhdmUtdHNoaXJ0IHYxAAIQc2t1PWJyYXZlLXRzaGlydAACCnByaWNlPTAuMjUAAgxjdXJyZW5jeT1CQVQAAgxkZXNjcmlwdGlvbj0AAhpjcmVkZW50aWFsX3R5cGU9c2luZ2xlLXVzZQAABiCcJ0zXGbSg+s3vsClkci44QQQTzWJb9UPyJASMVU11jw=="

	stagingBravePremium   = "MDAyOGxvY2F0aW9uIHByZW1pdW0uYnNnLmJyYXZlLnNvZnR3YXJlCjAwMmFpZGVudGlmaWVyIGJyYXZlLXByZW1pdW0gc2t1IHRva2VuIHYxCjAwMWFjaWQgc2t1PWJyYXZlLXByZW1pdW0KMDAxMWNpZCBwcmljZT0xMAowMDE1Y2lkIGN1cnJlbmN5PVVTRAowMDM1Y2lkIGRlc2NyaXB0aW9uPVByZW1pdW0gYWNjZXNzIHRvIEJyYXZlIHByb2R1Y3RzCjAwMjVjaWQgY3JlZGVudGlhbF90eXBlPXRpbWUtbGltaXRlZAowMDI2Y2lkIGNyZWRlbnRpYWxfdmFsaWRfZHVyYXRpb249UDFNCjAwMzljaWQgc3RyaXBlX3Byb2R1Y3RfaWQ9cHJpY2VfMUlxZ1ZwSG9mMjBicGhHNndjQ0J5clJmCjAwMmZzaWduYXR1cmUgJP7DPkq860UIU/2pkNWCAVJtr13+aEi+qMW2Sw5PMwYK"
	stagingBraveUnlimited = "MDAyYWxvY2F0aW9uIHVubGltaXRlZC5ic2cuYnJhdmUuc29mdHdhcmUKMDAyY2lkZW50aWZpZXIgYnJhdmUtdW5saW1pdGVkIHNrdSB0b2tlbiB2MQowMDFjY2lkIHNrdT1icmF2ZS11bmxpbWl0ZWQKMDAxMWNpZCBwcmljZT0xNQowMDE1Y2lkIGN1cnJlbmN5PVVTRAowMDM3Y2lkIGRlc2NyaXB0aW9uPVVubGltaXRlZCBhY2Nlc3MgdG8gQnJhdmUgcHJvZHVjdHMKMDAyNWNpZCBjcmVkZW50aWFsX3R5cGU9dGltZS1saW1pdGVkCjAwMjZjaWQgY3JlZGVudGlhbF92YWxpZF9kdXJhdGlvbj1QMU0KMDAzOWNpZCBzdHJpcGVfcHJvZHVjdF9pZD1wcmljZV8xSXFnV21Ib2YyMGJwaEc2Q0daanp3N1UKMDAyZnNpZ25hdHVyZSARGoBX5ESxmVJIkCkgFW5yH1WOByJVeOhW+1/Kt30rGAo="

	devUserWalletVote   = "AgEJYnJhdmUuY29tAiNicmF2ZSB1c2VyLXdhbGxldC12b3RlIHNrdSB0b2tlbiB2MQACFHNrdT11c2VyLXdhbGxldC12b3RlAAIKcHJpY2U9MC4yNQACDGN1cnJlbmN5PUJBVAACDGRlc2NyaXB0aW9uPQACGmNyZWRlbnRpYWxfdHlwZT1zaW5nbGUtdXNlAAAGINiB9dUmpqLyeSEdZ23E4dPXwIBOUNJCFN9d5toIME2M"
	devAnonCardVote     = "AgEJYnJhdmUuY29tAiFicmF2ZSBhbm9uLWNhcmQtdm90ZSBza3UgdG9rZW4gdjEAAhJza3U9YW5vbi1jYXJkLXZvdGUAAgpwcmljZT0wLjI1AAIMY3VycmVuY3k9QkFUAAIMZGVzY3JpcHRpb249AAIaY3JlZGVudGlhbF90eXBlPXNpbmdsZS11c2UAAAYgPpv+Al9jRgVCaR49/AoRrsjQqXGqkwaNfqVka00SJxQ="
	devSearchClosedBeta = "AgEVc2VhcmNoLmJyYXZlLnNvZnR3YXJlAh9zZWFyY2ggY2xvc2VkIGJldGEgcHJvZ3JhbSBkZW1vAAIWc2t1PXNlYXJjaC1iZXRhLWFjY2VzcwACB3ByaWNlPTAAAgxjdXJyZW5jeT1CQVQAAi1kZXNjcmlwdGlvbj1TZWFyY2ggY2xvc2VkIGJldGEgcHJvZ3JhbSBhY2Nlc3MAAhpjcmVkZW50aWFsX3R5cGU9c2luZ2xlLXVzZQAABiB3uXfAAkNSRQd24jSauRny3VM0BYZ8yOclPTEgPa0xrA=="
	devBraveTalkPremium = "MDAyMWxvY2F0aW9uIHRhbGsuYnJhdmUuc29mdHdhcmUKMDAyZmlkZW50aWZpZXIgYnJhdmUtdGFsay1wcmVtaXVtIHNrdSB0b2tlbiB2MQowMDFmY2lkIHNrdT1icmF2ZS10YWxrLXByZW1pdW0KMDAxM2NpZCBwcmljZT03LjAwCjAwMTVjaWQgY3VycmVuY3k9VVNECjAwMzFjaWQgZGVzY3JpcHRpb249UHJlbWl1bSBhY2Nlc3MgdG8gQnJhdmUgVGFsawowMDIzY2lkIGNyZWRlbnRpYWxfdHlwZT1zaW5nbGUtdXNlCjAxMTNjaWQgbWV0YWRhdGE9IHsgInN0cmlwZV9wcm9kdWN0X2lkIjogInByb2RfSmxIWjlSWjk4bDNOMmsiLCAic3RyaXBlX2l0ZW1faWQiOiAicHJpY2VfMUo3bDBDSG9mMjBicGhHNmlYR01Tb1RqIiwgInN0cmlwZV9zdWNjZXNzX3VyaSI6ICJodHRwczovL2FjY291bnQuYnJhdmUuc29mdHdhcmUvYWNjb3VudD9pbnRlbnQ9cHJvdmlzaW9uIiwgInN0cmlwZV9jYW5jZWxfdXJpIjogImh0dHBzOi8vYWNjb3VudC5icmF2ZS5zb2Z0d2FyZS9wbGFucz9pbnRlbnQ9Y2hlY2tvdXQiIH0KMDAyZnNpZ25hdHVyZSDOHnlW2JYs2mSrKATxn0+XaLpjdip/lwXkoOMlP40DmAo="
)

var skuMap = map[string]map[string]bool{
	"production": {
		prodUserWalletVote:    true,
		prodAnonCardVote:      true,
		prodBraveTogetherPaid: true,
	},
	"staging": {
		stagingBraveUnlimited:    true,
		stagingBravePremium:      true,
		stagingUserWalletVote:    true,
		stagingAnonCardVote:      true,
		stagingBraveTogetherPaid: true,
		stagingWebtestPJSKUDemo:  true,
	},
	"development": {
		devUserWalletVote:   true,
		devAnonCardVote:     true,
		devSearchClosedBeta: true,
		devBraveTalkPremium: true,
	},
}

// temporary, until we can validate macaroon signatures
func validateHardcodedSku(ctx context.Context, sku string) (bool, error) {
	// check sku white list from environment
	whitelistSKUs, ok := ctx.Value(appctx.WhitelistSKUsCTXKey).([]string)
	if ok {
		for _, whitelistSKU := range whitelistSKUs {
			if sku == whitelistSKU {
				return true, nil
			}
		}
	}

	// check hardcoded based on environment (non whitelisted)
	env, err := appctx.GetStringFromContext(ctx, appctx.EnvironmentCTXKey)
	if err != nil {
		return false, fmt.Errorf("failed to get environment: %w", err)
	}
	valid, ok := skuMap[env][sku]
	return valid && ok, nil
}
