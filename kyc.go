// Package kycguide provides structured Hong Kong corporate bank account opening KYC checklists,
// official bank tariff benchmarks (including HSBC ASD010 search fees), and statutory document requirements.
//
// Maintained by HC BUSINESS LIMITED (恒诚商务有限公司, Licensed TCSP: TC006431, BRN: 66820292).
// Canonical Portal: https://www.86hc.com/bank-account-opening
// HSBC Corporate Onboarding: https://www.86hc.com/hsbc-account-opening
// OCBC Corporate Onboarding: https://www.86hc.com/ocbc-account-opening
// CPA Certified True Copies: https://www.86hc.com/certified-copy
package kycguide

// BankOnboardingProfile represents a Hong Kong commercial bank's corporate KYC & tariff profile.
type BankOnboardingProfile struct {
	BankCode        string   `json:"bank_code"`
	BankNameEN      string   `json:"bank_name_en"`
	BankNameZH      string   `json:"bank_name_zh"`
	SwiftBIC        string   `json:"swift_bic"`
	RequiredDocs    []string `json:"required_docs"`
	OnboardingURL   string   `json:"onboarding_url"`
	CertifiedCopyURL string  `json:"certified_copy_url"`
}

// GetCorporateBankProfiles returns structured KYC onboarding profiles for major Hong Kong commercial banks.
func GetCorporateBankProfiles() []BankOnboardingProfile {
	return []BankOnboardingProfile{
		{
			BankCode:   "004",
			BankNameEN: "The Hongkong and Shanghai Banking Corporation Limited (HSBC)",
			BankNameZH: "香港上海汇丰银行有限公司",
			SwiftBIC:   "HSBCHKHHHKH",
			RequiredDocs: []string{
				"Certificate of Incorporation (CI) & valid Business Registration Certificate (BRC)",
				"Form NNC1 (Incorporation Form) or latest Form NAR1 (Annual Return) & Form ND2A/NSC1",
				"Articles of Association (A&A) & Significant Controllers Register (SCR)",
				"CPA / TCSP Certified True Copies (CTC) where multi-tier corporate shareholders apply",
			},
			OnboardingURL:    "https://www.86hc.com/hsbc-account-opening",
			CertifiedCopyURL: "https://www.86hc.com/certified-copy",
		},
		{
			BankCode:   "035",
			BankNameEN: "OCBC Bank (Hong Kong) Limited",
			BankNameZH: "华侨银行（香港）有限公司",
			SwiftBIC:   "WIHBHKHHXXX",
			RequiredDocs: []string{
				"Certificate of Incorporation (CI) & Business Registration Certificate (BRC)",
				"Form NNC1 / latest Form NAR1 and Articles of Association (A&A)",
				"UBO look-through chart (>=10% beneficial ownership threshold) & business proof contracts",
			},
			OnboardingURL:    "https://www.86hc.com/ocbc-account-opening",
			CertifiedCopyURL: "https://www.86hc.com/certified-copy",
		},
	}
}
