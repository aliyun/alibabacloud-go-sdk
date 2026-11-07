// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMseServiceApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateMseServiceApplicationRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *CreateMseServiceApplicationRequest
	GetAppName() *string
	SetExtraInfo(v string) *CreateMseServiceApplicationRequest
	GetExtraInfo() *string
	SetLanguage(v string) *CreateMseServiceApplicationRequest
	GetLanguage() *string
	SetMseVersion(v string) *CreateMseServiceApplicationRequest
	GetMseVersion() *string
	SetRegion(v string) *CreateMseServiceApplicationRequest
	GetRegion() *string
	SetSentinelEnable(v string) *CreateMseServiceApplicationRequest
	GetSentinelEnable() *string
	SetSource(v string) *CreateMseServiceApplicationRequest
	GetSource() *string
	SetSwitchEnable(v string) *CreateMseServiceApplicationRequest
	GetSwitchEnable() *string
}

type CreateMseServiceApplicationRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
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
	// bsd-xxyp-open-goods-server
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The version of the Microservices Engine (MSE) cluster to be activated. Valid values:
	//
	// - mse_pro: Professional Edition.
	//
	// - mse_dev: Developer Edition.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The region where the cluster is located, including but not limited to:
	//
	// - cn-hangzhou: Hangzhou
	//
	// - cn-beijing: Beijing
	//
	// - cn-shanghai: Shanghai
	//
	// - cn-zhangjiakou: Zhangjiakou
	//
	// - cn-shenzhen: Shenzhen
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Specifies whether to enable Sentinel.
	//
	// example:
	//
	// true
	SentinelEnable *string `json:"SentinelEnable,omitempty" xml:"SentinelEnable,omitempty"`
	// The source.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Specifies whether to enable Switch.
	//
	// example:
	//
	// true
	SwitchEnable *string `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
}

func (s CreateMseServiceApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMseServiceApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateMseServiceApplicationRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateMseServiceApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateMseServiceApplicationRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateMseServiceApplicationRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateMseServiceApplicationRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *CreateMseServiceApplicationRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateMseServiceApplicationRequest) GetSentinelEnable() *string {
	return s.SentinelEnable
}

func (s *CreateMseServiceApplicationRequest) GetSource() *string {
	return s.Source
}

func (s *CreateMseServiceApplicationRequest) GetSwitchEnable() *string {
	return s.SwitchEnable
}

func (s *CreateMseServiceApplicationRequest) SetAcceptLanguage(v string) *CreateMseServiceApplicationRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetAppName(v string) *CreateMseServiceApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetExtraInfo(v string) *CreateMseServiceApplicationRequest {
	s.ExtraInfo = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetLanguage(v string) *CreateMseServiceApplicationRequest {
	s.Language = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetMseVersion(v string) *CreateMseServiceApplicationRequest {
	s.MseVersion = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetRegion(v string) *CreateMseServiceApplicationRequest {
	s.Region = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetSentinelEnable(v string) *CreateMseServiceApplicationRequest {
	s.SentinelEnable = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetSource(v string) *CreateMseServiceApplicationRequest {
	s.Source = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) SetSwitchEnable(v string) *CreateMseServiceApplicationRequest {
	s.SwitchEnable = &v
	return s
}

func (s *CreateMseServiceApplicationRequest) Validate() error {
	return dara.Validate(s)
}
