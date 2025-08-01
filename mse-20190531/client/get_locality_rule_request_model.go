// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetLocalityRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetLocalityRuleRequest
	GetAppId() *string
	SetAppName(v string) *GetLocalityRuleRequest
	GetAppName() *string
	SetNamespace(v string) *GetLocalityRuleRequest
	GetNamespace() *string
	SetRegion(v string) *GetLocalityRuleRequest
	GetRegion() *string
	SetSource(v string) *GetLocalityRuleRequest
	GetSource() *string
}

type GetLocalityRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// hgxznfcvbe@be2c0228f******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetLocalityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLocalityRuleRequest) GoString() string {
	return s.String()
}

func (s *GetLocalityRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetLocalityRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetLocalityRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLocalityRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLocalityRuleRequest) GetRegion() *string {
	return s.Region
}

func (s *GetLocalityRuleRequest) GetSource() *string {
	return s.Source
}

func (s *GetLocalityRuleRequest) SetAcceptLanguage(v string) *GetLocalityRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetLocalityRuleRequest) SetAppId(v string) *GetLocalityRuleRequest {
	s.AppId = &v
	return s
}

func (s *GetLocalityRuleRequest) SetAppName(v string) *GetLocalityRuleRequest {
	s.AppName = &v
	return s
}

func (s *GetLocalityRuleRequest) SetNamespace(v string) *GetLocalityRuleRequest {
	s.Namespace = &v
	return s
}

func (s *GetLocalityRuleRequest) SetRegion(v string) *GetLocalityRuleRequest {
	s.Region = &v
	return s
}

func (s *GetLocalityRuleRequest) SetSource(v string) *GetLocalityRuleRequest {
	s.Source = &v
	return s
}

func (s *GetLocalityRuleRequest) Validate() error {
	return dara.Validate(s)
}
