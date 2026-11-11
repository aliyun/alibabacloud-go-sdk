// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafRuleConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *WafRuleConfig
	GetAction() *string
	SetActions(v *WafRuleConfigActions) *WafRuleConfig
	GetActions() *WafRuleConfigActions
	SetAppPackage(v *WafRuleConfigAppPackage) *WafRuleConfig
	GetAppPackage() *WafRuleConfigAppPackage
	SetAppSdk(v *WafRuleConfigAppSdk) *WafRuleConfig
	GetAppSdk() *WafRuleConfigAppSdk
	SetExpression(v string) *WafRuleConfig
	GetExpression() *string
	SetId(v int64) *WafRuleConfig
	GetId() *int64
	SetManagedGroupId(v int64) *WafRuleConfig
	GetManagedGroupId() *int64
	SetManagedList(v string) *WafRuleConfig
	GetManagedList() *string
	SetManagedRulesets(v []*WafRuleConfigManagedRulesets) *WafRuleConfig
	GetManagedRulesets() []*WafRuleConfigManagedRulesets
	SetName(v string) *WafRuleConfig
	GetName() *string
	SetNotes(v string) *WafRuleConfig
	GetNotes() *string
	SetRateLimit(v *WafRuleConfigRateLimit) *WafRuleConfig
	GetRateLimit() *WafRuleConfigRateLimit
	SetSecurityLevel(v *WafRuleConfigSecurityLevel) *WafRuleConfig
	GetSecurityLevel() *WafRuleConfigSecurityLevel
	SetSigchl(v []*string) *WafRuleConfig
	GetSigchl() []*string
	SetStatus(v string) *WafRuleConfig
	GetStatus() *string
	SetTimer(v *WafTimer) *WafRuleConfig
	GetTimer() *WafTimer
	SetType(v string) *WafRuleConfig
	GetType() *string
	SetValue(v string) *WafRuleConfig
	GetValue() *string
}

type WafRuleConfig struct {
	// The action performed after a WAF rule is matched.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The extended actions.
	Actions *WafRuleConfigActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	// The repackaging detection configuration.
	AppPackage *WafRuleConfigAppPackage `json:"AppPackage,omitempty" xml:"AppPackage,omitempty" type:"Struct"`
	// The app SDK configuration.
	AppSdk *WafRuleConfigAppSdk `json:"AppSdk,omitempty" xml:"AppSdk,omitempty" type:"Struct"`
	// The matching expression.
	//
	// example:
	//
	// ip.src eq 10.0.0.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The ID of the WAF rule.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// **[Deprecated]*	- The managed rule group ID.
	//
	// example:
	//
	// 30000001
	ManagedGroupId *int64 `json:"ManagedGroupId,omitempty" xml:"ManagedGroupId,omitempty"`
	// The managed list.
	//
	// example:
	//
	// intelligence_crawler
	ManagedList *string `json:"ManagedList,omitempty" xml:"ManagedList,omitempty"`
	// The list of managed rulesets.
	ManagedRulesets []*WafRuleConfigManagedRulesets `json:"ManagedRulesets,omitempty" xml:"ManagedRulesets,omitempty" type:"Repeated"`
	// The name of the WAF rule.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks.
	//
	// example:
	//
	// example notes
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// The configurations related to the rate limiting rule.
	RateLimit *WafRuleConfigRateLimit `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	// The security level.
	SecurityLevel *WafRuleConfigSecurityLevel `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty" type:"Struct"`
	// The token verification configurations.
	Sigchl []*string `json:"Sigchl,omitempty" xml:"Sigchl,omitempty" type:"Repeated"`
	// The status of the WAF rule.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The timer.
	Timer *WafTimer `json:"Timer,omitempty" xml:"Timer,omitempty"`
	// The type of the WAF rule.
	//
	// example:
	//
	// http_custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP access control value.
	//
	// example:
	//
	// 10.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfig) GoString() string {
	return s.String()
}

func (s *WafRuleConfig) GetAction() *string {
	return s.Action
}

func (s *WafRuleConfig) GetActions() *WafRuleConfigActions {
	return s.Actions
}

func (s *WafRuleConfig) GetAppPackage() *WafRuleConfigAppPackage {
	return s.AppPackage
}

func (s *WafRuleConfig) GetAppSdk() *WafRuleConfigAppSdk {
	return s.AppSdk
}

