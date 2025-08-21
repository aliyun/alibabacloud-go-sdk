// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafSiteSettings interface {
	dara.Model
	String() string
	GoString() string
	SetAddBotProtectionHeaders(v *WafSiteSettingsAddBotProtectionHeaders) *WafSiteSettings
	GetAddBotProtectionHeaders() *WafSiteSettingsAddBotProtectionHeaders
	SetAddSecurityHeaders(v *WafSiteSettingsAddSecurityHeaders) *WafSiteSettings
	GetAddSecurityHeaders() *WafSiteSettingsAddSecurityHeaders
	SetBandwidthAbuseProtection(v *WafSiteSettingsBandwidthAbuseProtection) *WafSiteSettings
	GetBandwidthAbuseProtection() *WafSiteSettingsBandwidthAbuseProtection
	SetBotManagement(v *WafSiteSettingsBotManagement) *WafSiteSettings
	GetBotManagement() *WafSiteSettingsBotManagement
	SetClientIpIdentifier(v *WafSiteSettingsClientIpIdentifier) *WafSiteSettings
	GetClientIpIdentifier() *WafSiteSettingsClientIpIdentifier
	SetDisableSecurityModule(v *WafSiteSettingsDisableSecurityModule) *WafSiteSettings
	GetDisableSecurityModule() *WafSiteSettingsDisableSecurityModule
	SetSecurityLevel(v *WafSiteSettingsSecurityLevel) *WafSiteSettings
	GetSecurityLevel() *WafSiteSettingsSecurityLevel
}

type WafSiteSettings struct {
	AddBotProtectionHeaders  *WafSiteSettingsAddBotProtectionHeaders  `json:"AddBotProtectionHeaders,omitempty" xml:"AddBotProtectionHeaders,omitempty" type:"Struct"`
	AddSecurityHeaders       *WafSiteSettingsAddSecurityHeaders       `json:"AddSecurityHeaders,omitempty" xml:"AddSecurityHeaders,omitempty" type:"Struct"`
	BandwidthAbuseProtection *WafSiteSettingsBandwidthAbuseProtection `json:"BandwidthAbuseProtection,omitempty" xml:"BandwidthAbuseProtection,omitempty" type:"Struct"`
	BotManagement            *WafSiteSettingsBotManagement            `json:"BotManagement,omitempty" xml:"BotManagement,omitempty" type:"Struct"`
	ClientIpIdentifier       *WafSiteSettingsClientIpIdentifier       `json:"ClientIpIdentifier,omitempty" xml:"ClientIpIdentifier,omitempty" type:"Struct"`
	DisableSecurityModule    *WafSiteSettingsDisableSecurityModule    `json:"DisableSecurityModule,omitempty" xml:"DisableSecurityModule,omitempty" type:"Struct"`
	SecurityLevel            *WafSiteSettingsSecurityLevel            `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty" type:"Struct"`
}

func (s WafSiteSettings) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettings) GoString() string {
	return s.String()
}

func (s *WafSiteSettings) GetAddBotProtectionHeaders() *WafSiteSettingsAddBotProtectionHeaders {
	return s.AddBotProtectionHeaders
}

func (s *WafSiteSettings) GetAddSecurityHeaders() *WafSiteSettingsAddSecurityHeaders {
	return s.AddSecurityHeaders
}

func (s *WafSiteSettings) GetBandwidthAbuseProtection() *WafSiteSettingsBandwidthAbuseProtection {
	return s.BandwidthAbuseProtection
}

func (s *WafSiteSettings) GetBotManagement() *WafSiteSettingsBotManagement {
	return s.BotManagement
}

func (s *WafSiteSettings) GetClientIpIdentifier() *WafSiteSettingsClientIpIdentifier {
	return s.ClientIpIdentifier
}

func (s *WafSiteSettings) GetDisableSecurityModule() *WafSiteSettingsDisableSecurityModule {
	return s.DisableSecurityModule
}

func (s *WafSiteSettings) GetSecurityLevel() *WafSiteSettingsSecurityLevel {
	return s.SecurityLevel
}

func (s *WafSiteSettings) SetAddBotProtectionHeaders(v *WafSiteSettingsAddBotProtectionHeaders) *WafSiteSettings {
	s.AddBotProtectionHeaders = v
	return s
}

func (s *WafSiteSettings) SetAddSecurityHeaders(v *WafSiteSettingsAddSecurityHeaders) *WafSiteSettings {
	s.AddSecurityHeaders = v
	return s
}

func (s *WafSiteSettings) SetBandwidthAbuseProtection(v *WafSiteSettingsBandwidthAbuseProtection) *WafSiteSettings {
	s.BandwidthAbuseProtection = v
	return s
}

func (s *WafSiteSettings) SetBotManagement(v *WafSiteSettingsBotManagement) *WafSiteSettings {
	s.BotManagement = v
	return s
}

func (s *WafSiteSettings) SetClientIpIdentifier(v *WafSiteSettingsClientIpIdentifier) *WafSiteSettings {
	s.ClientIpIdentifier = v
	return s
}

