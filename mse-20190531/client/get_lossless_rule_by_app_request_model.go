// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLosslessRuleByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetLosslessRuleByAppRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetLosslessRuleByAppRequest
	GetAppId() *string
	SetAppName(v string) *GetLosslessRuleByAppRequest
	GetAppName() *string
	SetNamespace(v string) *GetLosslessRuleByAppRequest
	GetNamespace() *string
	SetRegionId(v string) *GetLosslessRuleByAppRequest
	GetRegionId() *string
}

type GetLosslessRuleByAppRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522baa***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLosslessRuleByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLosslessRuleByAppRequest) GoString() string {
	return s.String()
}

func (s *GetLosslessRuleByAppRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetLosslessRuleByAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetLosslessRuleByAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLosslessRuleByAppRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLosslessRuleByAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLosslessRuleByAppRequest) SetAcceptLanguage(v string) *GetLosslessRuleByAppRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetLosslessRuleByAppRequest) SetAppId(v string) *GetLosslessRuleByAppRequest {
	s.AppId = &v
	return s
}

func (s *GetLosslessRuleByAppRequest) SetAppName(v string) *GetLosslessRuleByAppRequest {
	s.AppName = &v
	return s
}

func (s *GetLosslessRuleByAppRequest) SetNamespace(v string) *GetLosslessRuleByAppRequest {
	s.Namespace = &v
	return s
}

func (s *GetLosslessRuleByAppRequest) SetRegionId(v string) *GetLosslessRuleByAppRequest {
	s.RegionId = &v
	return s
}

func (s *GetLosslessRuleByAppRequest) Validate() error {
	return dara.Validate(s)
}
