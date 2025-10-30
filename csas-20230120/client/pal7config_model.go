// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPAL7Config interface {
	dara.Model
	String() string
	GoString() string
	SetBypassConfig(v *PAL7ConfigBypassConfig) *PAL7Config
	GetBypassConfig() *PAL7ConfigBypassConfig
	SetCertId(v string) *PAL7Config
	GetCertId() *string
	SetDnsConfig(v *PAL7ConfigDnsConfig) *PAL7Config
	GetDnsConfig() *PAL7ConfigDnsConfig
	SetJsHookConfig(v *PAL7ConfigJsHookConfig) *PAL7Config
	GetJsHookConfig() *PAL7ConfigJsHookConfig
	SetProxyDomainTypes(v []byte) *PAL7Config
	GetProxyDomainTypes() []byte
	SetRequestHeaderRewriteConfig(v *PAL7ConfigRequestHeaderRewriteConfig) *PAL7Config
	GetRequestHeaderRewriteConfig() *PAL7ConfigRequestHeaderRewriteConfig
	SetRequestQueryRewriteConfig(v *PAL7ConfigRequestQueryRewriteConfig) *PAL7Config
	GetRequestQueryRewriteConfig() *PAL7ConfigRequestQueryRewriteConfig
	SetResponseHeaderRewriteConfig(v *PAL7ConfigResponseHeaderRewriteConfig) *PAL7Config
	GetResponseHeaderRewriteConfig() *PAL7ConfigResponseHeaderRewriteConfig
	SetResponseRewriteConfig(v *PAL7ConfigResponseRewriteConfig) *PAL7Config
	GetResponseRewriteConfig() *PAL7ConfigResponseRewriteConfig
}

