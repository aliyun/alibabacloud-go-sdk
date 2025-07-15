// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTranscodeParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DescribeLiveDomainTranscodeParamsRequest
	GetSecurityToken() *string
	SetApp(v string) *DescribeLiveDomainTranscodeParamsRequest
	GetApp() *string
	SetPushdomain(v string) *DescribeLiveDomainTranscodeParamsRequest
	GetPushdomain() *string
	SetTemplateName(v string) *DescribeLiveDomainTranscodeParamsRequest
	GetTemplateName() *string
}

type DescribeLiveDomainTranscodeParamsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// This parameter is required.
	Pushdomain   *string `json:"pushdomain,omitempty" xml:"pushdomain,omitempty"`
	TemplateName *string `json:"template_name,omitempty" xml:"template_name,omitempty"`
}

func (s DescribeLiveDomainTranscodeParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTranscodeParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeParamsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveDomainTranscodeParamsRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveDomainTranscodeParamsRequest) GetPushdomain() *string {
	return s.Pushdomain
}

func (s *DescribeLiveDomainTranscodeParamsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeLiveDomainTranscodeParamsRequest) SetSecurityToken(v string) *DescribeLiveDomainTranscodeParamsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsRequest) SetApp(v string) *DescribeLiveDomainTranscodeParamsRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsRequest) SetPushdomain(v string) *DescribeLiveDomainTranscodeParamsRequest {
	s.Pushdomain = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsRequest) SetTemplateName(v string) *DescribeLiveDomainTranscodeParamsRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsRequest) Validate() error {
	return dara.Validate(s)
}