func (s *WafSiteSettings) SetDisableSecurityModule(v *WafSiteSettingsDisableSecurityModule) *WafSiteSettings {
	s.DisableSecurityModule = v
	return s
}

func (s *WafSiteSettings) SetSecurityLevel(v *WafSiteSettingsSecurityLevel) *WafSiteSettings {
	s.SecurityLevel = v
	return s
}

func (s *WafSiteSettings) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsAddBotProtectionHeaders struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsAddBotProtectionHeaders) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsAddBotProtectionHeaders) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsAddBotProtectionHeaders) GetEnable() *bool {
	return s.Enable
}

func (s *WafSiteSettingsAddBotProtectionHeaders) SetEnable(v bool) *WafSiteSettingsAddBotProtectionHeaders {
	s.Enable = &v
	return s
}

func (s *WafSiteSettingsAddBotProtectionHeaders) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsAddSecurityHeaders struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsAddSecurityHeaders) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsAddSecurityHeaders) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsAddSecurityHeaders) GetEnable() *bool {
	return s.Enable
}

func (s *WafSiteSettingsAddSecurityHeaders) SetEnable(v bool) *WafSiteSettingsAddSecurityHeaders {
	s.Enable = &v
	return s
}

func (s *WafSiteSettingsAddSecurityHeaders) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBandwidthAbuseProtection struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WafSiteSettingsBandwidthAbuseProtection) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBandwidthAbuseProtection) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBandwidthAbuseProtection) GetAction() *string {
	return s.Action
}

func (s *WafSiteSettingsBandwidthAbuseProtection) GetId() *int64 {
	return s.Id
}

func (s *WafSiteSettingsBandwidthAbuseProtection) GetStatus() *string {
	return s.Status
}

func (s *WafSiteSettingsBandwidthAbuseProtection) SetAction(v string) *WafSiteSettingsBandwidthAbuseProtection {
	s.Action = &v
	return s
}

func (s *WafSiteSettingsBandwidthAbuseProtection) SetId(v int64) *WafSiteSettingsBandwidthAbuseProtection {
	s.Id = &v
	return s
}

func (s *WafSiteSettingsBandwidthAbuseProtection) SetStatus(v string) *WafSiteSettingsBandwidthAbuseProtection {
	s.Status = &v
	return s
}

func (s *WafSiteSettingsBandwidthAbuseProtection) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagement struct {
	DefiniteBots   *WafSiteSettingsBotManagementDefiniteBots   `json:"DefiniteBots,omitempty" xml:"DefiniteBots,omitempty" type:"Struct"`
	EffectOnStatic *WafSiteSettingsBotManagementEffectOnStatic `json:"EffectOnStatic,omitempty" xml:"EffectOnStatic,omitempty" type:"Struct"`
	JSDetection    *WafSiteSettingsBotManagementJSDetection    `json:"JSDetection,omitempty" xml:"JSDetection,omitempty" type:"Struct"`
	LikelyBots     *WafSiteSettingsBotManagementLikelyBots     `json:"LikelyBots,omitempty" xml:"LikelyBots,omitempty" type:"Struct"`
	VerifiedBots   *WafSiteSettingsBotManagementVerifiedBots   `json:"VerifiedBots,omitempty" xml:"VerifiedBots,omitempty" type:"Struct"`
}

func (s WafSiteSettingsBotManagement) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagement) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagement) GetDefiniteBots() *WafSiteSettingsBotManagementDefiniteBots {
	return s.DefiniteBots
}

func (s *WafSiteSettingsBotManagement) GetEffectOnStatic() *WafSiteSettingsBotManagementEffectOnStatic {
	return s.EffectOnStatic
}

func (s *WafSiteSettingsBotManagement) GetJSDetection() *WafSiteSettingsBotManagementJSDetection {
	return s.JSDetection
}

func (s *WafSiteSettingsBotManagement) GetLikelyBots() *WafSiteSettingsBotManagementLikelyBots {
	return s.LikelyBots
}

func (s *WafSiteSettingsBotManagement) GetVerifiedBots() *WafSiteSettingsBotManagementVerifiedBots {
	return s.VerifiedBots
}

func (s *WafSiteSettingsBotManagement) SetDefiniteBots(v *WafSiteSettingsBotManagementDefiniteBots) *WafSiteSettingsBotManagement {
	s.DefiniteBots = v
	return s
}

func (s *WafSiteSettingsBotManagement) SetEffectOnStatic(v *WafSiteSettingsBotManagementEffectOnStatic) *WafSiteSettingsBotManagement {
	s.EffectOnStatic = v
	return s
}

func (s *WafSiteSettingsBotManagement) SetJSDetection(v *WafSiteSettingsBotManagementJSDetection) *WafSiteSettingsBotManagement {
	s.JSDetection = v
	return s
}

func (s *WafSiteSettingsBotManagement) SetLikelyBots(v *WafSiteSettingsBotManagementLikelyBots) *WafSiteSettingsBotManagement {
	s.LikelyBots = v
	return s
}

