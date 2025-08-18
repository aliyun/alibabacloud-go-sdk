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
	Action          *string                         `json:"Action,omitempty" xml:"Action,omitempty"`
	Actions         *WafRuleConfigActions           `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	AppPackage      *WafRuleConfigAppPackage        `json:"AppPackage,omitempty" xml:"AppPackage,omitempty" type:"Struct"`
	AppSdk          *WafRuleConfigAppSdk            `json:"AppSdk,omitempty" xml:"AppSdk,omitempty" type:"Struct"`
	Expression      *string                         `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Id              *int64                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ManagedGroupId  *int64                          `json:"ManagedGroupId,omitempty" xml:"ManagedGroupId,omitempty"`
	ManagedList     *string                         `json:"ManagedList,omitempty" xml:"ManagedList,omitempty"`
	ManagedRulesets []*WafRuleConfigManagedRulesets `json:"ManagedRulesets,omitempty" xml:"ManagedRulesets,omitempty" type:"Repeated"`
	Name            *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Notes           *string                         `json:"Notes,omitempty" xml:"Notes,omitempty"`
	RateLimit       *WafRuleConfigRateLimit         `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	SecurityLevel   *WafRuleConfigSecurityLevel     `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty" type:"Struct"`
	Sigchl          []*string                       `json:"Sigchl,omitempty" xml:"Sigchl,omitempty" type:"Repeated"`
	Status          *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Timer           *WafTimer                       `json:"Timer,omitempty" xml:"Timer,omitempty"`
	Type            *string                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Value           *string                         `json:"Value,omitempty" xml:"Value,omitempty"`
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
	return dara.Validate(s)
}

type WafRuleConfigActions struct {
	Bypass   *WafRuleConfigActionsBypass   `json:"Bypass,omitempty" xml:"Bypass,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type WafRuleConfigActionsBypass struct {
	CustomRules  []*int64  `json:"CustomRules,omitempty" xml:"CustomRules,omitempty" type:"Repeated"`
	RegularRules []*int64  `json:"RegularRules,omitempty" xml:"RegularRules,omitempty" type:"Repeated"`
	RegularTypes []*string `json:"RegularTypes,omitempty" xml:"RegularTypes,omitempty" type:"Repeated"`
	Skip         *string   `json:"Skip,omitempty" xml:"Skip,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	Id   *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	return dara.Validate(s)
}

type WafRuleConfigAppPackagePackageSigns struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	CustomSign       *WafRuleConfigAppSdkCustomSign `json:"CustomSign,omitempty" xml:"CustomSign,omitempty" type:"Struct"`
	CustomSignStatus *string                        `json:"CustomSignStatus,omitempty" xml:"CustomSignStatus,omitempty"`
	FeatureAbnormal  []*string                      `json:"FeatureAbnormal,omitempty" xml:"FeatureAbnormal,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type WafRuleConfigAppSdkCustomSign struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Action          *string                                     `json:"Action,omitempty" xml:"Action,omitempty"`
	AttackType      *int32                                      `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	ManagedRules    []*WafRuleConfigManagedRulesetsManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	NumberEnabled   *int32                                      `json:"NumberEnabled,omitempty" xml:"NumberEnabled,omitempty"`
	NumberTotal     *int32                                      `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
	ProtectionLevel *int32                                      `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
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
	return dara.Validate(s)
}

type WafRuleConfigManagedRulesetsManagedRules struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
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
	Characteristics *WafRuleMatch2                   `json:"Characteristics,omitempty" xml:"Characteristics,omitempty"`
	Interval        *int32                           `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OnHit           *bool                            `json:"OnHit,omitempty" xml:"OnHit,omitempty"`
	TTL             *int32                           `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Threshold       *WafRuleConfigRateLimitThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s WafRuleConfigRateLimit) String() string {
	return dara.Prettify(s)
}

func (s WafRuleConfigRateLimit) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimit) GetCharacteristics() *WafRuleMatch2 {
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

func (s *WafRuleConfigRateLimit) SetCharacteristics(v *WafRuleMatch2) *WafRuleConfigRateLimit {
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
	return dara.Validate(s)
}

type WafRuleConfigRateLimitThreshold struct {
	DistinctManagedRules *int32                                         `json:"DistinctManagedRules,omitempty" xml:"DistinctManagedRules,omitempty"`
	ManagedRulesBlocked  *int32                                         `json:"ManagedRulesBlocked,omitempty" xml:"ManagedRulesBlocked,omitempty"`
	Request              *int32                                         `json:"Request,omitempty" xml:"Request,omitempty"`
	ResponseStatus       *WafRuleConfigRateLimitThresholdResponseStatus `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty" type:"Struct"`
	Traffic              *string                                        `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
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
	return dara.Validate(s)
}

type WafRuleConfigRateLimitThresholdResponseStatus struct {
	Code  *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
