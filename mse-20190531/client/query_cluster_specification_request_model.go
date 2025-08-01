// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterSpecificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryClusterSpecificationRequest
	GetAcceptLanguage() *string
	SetConnectType(v string) *QueryClusterSpecificationRequest
	GetConnectType() *string
	SetMseVersion(v string) *QueryClusterSpecificationRequest
	GetMseVersion() *string
}

type QueryClusterSpecificationRequest struct {
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
	// The network type. Valid values:
	//
	// 	- slb
	//
	// 	- eni
	//
	// example:
	//
	// slb
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The edition of the MSE instance that you want to purchase.
	//
	// 	- mse_pro: Professional Edition
	//
	// 	- mse_dev: Developer Edition
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
}

func (s QueryClusterSpecificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterSpecificationRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryClusterSpecificationRequest) GetConnectType() *string {
	return s.ConnectType
}

func (s *QueryClusterSpecificationRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *QueryClusterSpecificationRequest) SetAcceptLanguage(v string) *QueryClusterSpecificationRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryClusterSpecificationRequest) SetConnectType(v string) *QueryClusterSpecificationRequest {
	s.ConnectType = &v
	return s
}

func (s *QueryClusterSpecificationRequest) SetMseVersion(v string) *QueryClusterSpecificationRequest {
	s.MseVersion = &v
	return s
}

func (s *QueryClusterSpecificationRequest) Validate() error {
	return dara.Validate(s)
}
