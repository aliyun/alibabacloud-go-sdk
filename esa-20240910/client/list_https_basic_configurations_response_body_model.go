// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsBasicConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpsBasicConfigurationsResponseBodyConfigs) *ListHttpsBasicConfigurationsResponseBody
	GetConfigs() []*ListHttpsBasicConfigurationsResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpsBasicConfigurationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetTotalPage() *int32
}

type ListHttpsBasicConfigurationsResponseBody struct {
	Configs    []*ListHttpsBasicConfigurationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpsBasicConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetConfigs() []*ListHttpsBasicConfigurationsResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetConfigs(v []*ListHttpsBasicConfigurationsResponseBodyConfigs) *ListHttpsBasicConfigurationsResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetPageNumber(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetPageSize(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetRequestId(v string) *ListHttpsBasicConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetTotalCount(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetTotalPage(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHttpsBasicConfigurationsResponseBodyConfigs struct {
	Ciphersuite      *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	ConfigId         *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType       *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Http2            *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	Http3            *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	Https            *string `json:"Https,omitempty" xml:"Https,omitempty"`
	OcspStapling     *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable       *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName         *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence         *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Tls10            *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	Tls11            *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	Tls12            *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	Tls13            *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
}

func (s ListHttpsBasicConfigurationsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttp2() *string {
	return s.Http2
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttp3() *string {
	return s.Http3
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttps() *string {
	return s.Https
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls10() *string {
	return s.Tls10
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls11() *string {
	return s.Tls11
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls12() *string {
	return s.Tls12
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls13() *string {
	return s.Tls13
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetCiphersuite(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Ciphersuite = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetCiphersuiteGroup(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.CiphersuiteGroup = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetConfigId(v int64) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetConfigType(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttp2(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Http2 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttp3(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Http3 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttps(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Https = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetOcspStapling(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.OcspStapling = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRule(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRuleEnable(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRuleName(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetSequence(v int32) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls10(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls10 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls11(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls11 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls12(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls12 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls13(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls13 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
