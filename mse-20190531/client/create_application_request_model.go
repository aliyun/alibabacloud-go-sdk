// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateApplicationRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *CreateApplicationRequest
	GetAppName() *string
	SetLanguage(v string) *CreateApplicationRequest
	GetLanguage() *string
	SetNamespace(v string) *CreateApplicationRequest
	GetNamespace() *string
	SetRegion(v string) *CreateApplicationRequest
	GetRegion() *string
	SetSentinelEnable(v string) *CreateApplicationRequest
	GetSentinelEnable() *string
	SetSource(v string) *CreateApplicationRequest
	GetSource() *string
	SetSwitchEnable(v string) *CreateApplicationRequest
	GetSwitchEnable() *string
	SetTags(v []*CreateApplicationRequestTags) *CreateApplicationRequest
	GetTags() []*CreateApplicationRequestTags
}

type CreateApplicationRequest struct {
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
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// MSE命名空间名字。
	//
	// example:
	//
	// prod
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Specifies whether to start the switch.
	//
	// example:
	//
	// true
	SentinelEnable *string `json:"SentinelEnable,omitempty" xml:"SentinelEnable,omitempty"`
	// The service where the application is deployed. A value of ACK indicates Container Service for Kubernetes.
	//
	// example:
	//
	// ACK
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// true
	SwitchEnable *string                         `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
	Tags         []*CreateApplicationRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateApplicationRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateApplicationRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateApplicationRequest) GetSentinelEnable() *string {
	return s.SentinelEnable
}

func (s *CreateApplicationRequest) GetSource() *string {
	return s.Source
}

func (s *CreateApplicationRequest) GetSwitchEnable() *string {
	return s.SwitchEnable
}

func (s *CreateApplicationRequest) GetTags() []*CreateApplicationRequestTags {
	return s.Tags
}

func (s *CreateApplicationRequest) SetAcceptLanguage(v string) *CreateApplicationRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetLanguage(v string) *CreateApplicationRequest {
	s.Language = &v
	return s
}

func (s *CreateApplicationRequest) SetNamespace(v string) *CreateApplicationRequest {
	s.Namespace = &v
	return s
}

func (s *CreateApplicationRequest) SetRegion(v string) *CreateApplicationRequest {
	s.Region = &v
	return s
}

func (s *CreateApplicationRequest) SetSentinelEnable(v string) *CreateApplicationRequest {
	s.SentinelEnable = &v
	return s
}

func (s *CreateApplicationRequest) SetSource(v string) *CreateApplicationRequest {
	s.Source = &v
	return s
}

func (s *CreateApplicationRequest) SetSwitchEnable(v string) *CreateApplicationRequest {
	s.SwitchEnable = &v
	return s
}

func (s *CreateApplicationRequest) SetTags(v []*CreateApplicationRequestTags) *CreateApplicationRequest {
	s.Tags = v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApplicationRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestTags) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateApplicationRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateApplicationRequestTags) SetKey(v string) *CreateApplicationRequestTags {
	s.Key = &v
	return s
}

func (s *CreateApplicationRequestTags) SetValue(v string) *CreateApplicationRequestTags {
	s.Value = &v
	return s
}

func (s *CreateApplicationRequestTags) Validate() error {
	return dara.Validate(s)
}
