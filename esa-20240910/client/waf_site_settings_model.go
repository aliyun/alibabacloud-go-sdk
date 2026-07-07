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
	SetRequestBodyInspection(v *WafSiteSettingsRequestBodyInspection) *WafSiteSettings
	GetRequestBodyInspection() *WafSiteSettingsRequestBodyInspection
	SetSecurityLevel(v *WafSiteSettingsSecurityLevel) *WafSiteSettings
	GetSecurityLevel() *WafSiteSettingsSecurityLevel
}

type WafSiteSettings struct {
	// The bot protection headers.
	AddBotProtectionHeaders *WafSiteSettingsAddBotProtectionHeaders `json:"AddBotProtectionHeaders,omitempty" xml:"AddBotProtectionHeaders,omitempty" type:"Struct"`
	// The security headers.
	AddSecurityHeaders *WafSiteSettingsAddSecurityHeaders `json:"AddSecurityHeaders,omitempty" xml:"AddSecurityHeaders,omitempty" type:"Struct"`
	// The bandwidth abuse protection.
	BandwidthAbuseProtection *WafSiteSettingsBandwidthAbuseProtection `json:"BandwidthAbuseProtection,omitempty" xml:"BandwidthAbuseProtection,omitempty" type:"Struct"`
	// The bot management.
	BotManagement *WafSiteSettingsBotManagement `json:"BotManagement,omitempty" xml:"BotManagement,omitempty" type:"Struct"`
	// The client IP identification.
	ClientIpIdentifier *WafSiteSettingsClientIpIdentifier `json:"ClientIpIdentifier,omitempty" xml:"ClientIpIdentifier,omitempty" type:"Struct"`
	// The configuration for disabling the security module.
	DisableSecurityModule *WafSiteSettingsDisableSecurityModule `json:"DisableSecurityModule,omitempty" xml:"DisableSecurityModule,omitempty" type:"Struct"`
	// The request body inspection configuration. Controls the deep packet inspection behavior of WAF for HTTP request bodies. After this feature is enabled, content-based matching rules such as SQL injection and XSS detection take effect on request bodies.
	//
	// This structure can contain the following fields:
	//
	// - Id: The unique identifier of the built-in inspection rule.
	//
	// - SizeLimit: The maximum size of the request body to inspect.
	//
	// - Action: The action to take when the request body exceeds the size limit.
	RequestBodyInspection *WafSiteSettingsRequestBodyInspection `json:"RequestBodyInspection,omitempty" xml:"RequestBodyInspection,omitempty" type:"Struct"`
	// The security level.
	SecurityLevel *WafSiteSettingsSecurityLevel `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty" type:"Struct"`
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

func (s *WafSiteSettings) GetRequestBodyInspection() *WafSiteSettingsRequestBodyInspection {
	return s.RequestBodyInspection
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

func (s *WafSiteSettings) SetRequestBodyInspection(v *WafSiteSettingsRequestBodyInspection) *WafSiteSettings {
	s.RequestBodyInspection = v
	return s
}

func (s *WafSiteSettings) SetSecurityLevel(v *WafSiteSettingsSecurityLevel) *WafSiteSettings {
	s.SecurityLevel = v
	return s
}

func (s *WafSiteSettings) Validate() error {
	if s.AddBotProtectionHeaders != nil {
		if err := s.AddBotProtectionHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.AddSecurityHeaders != nil {
		if err := s.AddSecurityHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.BandwidthAbuseProtection != nil {
		if err := s.BandwidthAbuseProtection.Validate(); err != nil {
			return err
		}
	}
	if s.BotManagement != nil {
		if err := s.BotManagement.Validate(); err != nil {
			return err
		}
	}
	if s.ClientIpIdentifier != nil {
		if err := s.ClientIpIdentifier.Validate(); err != nil {
			return err
		}
	}
	if s.DisableSecurityModule != nil {
		if err := s.DisableSecurityModule.Validate(); err != nil {
			return err
		}
	}
	if s.RequestBodyInspection != nil {
		if err := s.RequestBodyInspection.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityLevel != nil {
		if err := s.SecurityLevel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafSiteSettingsAddBotProtectionHeaders struct {
	// The switch.
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
	// The switch.
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
	// The action of the bandwidth abuse protection rule.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ID of the bandwidth abuse protection rule.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the bandwidth abuse protection rule.
	//
	// example:
	//
	// on
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
	// The definite bots.
	DefiniteBots *WafSiteSettingsBotManagementDefiniteBots `json:"DefiniteBots,omitempty" xml:"DefiniteBots,omitempty" type:"Struct"`
	// Specifies whether the rule applies to static resource requests.
	EffectOnStatic *WafSiteSettingsBotManagementEffectOnStatic `json:"EffectOnStatic,omitempty" xml:"EffectOnStatic,omitempty" type:"Struct"`
	// The JavaScript detection.
	JSDetection *WafSiteSettingsBotManagementJSDetection `json:"JSDetection,omitempty" xml:"JSDetection,omitempty" type:"Struct"`
	// The likely bots.
	LikelyBots *WafSiteSettingsBotManagementLikelyBots `json:"LikelyBots,omitempty" xml:"LikelyBots,omitempty" type:"Struct"`
	// The verified bots.
	VerifiedBots *WafSiteSettingsBotManagementVerifiedBots `json:"VerifiedBots,omitempty" xml:"VerifiedBots,omitempty" type:"Struct"`
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
	if s.DefiniteBots != nil {
		if err := s.DefiniteBots.Validate(); err != nil {
			return err
		}
	}
	if s.EffectOnStatic != nil {
		if err := s.EffectOnStatic.Validate(); err != nil {
			return err
		}
	}
	if s.JSDetection != nil {
		if err := s.JSDetection.Validate(); err != nil {
			return err
		}
	}
	if s.LikelyBots != nil {
		if err := s.LikelyBots.Validate(); err != nil {
			return err
		}
	}
	if s.VerifiedBots != nil {
		if err := s.VerifiedBots.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafSiteSettingsBotManagementDefiniteBots struct {
	// The action.
	//
	// example:
	//
	// captcha
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The switch.
	//
	// example:
	//
	// true
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
	// The switch.
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
	// The action.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 20000002
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The action.
	//
	// example:
	//
	// bypass
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 20000003
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The specified headers.
	Headers []*string `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The identification mode.
	//
	// example:
	//
	// headers
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
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
	// The status switch for disabling the security module.
	//
	// example:
	//
	// on
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

type WafSiteSettingsRequestBodyInspection struct {
	// The action to take when the request body size exceeds SizeLimit.
	//
	// Common valid values (the complete list is determined by the server-side configuration):
	//
	// - allow: allows the request without performing deep packet inspection on the portion that exceeds the limit.
	//
	// > The complete enumeration is determined by the WAF server-side configuration.
	//
	// example:
	//
	// allow
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The request body inspection rule ID, which is the unique identifier of the built-in rule. When request body inspection is enabled, the server uses this ID to associate the matching logic of the built-in inspection rule. The valid values are based on the built-in rule list of WAF.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum size of the request body to inspect, in bytes.
	//
	// - If the request body is less than or equal to this value, the entire content is subject to WAF matching.
	//
	// - If the request body exceeds this value, the action specified in the Action field is taken, such as inspecting only the first N bytes, rejecting the request, or allowing the request.
	//
	// > The valid value range and default value are determined by the WAF server-side configuration.
	//
	// example:
	//
	// 16KB
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s WafSiteSettingsRequestBodyInspection) String() string {
	return dara.Prettify(s)
}

func (s WafSiteSettingsRequestBodyInspection) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsRequestBodyInspection) GetAction() *string {
	return s.Action
}

func (s *WafSiteSettingsRequestBodyInspection) GetId() *int64 {
	return s.Id
}

func (s *WafSiteSettingsRequestBodyInspection) GetSizeLimit() *string {
	return s.SizeLimit
}

func (s *WafSiteSettingsRequestBodyInspection) SetAction(v string) *WafSiteSettingsRequestBodyInspection {
	s.Action = &v
	return s
}

func (s *WafSiteSettingsRequestBodyInspection) SetId(v int64) *WafSiteSettingsRequestBodyInspection {
	s.Id = &v
	return s
}

func (s *WafSiteSettingsRequestBodyInspection) SetSizeLimit(v string) *WafSiteSettingsRequestBodyInspection {
	s.SizeLimit = &v
	return s
}

func (s *WafSiteSettingsRequestBodyInspection) Validate() error {
	return dara.Validate(s)
}

type WafSiteSettingsSecurityLevel struct {
	// The security level value.
	//
	// example:
	//
	// low
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
