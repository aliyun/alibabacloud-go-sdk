// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAliyunCertUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunPk(v string) *GenerateAliyunCertUrlRequest
	GetAliyunPk() *string
	SetAppName(v string) *GenerateAliyunCertUrlRequest
	GetAppName() *string
	SetApplyChannel(v string) *GenerateAliyunCertUrlRequest
	GetApplyChannel() *string
	SetApplyType(v string) *GenerateAliyunCertUrlRequest
	GetApplyType() *string
	SetCallback(v string) *GenerateAliyunCertUrlRequest
	GetCallback() *string
	SetCertWay(v string) *GenerateAliyunCertUrlRequest
	GetCertWay() *string
	SetIgnoreAlreadyCert(v bool) *GenerateAliyunCertUrlRequest
	GetIgnoreAlreadyCert() *bool
	SetIsMobile(v bool) *GenerateAliyunCertUrlRequest
	GetIsMobile() *bool
	SetIsOpenApp(v bool) *GenerateAliyunCertUrlRequest
	GetIsOpenApp() *bool
	SetPlatform(v string) *GenerateAliyunCertUrlRequest
	GetPlatform() *string
	SetSource(v string) *GenerateAliyunCertUrlRequest
	GetSource() *string
	SetSubjectType(v string) *GenerateAliyunCertUrlRequest
	GetSubjectType() *string
}

type GenerateAliyunCertUrlRequest struct {
	// This parameter is required.
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// This parameter is required.
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ApplyChannel      *string `json:"ApplyChannel,omitempty" xml:"ApplyChannel,omitempty"`
	ApplyType         *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	Callback          *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CertWay           *string `json:"CertWay,omitempty" xml:"CertWay,omitempty"`
	IgnoreAlreadyCert *bool   `json:"IgnoreAlreadyCert,omitempty" xml:"IgnoreAlreadyCert,omitempty"`
	IsMobile          *bool   `json:"IsMobile,omitempty" xml:"IsMobile,omitempty"`
	IsOpenApp         *bool   `json:"IsOpenApp,omitempty" xml:"IsOpenApp,omitempty"`
	// This parameter is required.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s GenerateAliyunCertUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAliyunCertUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateAliyunCertUrlRequest) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *GenerateAliyunCertUrlRequest) GetAppName() *string {
	return s.AppName
}

func (s *GenerateAliyunCertUrlRequest) GetApplyChannel() *string {
	return s.ApplyChannel
}

func (s *GenerateAliyunCertUrlRequest) GetApplyType() *string {
	return s.ApplyType
}

func (s *GenerateAliyunCertUrlRequest) GetCallback() *string {
	return s.Callback
}

func (s *GenerateAliyunCertUrlRequest) GetCertWay() *string {
	return s.CertWay
}

func (s *GenerateAliyunCertUrlRequest) GetIgnoreAlreadyCert() *bool {
	return s.IgnoreAlreadyCert
}

func (s *GenerateAliyunCertUrlRequest) GetIsMobile() *bool {
	return s.IsMobile
}

func (s *GenerateAliyunCertUrlRequest) GetIsOpenApp() *bool {
	return s.IsOpenApp
}

func (s *GenerateAliyunCertUrlRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GenerateAliyunCertUrlRequest) GetSource() *string {
	return s.Source
}

func (s *GenerateAliyunCertUrlRequest) GetSubjectType() *string {
	return s.SubjectType
}

func (s *GenerateAliyunCertUrlRequest) SetAliyunPk(v string) *GenerateAliyunCertUrlRequest {
	s.AliyunPk = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetAppName(v string) *GenerateAliyunCertUrlRequest {
	s.AppName = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetApplyChannel(v string) *GenerateAliyunCertUrlRequest {
	s.ApplyChannel = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetApplyType(v string) *GenerateAliyunCertUrlRequest {
	s.ApplyType = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetCallback(v string) *GenerateAliyunCertUrlRequest {
	s.Callback = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetCertWay(v string) *GenerateAliyunCertUrlRequest {
	s.CertWay = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetIgnoreAlreadyCert(v bool) *GenerateAliyunCertUrlRequest {
	s.IgnoreAlreadyCert = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetIsMobile(v bool) *GenerateAliyunCertUrlRequest {
	s.IsMobile = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetIsOpenApp(v bool) *GenerateAliyunCertUrlRequest {
	s.IsOpenApp = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetPlatform(v string) *GenerateAliyunCertUrlRequest {
	s.Platform = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetSource(v string) *GenerateAliyunCertUrlRequest {
	s.Source = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) SetSubjectType(v string) *GenerateAliyunCertUrlRequest {
	s.SubjectType = &v
	return s
}

func (s *GenerateAliyunCertUrlRequest) Validate() error {
	return dara.Validate(s)
}
