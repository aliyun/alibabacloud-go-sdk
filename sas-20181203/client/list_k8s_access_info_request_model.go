// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sAccessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunYundunGatewayApiName(v string) *ListK8sAccessInfoRequest
	GetAliyunYundunGatewayApiName() *string
	SetAliyunYundunGatewayPopName(v string) *ListK8sAccessInfoRequest
	GetAliyunYundunGatewayPopName() *string
	SetAliyunYundunGatewayProjectName(v string) *ListK8sAccessInfoRequest
	GetAliyunYundunGatewayProjectName() *string
	SetLang(v string) *ListK8sAccessInfoRequest
	GetLang() *string
}

type ListK8sAccessInfoRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayApiName *string `json:"AliyunYundunGatewayApiName,omitempty" xml:"AliyunYundunGatewayApiName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayPopName *string `json:"AliyunYundunGatewayPopName,omitempty" xml:"AliyunYundunGatewayPopName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayProjectName *string `json:"AliyunYundunGatewayProjectName,omitempty" xml:"AliyunYundunGatewayProjectName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListK8sAccessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sAccessInfoRequest) GoString() string {
	return s.String()
}

func (s *ListK8sAccessInfoRequest) GetAliyunYundunGatewayApiName() *string {
	return s.AliyunYundunGatewayApiName
}

func (s *ListK8sAccessInfoRequest) GetAliyunYundunGatewayPopName() *string {
	return s.AliyunYundunGatewayPopName
}

func (s *ListK8sAccessInfoRequest) GetAliyunYundunGatewayProjectName() *string {
	return s.AliyunYundunGatewayProjectName
}

func (s *ListK8sAccessInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListK8sAccessInfoRequest) SetAliyunYundunGatewayApiName(v string) *ListK8sAccessInfoRequest {
	s.AliyunYundunGatewayApiName = &v
	return s
}

func (s *ListK8sAccessInfoRequest) SetAliyunYundunGatewayPopName(v string) *ListK8sAccessInfoRequest {
	s.AliyunYundunGatewayPopName = &v
	return s
}

func (s *ListK8sAccessInfoRequest) SetAliyunYundunGatewayProjectName(v string) *ListK8sAccessInfoRequest {
	s.AliyunYundunGatewayProjectName = &v
	return s
}

func (s *ListK8sAccessInfoRequest) SetLang(v string) *ListK8sAccessInfoRequest {
	s.Lang = &v
	return s
}

func (s *ListK8sAccessInfoRequest) Validate() error {
	return dara.Validate(s)
}