func (s *WafSiteSettingsBotManagement) SetVerifiedBots(v *WafSiteSettingsBotManagementVerifiedBots) *WafSiteSettingsBotManagement {
	s.VerifiedBots = v
	return s
}

func (s *WafSiteSettingsBotManagement) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagementDefiniteBots struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafSiteSettingsBotManagementDefiniteBots) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagementDefiniteBots) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagementDefiniteBots) GetAction() *string {
	return s.Action
}

func (s *WafSiteSettingsBotManagementDefiniteBots) GetId() *int64 {
	return s.Id
}

func (s *WafSiteSettingsBotManagementDefiniteBots) SetAction(v string) *WafSiteSettingsBotManagementDefiniteBots {
	s.Action = &v
	return s
}

func (s *WafSiteSettingsBotManagementDefiniteBots) SetId(v int64) *WafSiteSettingsBotManagementDefiniteBots {
	s.Id = &v
	return s
}

func (s *WafSiteSettingsBotManagementDefiniteBots) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagementEffectOnStatic struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsBotManagementEffectOnStatic) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagementEffectOnStatic) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagementEffectOnStatic) GetEnable() *bool {
	return s.Enable
}

func (s *WafSiteSettingsBotManagementEffectOnStatic) SetEnable(v bool) *WafSiteSettingsBotManagementEffectOnStatic {
	s.Enable = &v
	return s
}

func (s *WafSiteSettingsBotManagementEffectOnStatic) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagementJSDetection struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsBotManagementJSDetection) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagementJSDetection) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagementJSDetection) GetEnable() *bool {
	return s.Enable
}

func (s *WafSiteSettingsBotManagementJSDetection) SetEnable(v bool) *WafSiteSettingsBotManagementJSDetection {
	s.Enable = &v
	return s
}

func (s *WafSiteSettingsBotManagementJSDetection) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagementLikelyBots struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafSiteSettingsBotManagementLikelyBots) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagementLikelyBots) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagementLikelyBots) GetAction() *string {
	return s.Action
}

func (s *WafSiteSettingsBotManagementLikelyBots) GetId() *int64 {
	return s.Id
}

func (s *WafSiteSettingsBotManagementLikelyBots) SetAction(v string) *WafSiteSettingsBotManagementLikelyBots {
	s.Action = &v
	return s
}

func (s *WafSiteSettingsBotManagementLikelyBots) SetId(v int64) *WafSiteSettingsBotManagementLikelyBots {
	s.Id = &v
	return s
}

func (s *WafSiteSettingsBotManagementLikelyBots) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsBotManagementVerifiedBots struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafSiteSettingsBotManagementVerifiedBots) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsBotManagementVerifiedBots) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsBotManagementVerifiedBots) GetAction() *string {
	return s.Action
}

func (s *WafSiteSettingsBotManagementVerifiedBots) GetId() *int64 {
	return s.Id
}

func (s *WafSiteSettingsBotManagementVerifiedBots) SetAction(v string) *WafSiteSettingsBotManagementVerifiedBots {
	s.Action = &v
	return s
}

func (s *WafSiteSettingsBotManagementVerifiedBots) SetId(v int64) *WafSiteSettingsBotManagementVerifiedBots {
	s.Id = &v
	return s
}

func (s *WafSiteSettingsBotManagementVerifiedBots) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsClientIpIdentifier struct {
	Headers []*string `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Mode    *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s WafSiteSettingsClientIpIdentifier) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsClientIpIdentifier) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsClientIpIdentifier) GetHeaders() []*string {
	return s.Headers
}

func (s *WafSiteSettingsClientIpIdentifier) GetMode() *string {
	return s.Mode
}

func (s *WafSiteSettingsClientIpIdentifier) SetHeaders(v []*string) *WafSiteSettingsClientIpIdentifier {
	s.Headers = v
	return s
}

func (s *WafSiteSettingsClientIpIdentifier) SetMode(v string) *WafSiteSettingsClientIpIdentifier {
	s.Mode = &v
	return s
}

func (s *WafSiteSettingsClientIpIdentifier) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsDisableSecurityModule struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WafSiteSettingsDisableSecurityModule) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsDisableSecurityModule) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsDisableSecurityModule) GetStatus() *string {
	return s.Status
}

func (s *WafSiteSettingsDisableSecurityModule) SetStatus(v string) *WafSiteSettingsDisableSecurityModule {
	s.Status = &v
	return s
}

func (s *WafSiteSettingsDisableSecurityModule) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsSecurityLevel struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafSiteSettingsSecurityLevel) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsSecurityLevel) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsSecurityLevel) GetValue() *string {
	return s.Value
}

func (s *WafSiteSettingsSecurityLevel) SetValue(v string) *WafSiteSettingsSecurityLevel {
	s.Value = &v
	return s
}

func (s *WafSiteSettingsSecurityLevel) Validate() error {
	return dara.Validate(s)
}