type PAL7Config struct {
	BypassConfig *PAL7ConfigBypassConfig `json:"BypassConfig,omitempty" xml:"BypassConfig,omitempty" type:"Struct"`
	CertId       *string                 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DnsConfig    *PAL7ConfigDnsConfig    `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	JsHookConfig *PAL7ConfigJsHookConfig `json:"JsHookConfig,omitempty" xml:"JsHookConfig,omitempty" type:"Struct"`
	// example:
	//
	// 逗号分隔的枚举值：automatic,custom
	ProxyDomainTypes            []byte                                 `json:"ProxyDomainTypes,omitempty" xml:"ProxyDomainTypes,omitempty"`
	RequestHeaderRewriteConfig  *PAL7ConfigRequestHeaderRewriteConfig  `json:"RequestHeaderRewriteConfig,omitempty" xml:"RequestHeaderRewriteConfig,omitempty" type:"Struct"`
	RequestQueryRewriteConfig   *PAL7ConfigRequestQueryRewriteConfig   `json:"RequestQueryRewriteConfig,omitempty" xml:"RequestQueryRewriteConfig,omitempty" type:"Struct"`
	ResponseHeaderRewriteConfig *PAL7ConfigResponseHeaderRewriteConfig `json:"ResponseHeaderRewriteConfig,omitempty" xml:"ResponseHeaderRewriteConfig,omitempty" type:"Struct"`
	ResponseRewriteConfig       *PAL7ConfigResponseRewriteConfig       `json:"ResponseRewriteConfig,omitempty" xml:"ResponseRewriteConfig,omitempty" type:"Struct"`
}

func (s PAL7Config) String() string {
	return dara.Prettify(s)
}

func (s PAL7Config) GoString() string {
	return s.String()
}

func (s *PAL7Config) GetBypassConfig() *PAL7ConfigBypassConfig {
	return s.BypassConfig
}

func (s *PAL7Config) GetCertId() *string {
	return s.CertId
}

func (s *PAL7Config) GetDnsConfig() *PAL7ConfigDnsConfig {
	return s.DnsConfig
}

func (s *PAL7Config) GetJsHookConfig() *PAL7ConfigJsHookConfig {
	return s.JsHookConfig
}

func (s *PAL7Config) GetProxyDomainTypes() []byte {
	return s.ProxyDomainTypes
}

func (s *PAL7Config) GetRequestHeaderRewriteConfig() *PAL7ConfigRequestHeaderRewriteConfig {
	return s.RequestHeaderRewriteConfig
}

func (s *PAL7Config) GetRequestQueryRewriteConfig() *PAL7ConfigRequestQueryRewriteConfig {
	return s.RequestQueryRewriteConfig
}

func (s *PAL7Config) GetResponseHeaderRewriteConfig() *PAL7ConfigResponseHeaderRewriteConfig {
	return s.ResponseHeaderRewriteConfig
}

func (s *PAL7Config) GetResponseRewriteConfig() *PAL7ConfigResponseRewriteConfig {
	return s.ResponseRewriteConfig
}

func (s *PAL7Config) SetBypassConfig(v *PAL7ConfigBypassConfig) *PAL7Config {
	s.BypassConfig = v
	return s
}

func (s *PAL7Config) SetCertId(v string) *PAL7Config {
	s.CertId = &v
	return s
}

func (s *PAL7Config) SetDnsConfig(v *PAL7ConfigDnsConfig) *PAL7Config {
	s.DnsConfig = v
	return s
}

func (s *PAL7Config) SetJsHookConfig(v *PAL7ConfigJsHookConfig) *PAL7Config {
	s.JsHookConfig = v
	return s
}

func (s *PAL7Config) SetProxyDomainTypes(v []byte) *PAL7Config {
	s.ProxyDomainTypes = v
	return s
}

func (s *PAL7Config) SetRequestHeaderRewriteConfig(v *PAL7ConfigRequestHeaderRewriteConfig) *PAL7Config {
	s.RequestHeaderRewriteConfig = v
	return s
}

func (s *PAL7Config) SetRequestQueryRewriteConfig(v *PAL7ConfigRequestQueryRewriteConfig) *PAL7Config {
	s.RequestQueryRewriteConfig = v
	return s
}

func (s *PAL7Config) SetResponseHeaderRewriteConfig(v *PAL7ConfigResponseHeaderRewriteConfig) *PAL7Config {
	s.ResponseHeaderRewriteConfig = v
	return s
}

func (s *PAL7Config) SetResponseRewriteConfig(v *PAL7ConfigResponseRewriteConfig) *PAL7Config {
	s.ResponseRewriteConfig = v
	return s
}

func (s *PAL7Config) Validate() error {
	if s.BypassConfig != nil {
		if err := s.BypassConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DnsConfig != nil {
		if err := s.DnsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JsHookConfig != nil {
		if err := s.JsHookConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RequestHeaderRewriteConfig != nil {
		if err := s.RequestHeaderRewriteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RequestQueryRewriteConfig != nil {
		if err := s.RequestQueryRewriteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseHeaderRewriteConfig != nil {
		if err := s.ResponseHeaderRewriteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseRewriteConfig != nil {
		if err := s.ResponseRewriteConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PAL7ConfigBypassConfig struct {
	AppBypassFroms []*string                               `json:"AppBypassFroms,omitempty" xml:"AppBypassFroms,omitempty" type:"Repeated"`
	Mode           *string                                 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UrlBypassRules []*PAL7ConfigBypassConfigUrlBypassRules `json:"UrlBypassRules,omitempty" xml:"UrlBypassRules,omitempty" type:"Repeated"`
}

func (s PAL7ConfigBypassConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigBypassConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigBypassConfig) GetAppBypassFroms() []*string {
	return s.AppBypassFroms
}

func (s *PAL7ConfigBypassConfig) GetMode() *string {
	return s.Mode
}

func (s *PAL7ConfigBypassConfig) GetUrlBypassRules() []*PAL7ConfigBypassConfigUrlBypassRules {
	return s.UrlBypassRules
}

func (s *PAL7ConfigBypassConfig) SetAppBypassFroms(v []*string) *PAL7ConfigBypassConfig {
	s.AppBypassFroms = v
	return s
}

func (s *PAL7ConfigBypassConfig) SetMode(v string) *PAL7ConfigBypassConfig {
	s.Mode = &v
	return s
}

func (s *PAL7ConfigBypassConfig) SetUrlBypassRules(v []*PAL7ConfigBypassConfigUrlBypassRules) *PAL7ConfigBypassConfig {
	s.UrlBypassRules = v
	return s
}

func (s *PAL7ConfigBypassConfig) Validate() error {
	if s.UrlBypassRules != nil {
		for _, item := range s.UrlBypassRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PAL7ConfigBypassConfigUrlBypassRules struct {
	Froms []*string `json:"Froms,omitempty" xml:"Froms,omitempty" type:"Repeated"`
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s PAL7ConfigBypassConfigUrlBypassRules) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigBypassConfigUrlBypassRules) GoString() string {
	return s.String()
}

func (s *PAL7ConfigBypassConfigUrlBypassRules) GetFroms() []*string {
	return s.Froms
}

func (s *PAL7ConfigBypassConfigUrlBypassRules) GetPaths() []*string {
	return s.Paths
}

func (s *PAL7ConfigBypassConfigUrlBypassRules) SetFroms(v []*string) *PAL7ConfigBypassConfigUrlBypassRules {
	s.Froms = v
	return s
}

func (s *PAL7ConfigBypassConfigUrlBypassRules) SetPaths(v []*string) *PAL7ConfigBypassConfigUrlBypassRules {
	s.Paths = v
	return s
}

func (s *PAL7ConfigBypassConfigUrlBypassRules) Validate() error {
	return dara.Validate(s)
}

type PAL7ConfigDnsConfig struct {
	DnsServers []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
}

func (s PAL7ConfigDnsConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigDnsConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigDnsConfig) GetDnsServers() []*string {
	return s.DnsServers
}

func (s *PAL7ConfigDnsConfig) SetDnsServers(v []*string) *PAL7ConfigDnsConfig {
	s.DnsServers = v
	return s
}

func (s *PAL7ConfigDnsConfig) Validate() error {
	return dara.Validate(s)
}

type PAL7ConfigJsHookConfig struct {
	Mode         *string                  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReplaceRules []*PAL7ConfigReplaceRule `json:"ReplaceRules,omitempty" xml:"ReplaceRules,omitempty" type:"Repeated"`
}

func (s PAL7ConfigJsHookConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigJsHookConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigJsHookConfig) GetMode() *string {
	return s.Mode
}

func (s *PAL7ConfigJsHookConfig) GetReplaceRules() []*PAL7ConfigReplaceRule {
	return s.ReplaceRules
}

func (s *PAL7ConfigJsHookConfig) SetMode(v string) *PAL7ConfigJsHookConfig {
	s.Mode = &v
	return s
}

func (s *PAL7ConfigJsHookConfig) SetReplaceRules(v []*PAL7ConfigReplaceRule) *PAL7ConfigJsHookConfig {
	s.ReplaceRules = v
	return s
}

func (s *PAL7ConfigJsHookConfig) Validate() error {
	if s.ReplaceRules != nil {
		for _, item := range s.ReplaceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PAL7ConfigRequestHeaderRewriteConfig struct {
	Ops []*PAL7ConfigRewriteOp `json:"Ops,omitempty" xml:"Ops,omitempty" type:"Repeated"`
}

func (s PAL7ConfigRequestHeaderRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigRequestHeaderRewriteConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigRequestHeaderRewriteConfig) GetOps() []*PAL7ConfigRewriteOp {
	return s.Ops
}

func (s *PAL7ConfigRequestHeaderRewriteConfig) SetOps(v []*PAL7ConfigRewriteOp) *PAL7ConfigRequestHeaderRewriteConfig {
	s.Ops = v
	return s
}

func (s *PAL7ConfigRequestHeaderRewriteConfig) Validate() error {
	if s.Ops != nil {
		for _, item := range s.Ops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PAL7ConfigRequestQueryRewriteConfig struct {
	Ops []*PAL7ConfigRewriteOp `json:"Ops,omitempty" xml:"Ops,omitempty" type:"Repeated"`
}

func (s PAL7ConfigRequestQueryRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigRequestQueryRewriteConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigRequestQueryRewriteConfig) GetOps() []*PAL7ConfigRewriteOp {
	return s.Ops
}

func (s *PAL7ConfigRequestQueryRewriteConfig) SetOps(v []*PAL7ConfigRewriteOp) *PAL7ConfigRequestQueryRewriteConfig {
	s.Ops = v
	return s
}

func (s *PAL7ConfigRequestQueryRewriteConfig) Validate() error {
	if s.Ops != nil {
		for _, item := range s.Ops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PAL7ConfigResponseHeaderRewriteConfig struct {
	Ops []*PAL7ConfigRewriteOp `json:"Ops,omitempty" xml:"Ops,omitempty" type:"Repeated"`
}

func (s PAL7ConfigResponseHeaderRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigResponseHeaderRewriteConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigResponseHeaderRewriteConfig) GetOps() []*PAL7ConfigRewriteOp {
	return s.Ops
}

func (s *PAL7ConfigResponseHeaderRewriteConfig) SetOps(v []*PAL7ConfigRewriteOp) *PAL7ConfigResponseHeaderRewriteConfig {
	s.Ops = v
	return s
}

func (s *PAL7ConfigResponseHeaderRewriteConfig) Validate() error {
	if s.Ops != nil {
		for _, item := range s.Ops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PAL7ConfigResponseRewriteConfig struct {
	Mode         *string                  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReplaceRules []*PAL7ConfigReplaceRule `json:"ReplaceRules,omitempty" xml:"ReplaceRules,omitempty" type:"Repeated"`
}

func (s PAL7ConfigResponseRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigResponseRewriteConfig) GoString() string {
	return s.String()
}

func (s *PAL7ConfigResponseRewriteConfig) GetMode() *string {
	return s.Mode
}

func (s *PAL7ConfigResponseRewriteConfig) GetReplaceRules() []*PAL7ConfigReplaceRule {
	return s.ReplaceRules
}

func (s *PAL7ConfigResponseRewriteConfig) SetMode(v string) *PAL7ConfigResponseRewriteConfig {
	s.Mode = &v
	return s
}

func (s *PAL7ConfigResponseRewriteConfig) SetReplaceRules(v []*PAL7ConfigReplaceRule) *PAL7ConfigResponseRewriteConfig {
	s.ReplaceRules = v
	return s
}

func (s *PAL7ConfigResponseRewriteConfig) Validate() error {
	if s.ReplaceRules != nil {
		for _, item := range s.ReplaceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
