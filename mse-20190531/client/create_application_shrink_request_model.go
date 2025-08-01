// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateApplicationShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *CreateApplicationShrinkRequest
	GetAppName() *string
	SetLanguage(v string) *CreateApplicationShrinkRequest
	GetLanguage() *string
	SetNamespace(v string) *CreateApplicationShrinkRequest
	GetNamespace() *string
	SetRegion(v string) *CreateApplicationShrinkRequest
	GetRegion() *string
	SetSentinelEnable(v string) *CreateApplicationShrinkRequest
	GetSentinelEnable() *string
	SetSource(v string) *CreateApplicationShrinkRequest
	GetSource() *string
	SetSwitchEnable(v string) *CreateApplicationShrinkRequest
	GetSwitchEnable() *string
	SetTagsShrink(v string) *CreateApplicationShrinkRequest
	GetTagsShrink() *string
}

type CreateApplicationShrinkRequest struct {
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
	SwitchEnable *string `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
	TagsShrink   *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateApplicationShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateApplicationShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateApplicationShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateApplicationShrinkRequest) GetSentinelEnable() *string {
	return s.SentinelEnable
}

func (s *CreateApplicationShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CreateApplicationShrinkRequest) GetSwitchEnable() *string {
	return s.SwitchEnable
}

func (s *CreateApplicationShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateApplicationShrinkRequest) SetAcceptLanguage(v string) *CreateApplicationShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAppName(v string) *CreateApplicationShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetLanguage(v string) *CreateApplicationShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetNamespace(v string) *CreateApplicationShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRegion(v string) *CreateApplicationShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSentinelEnable(v string) *CreateApplicationShrinkRequest {
	s.SentinelEnable = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSource(v string) *CreateApplicationShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSwitchEnable(v string) *CreateApplicationShrinkRequest {
	s.SwitchEnable = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTagsShrink(v string) *CreateApplicationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