func (s *WafRuleConfig) GetExpression() *string {
	return s.Expression
}

func (s *WafRuleConfig) GetId() *int64 {
	return s.Id
}

func (s *WafRuleConfig) GetManagedGroupId() *int64 {
	return s.ManagedGroupId
}

func (s *WafRuleConfig) GetManagedList() *string {
	return s.ManagedList
}

func (s *WafRuleConfig) GetManagedRulesets() []*WafRuleConfigManagedRulesets {
	return s.ManagedRulesets
}

func (s *WafRuleConfig) GetName() *string {
	return s.Name
}

func (s *WafRuleConfig) GetNotes() *string {
	return s.Notes
}

func (s *WafRuleConfig) GetRateLimit() *WafRuleConfigRateLimit {
	return s.RateLimit
}

func (s *WafRuleConfig) GetSecurityLevel() *WafRuleConfigSecurityLevel {
	return s.SecurityLevel
}

func (s *WafRuleConfig) GetSigchl() []*string {
	return s.Sigchl
}

func (s *WafRuleConfig) GetStatus() *string {
	return s.Status
}

func (s *WafRuleConfig) GetTimer() *WafTimer {
	return s.Timer
}

func (s *WafRuleConfig) GetType() *string {
	return s.Type
}

func (s *WafRuleConfig) GetValue() *string {
	return s.Value
}

func (s *WafRuleConfig) SetAction(v string) *WafRuleConfig {
	s.Action = &v
	return s
}

func (s *WafRuleConfig) SetActions(v *WafRuleConfigActions) *WafRuleConfig {
	s.Actions = v
	return s
}

func (s *WafRuleConfig) SetAppPackage(v *WafRuleConfigAppPackage) *WafRuleConfig {
	s.AppPackage = v
	return s
}

func (s *WafRuleConfig) SetAppSdk(v *WafRuleConfigAppSdk) *WafRuleConfig {
	s.AppSdk = v
	return s
}

func (s *WafRuleConfig) SetExpression(v string) *WafRuleConfig {
	s.Expression = &v
	return s
}

func (s *WafRuleConfig) SetId(v int64) *WafRuleConfig {
	s.Id = &v
	return s
}

func (s *WafRuleConfig) SetManagedGroupId(v int64) *WafRuleConfig {
	s.ManagedGroupId = &v
	return s
}

func (s *WafRuleConfig) SetManagedList(v string) *WafRuleConfig {
	s.ManagedList = &v
	return s
}

func (s *WafRuleConfig) SetManagedRulesets(v []*WafRuleConfigManagedRulesets) *WafRuleConfig {
	s.ManagedRulesets = v
	return s
}

func (s *WafRuleConfig) SetName(v string) *WafRuleConfig {
	s.Name = &v
	return s
}

func (s *WafRuleConfig) SetNotes(v string) *WafRuleConfig {
	s.Notes = &v
	return s
}

func (s *WafRuleConfig) SetRateLimit(v *WafRuleConfigRateLimit) *WafRuleConfig {
	s.RateLimit = v
	return s
}

func (s *WafRuleConfig) SetSecurityLevel(v *WafRuleConfigSecurityLevel) *WafRuleConfig {
	s.SecurityLevel = v
	return s
}

func (s *WafRuleConfig) SetSigchl(v []*string) *WafRuleConfig {
	s.Sigchl = v
	return s
}

func (s *WafRuleConfig) SetStatus(v string) *WafRuleConfig {
	s.Status = &v
	return s
}

func (s *WafRuleConfig) SetTimer(v *WafTimer) *WafRuleConfig {
	s.Timer = v
	return s
}

func (s *WafRuleConfig) SetType(v string) *WafRuleConfig {
	s.Type = &v
	return s
}

func (s *WafRuleConfig) SetValue(v string) *WafRuleConfig {
	s.Value = &v
	return s
}

