// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAuthPolicyRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListAuthPolicyRequest
	GetAppId() *string
	SetName(v string) *ListAuthPolicyRequest
	GetName() *string
	SetNamespace(v string) *ListAuthPolicyRequest
	GetNamespace() *string
	SetPageNumber(v string) *ListAuthPolicyRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthPolicyRequest
	GetPageSize() *string
	SetProtocol(v string) *ListAuthPolicyRequest
	GetProtocol() *string
	SetRegion(v string) *ListAuthPolicyRequest
	GetRegion() *string
	SetSource(v string) *ListAuthPolicyRequest
	GetSource() *string
}

type ListAuthPolicyRequest struct {
	// The language of the response. Valid values: zh-CN and en-US. Default value: zh-CN. The value zh-CN indicates Chinese and the value en-US indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID.
	//
	// example:
	//
	// jgy4cadmqo@***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the authentication rule.
	//
	// example:
	//
	// auto-rule-**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **SPRING_CLOUD**
	//
	// 	- **DUBBO**
	//
	// 	- **istio**
	//
	// example:
	//
	// SPRING_CLOUD
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The service source.
	//
	// This parameter is required.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListAuthPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAuthPolicyRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ListAuthPolicyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAuthPolicyRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthPolicyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthPolicyRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListAuthPolicyRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAuthPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *ListAuthPolicyRequest) SetAcceptLanguage(v string) *ListAuthPolicyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAuthPolicyRequest) SetAppId(v string) *ListAuthPolicyRequest {
	s.AppId = &v
	return s
}

func (s *ListAuthPolicyRequest) SetName(v string) *ListAuthPolicyRequest {
	s.Name = &v
	return s
}

func (s *ListAuthPolicyRequest) SetNamespace(v string) *ListAuthPolicyRequest {
	s.Namespace = &v
	return s
}

func (s *ListAuthPolicyRequest) SetPageNumber(v string) *ListAuthPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthPolicyRequest) SetPageSize(v string) *ListAuthPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthPolicyRequest) SetProtocol(v string) *ListAuthPolicyRequest {
	s.Protocol = &v
	return s
}

func (s *ListAuthPolicyRequest) SetRegion(v string) *ListAuthPolicyRequest {
	s.Region = &v
	return s
}

func (s *ListAuthPolicyRequest) SetSource(v string) *ListAuthPolicyRequest {
	s.Source = &v
	return s
}

func (s *ListAuthPolicyRequest) Validate() error {
	return dara.Validate(s)
}