func (s *WafRuleConfig) Validate() error {
	if s.Actions != nil {
		if err := s.Actions.Validate(); err != nil {
			return err
		}
	}
	if s.AppPackage != nil {
		if err := s.AppPackage.Validate(); err != nil {
			return err
		}
	}
	if s.AppSdk != nil {
		if err := s.AppSdk.Validate(); err != nil {
			return err
		}
	}
	if s.ManagedRulesets != nil {
		for _, item := range s.ManagedRulesets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RateLimit != nil {
		if err := s.RateLimit.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityLevel != nil {
		if err := s.SecurityLevel.Validate(); err != nil {
			return err
		}
	}
	if s.Timer != nil {
		if err := s.Timer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafRuleConfigActions struct {
	// The modules bypassed by the whitelist rule.
	Bypass *WafRuleConfigActionsBypass `json:"Bypass,omitempty" xml:"Bypass,omitempty" type:"Struct"`
	// The custom response page.
	Response *WafRuleConfigActionsResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s WafRuleConfigActions) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigActions) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActions) GetBypass() *WafRuleConfigActionsBypass {
	return s.Bypass
}

func (s *WafRuleConfigActions) GetResponse() *WafRuleConfigActionsResponse {
	return s.Response
}

func (s *WafRuleConfigActions) SetBypass(v *WafRuleConfigActionsBypass) *WafRuleConfigActions {
	s.Bypass = v
	return s
}

func (s *WafRuleConfigActions) SetResponse(v *WafRuleConfigActionsResponse) *WafRuleConfigActions {
	s.Response = v
	return s
}

func (s *WafRuleConfigActions) Validate() error {
	if s.Bypass != nil {
		if err := s.Bypass.Validate(); err != nil {
			return err
		}
	}
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafRuleConfigActionsBypass struct {
	// The list of custom rule IDs.
	CustomRules []*int64 `json:"CustomRules,omitempty" xml:"CustomRules,omitempty" type:"Repeated"`
	// The list of managed rule IDs.
	RegularRules []*int64 `json:"RegularRules,omitempty" xml:"RegularRules,omitempty" type:"Repeated"`
	// The list of managed rule types.
	RegularTypes []*string `json:"RegularTypes,omitempty" xml:"RegularTypes,omitempty" type:"Repeated"`
	// The type of the bypassed module.
	//
	// example:
	//
	// part
	Skip *string `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The list of bypassed module phases.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s WafRuleConfigActionsBypass) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigActionsBypass) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActionsBypass) GetCustomRules() []*int64 {
	return s.CustomRules
}

func (s *WafRuleConfigActionsBypass) GetRegularRules() []*int64 {
	return s.RegularRules
}

func (s *WafRuleConfigActionsBypass) GetRegularTypes() []*string {
	return s.RegularTypes
}

func (s *WafRuleConfigActionsBypass) GetSkip() *string {
	return s.Skip
}

func (s *WafRuleConfigActionsBypass) GetTags() []*string {
	return s.Tags
}

func (s *WafRuleConfigActionsBypass) SetCustomRules(v []*int64) *WafRuleConfigActionsBypass {
	s.CustomRules = v
	return s
}

func (s *WafRuleConfigActionsBypass) SetRegularRules(v []*int64) *WafRuleConfigActionsBypass {
	s.RegularRules = v
	return s
}

func (s *WafRuleConfigActionsBypass) SetRegularTypes(v []*string) *WafRuleConfigActionsBypass {
	s.RegularTypes = v
	return s
}

func (s *WafRuleConfigActionsBypass) SetSkip(v string) *WafRuleConfigActionsBypass {
	s.Skip = &v
	return s
}

func (s *WafRuleConfigActionsBypass) SetTags(v []*string) *WafRuleConfigActionsBypass {
	s.Tags = v
	return s
}

func (s *WafRuleConfigActionsBypass) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigActionsResponse struct {
	// The custom response code.
	//
	// example:
	//
	// 403
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the custom response page. You can obtain the ID by calling the [ListPages](https://help.aliyun.com/document_detail/2850223.html) operation.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafRuleConfigActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigActionsResponse) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActionsResponse) GetCode() *int32 {
	return s.Code
}

func (s *WafRuleConfigActionsResponse) GetId() *int64 {
	return s.Id
}

func (s *WafRuleConfigActionsResponse) SetCode(v int32) *WafRuleConfigActionsResponse {
	s.Code = &v
	return s
}

func (s *WafRuleConfigActionsResponse) SetId(v int64) *WafRuleConfigActionsResponse {
	s.Id = &v
	return s
}

func (s *WafRuleConfigActionsResponse) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigAppPackage struct {
	// The repackaging detection.
	PackageSigns []*WafRuleConfigAppPackagePackageSigns `json:"PackageSigns,omitempty" xml:"PackageSigns,omitempty" type:"Repeated"`
}

func (s WafRuleConfigAppPackage) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigAppPackage) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppPackage) GetPackageSigns() []*WafRuleConfigAppPackagePackageSigns {
	return s.PackageSigns
}

func (s *WafRuleConfigAppPackage) SetPackageSigns(v []*WafRuleConfigAppPackagePackageSigns) *WafRuleConfigAppPackage {
	s.PackageSigns = v
	return s
}

func (s *WafRuleConfigAppPackage) Validate() error {
	if s.PackageSigns != nil {
		for _, item := range s.PackageSigns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafRuleConfigAppPackagePackageSigns struct {
	// The specified valid package name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The package signature.
	//
	// example:
	//
	// sign
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s WafRuleConfigAppPackagePackageSigns) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigAppPackagePackageSigns) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppPackagePackageSigns) GetName() *string {
	return s.Name
}

func (s *WafRuleConfigAppPackagePackageSigns) GetSign() *string {
	return s.Sign
}

func (s *WafRuleConfigAppPackagePackageSigns) SetName(v string) *WafRuleConfigAppPackagePackageSigns {
	s.Name = &v
	return s
}

func (s *WafRuleConfigAppPackagePackageSigns) SetSign(v string) *WafRuleConfigAppPackagePackageSigns {
	s.Sign = &v
	return s
}

func (s *WafRuleConfigAppPackagePackageSigns) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigAppSdk struct {
	// The custom signature field.
	CustomSign *WafRuleConfigAppSdkCustomSign `json:"CustomSign,omitempty" xml:"CustomSign,omitempty" type:"Struct"`
	// The custom signature field switch.
	//
	// example:
	//
	// on
	CustomSignStatus *string `json:"CustomSignStatus,omitempty" xml:"CustomSignStatus,omitempty"`
	// The app feature anomaly.
	FeatureAbnormal []*string `json:"FeatureAbnormal,omitempty" xml:"FeatureAbnormal,omitempty" type:"Repeated"`
}

func (s WafRuleConfigAppSdk) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigAppSdk) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppSdk) GetCustomSign() *WafRuleConfigAppSdkCustomSign {
	return s.CustomSign
}

func (s *WafRuleConfigAppSdk) GetCustomSignStatus() *string {
	return s.CustomSignStatus
}

func (s *WafRuleConfigAppSdk) GetFeatureAbnormal() []*string {
	return s.FeatureAbnormal
}

func (s *WafRuleConfigAppSdk) SetCustomSign(v *WafRuleConfigAppSdkCustomSign) *WafRuleConfigAppSdk {
	s.CustomSign = v
	return s
}

func (s *WafRuleConfigAppSdk) SetCustomSignStatus(v string) *WafRuleConfigAppSdk {
	s.CustomSignStatus = &v
	return s
}

func (s *WafRuleConfigAppSdk) SetFeatureAbnormal(v []*string) *WafRuleConfigAppSdk {
	s.FeatureAbnormal = v
	return s
}

func (s *WafRuleConfigAppSdk) Validate() error {
	if s.CustomSign != nil {
		if err := s.CustomSign.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafRuleConfigAppSdkCustomSign struct {
	// The field name.
	//
	// example:
	//
	// sign
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The field value.
	//
	// example:
	//
	// examplesignvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafRuleConfigAppSdkCustomSign) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigAppSdkCustomSign) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppSdkCustomSign) GetKey() *string {
	return s.Key
}

func (s *WafRuleConfigAppSdkCustomSign) GetValue() *string {
	return s.Value
}

func (s *WafRuleConfigAppSdkCustomSign) SetKey(v string) *WafRuleConfigAppSdkCustomSign {
	s.Key = &v
	return s
}

func (s *WafRuleConfigAppSdkCustomSign) SetValue(v string) *WafRuleConfigAppSdkCustomSign {
	s.Value = &v
	return s
}

func (s *WafRuleConfigAppSdkCustomSign) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigManagedRulesets struct {
	// The action to perform.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The attack type.
	//
	// example:
	//
	// 11
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The list of managed rules.
	ManagedRules []*WafRuleConfigManagedRulesetsManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	// The number of enabled rules.
	//
	// example:
	//
	// 50
	NumberEnabled *int32 `json:"NumberEnabled,omitempty" xml:"NumberEnabled,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 100
	NumberTotal *int32 `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
	// The protection level.
	//
	// example:
	//
	// 4
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
}

func (s WafRuleConfigManagedRulesets) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigManagedRulesets) GoString() string {
	return s.String()
}

func (s *WafRuleConfigManagedRulesets) GetAction() *string {
	return s.Action
}

func (s *WafRuleConfigManagedRulesets) GetAttackType() *int32 {
	return s.AttackType
}

func (s *WafRuleConfigManagedRulesets) GetManagedRules() []*WafRuleConfigManagedRulesetsManagedRules {
	return s.ManagedRules
}

func (s *WafRuleConfigManagedRulesets) GetNumberEnabled() *int32 {
	return s.NumberEnabled
}

func (s *WafRuleConfigManagedRulesets) GetNumberTotal() *int32 {
	return s.NumberTotal
}

func (s *WafRuleConfigManagedRulesets) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *WafRuleConfigManagedRulesets) SetAction(v string) *WafRuleConfigManagedRulesets {
	s.Action = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetAttackType(v int32) *WafRuleConfigManagedRulesets {
	s.AttackType = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetManagedRules(v []*WafRuleConfigManagedRulesetsManagedRules) *WafRuleConfigManagedRulesets {
	s.ManagedRules = v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetNumberEnabled(v int32) *WafRuleConfigManagedRulesets {
	s.NumberEnabled = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetNumberTotal(v int32) *WafRuleConfigManagedRulesets {
	s.NumberTotal = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetProtectionLevel(v int32) *WafRuleConfigManagedRulesets {
	s.ProtectionLevel = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) Validate() error {
	if s.ManagedRules != nil {
		for _, item := range s.ManagedRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafRuleConfigManagedRulesetsManagedRules struct {
	// The action of the managed rule.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The managed rule ID.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the managed rule.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WafRuleConfigManagedRulesetsManagedRules) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigManagedRulesetsManagedRules) GoString() string {
	return s.String()
}

func (s *WafRuleConfigManagedRulesetsManagedRules) GetAction() *string {
	return s.Action
}

func (s *WafRuleConfigManagedRulesetsManagedRules) GetId() *int64 {
	return s.Id
}

func (s *WafRuleConfigManagedRulesetsManagedRules) GetStatus() *string {
	return s.Status
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetAction(v string) *WafRuleConfigManagedRulesetsManagedRules {
	s.Action = &v
	return s
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetId(v int64) *WafRuleConfigManagedRulesetsManagedRules {
	s.Id = &v
	return s
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetStatus(v string) *WafRuleConfigManagedRulesetsManagedRules {
	s.Status = &v
	return s
}

func (s *WafRuleConfigManagedRulesetsManagedRules) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigRateLimit struct {
	// The specified object for rate statistics.
	Characteristics *WafRatelimitCharacteristics `json:"Characteristics,omitempty" xml:"Characteristics,omitempty"`
	// The statistical period.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Specifies whether to apply the rule to requests that hit the cache.
	//
	// example:
	//
	// true
	OnHit *bool `json:"OnHit,omitempty" xml:"OnHit,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 10
	TTL *int32 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The statistical threshold.
	Threshold *WafRuleConfigRateLimitThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s WafRuleConfigRateLimit) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigRateLimit) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimit) GetCharacteristics() *WafRatelimitCharacteristics {
	return s.Characteristics
}

func (s *WafRuleConfigRateLimit) GetInterval() *int32 {
	return s.Interval
}

func (s *WafRuleConfigRateLimit) GetOnHit() *bool {
	return s.OnHit
}

func (s *WafRuleConfigRateLimit) GetTTL() *int32 {
	return s.TTL
}

func (s *WafRuleConfigRateLimit) GetThreshold() *WafRuleConfigRateLimitThreshold {
	return s.Threshold
}

func (s *WafRuleConfigRateLimit) SetCharacteristics(v *WafRatelimitCharacteristics) *WafRuleConfigRateLimit {
	s.Characteristics = v
	return s
}

func (s *WafRuleConfigRateLimit) SetInterval(v int32) *WafRuleConfigRateLimit {
	s.Interval = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetOnHit(v bool) *WafRuleConfigRateLimit {
	s.OnHit = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetTTL(v int32) *WafRuleConfigRateLimit {
	s.TTL = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetThreshold(v *WafRuleConfigRateLimitThreshold) *WafRuleConfigRateLimit {
	s.Threshold = v
	return s
}

func (s *WafRuleConfigRateLimit) Validate() error {
	if s.Characteristics != nil {
		if err := s.Characteristics.Validate(); err != nil {
			return err
		}
	}
	if s.Threshold != nil {
		if err := s.Threshold.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafRuleConfigRateLimitThreshold struct {
	// The threshold for the number of distinct managed rules.
	//
	// example:
	//
	// 10
	DistinctManagedRules *int32 `json:"DistinctManagedRules,omitempty" xml:"DistinctManagedRules,omitempty"`
	// The managed rule hit threshold.
	//
	// example:
	//
	// 10
	ManagedRulesBlocked *int32 `json:"ManagedRulesBlocked,omitempty" xml:"ManagedRulesBlocked,omitempty"`
	// The request threshold.
	//
	// example:
	//
	// 10
	Request *int32 `json:"Request,omitempty" xml:"Request,omitempty"`
	// The response code threshold.
	ResponseStatus *WafRuleConfigRateLimitThresholdResponseStatus `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty" type:"Struct"`
	// **[Deprecated]*	- The traffic threshold.
	//
	// example:
	//
	// 10Gb
	Traffic *string `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s WafRuleConfigRateLimitThreshold) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigRateLimitThreshold) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimitThreshold) GetDistinctManagedRules() *int32 {
	return s.DistinctManagedRules
}

func (s *WafRuleConfigRateLimitThreshold) GetManagedRulesBlocked() *int32 {
	return s.ManagedRulesBlocked
}

func (s *WafRuleConfigRateLimitThreshold) GetRequest() *int32 {
	return s.Request
}

func (s *WafRuleConfigRateLimitThreshold) GetResponseStatus() *WafRuleConfigRateLimitThresholdResponseStatus {
	return s.ResponseStatus
}

func (s *WafRuleConfigRateLimitThreshold) GetTraffic() *string {
	return s.Traffic
}

func (s *WafRuleConfigRateLimitThreshold) SetDistinctManagedRules(v int32) *WafRuleConfigRateLimitThreshold {
	s.DistinctManagedRules = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetManagedRulesBlocked(v int32) *WafRuleConfigRateLimitThreshold {
	s.ManagedRulesBlocked = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetRequest(v int32) *WafRuleConfigRateLimitThreshold {
	s.Request = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetResponseStatus(v *WafRuleConfigRateLimitThresholdResponseStatus) *WafRuleConfigRateLimitThreshold {
	s.ResponseStatus = v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetTraffic(v string) *WafRuleConfigRateLimitThreshold {
	s.Traffic = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) Validate() error {
	if s.ResponseStatus != nil {
		if err := s.ResponseStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafRuleConfigRateLimitThresholdResponseStatus struct {
	// The HTTP response code.
	//
	// example:
	//
	// 404
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code count threshold.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The response code percentage.
	//
	// example:
	//
	// 10
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s WafRuleConfigRateLimitThresholdResponseStatus) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigRateLimitThresholdResponseStatus) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) GetCode() *int32 {
	return s.Code
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) GetCount() *int32 {
	return s.Count
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) GetRatio() *int32 {
	return s.Ratio
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetCode(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Code = &v
	return s
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetCount(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Count = &v
	return s
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetRatio(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Ratio = &v
	return s
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) Validate() error {
	return dara.Validate(s)
}

type WafRuleConfigSecurityLevel struct {
	// The security level value. Valid values:
	//
	// 	- off: Disabled.
	//
	// 	- essentially_off: Essentially disabled.
	//
	// 	- low: Low.
	//
	// 	- medium: Medium.
	//
	// 	- high: High.
	//
	// 	- under_attack: Under attack.
	//
	// example:
	//
	// low
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafRuleConfigSecurityLevel) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigSecurityLevel) GoString() string {
	return s.String()
}

func (s *WafRuleConfigSecurityLevel) GetValue() *string {
	return s.Value
}

func (s *WafRuleConfigSecurityLevel) SetValue(v string) *WafRuleConfigSecurityLevel {
	s.Value = &v
	return s
}

func (s *WafRuleConfigSecurityLevel) Validate() error {
	return dara.Validate(s)
}
