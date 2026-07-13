// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressRemarkRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *UpdateCloudGtmAddressRemarkRequest
	GetAddressId() *string
	SetClientToken(v string) *UpdateCloudGtmAddressRemarkRequest
	GetClientToken() *string
	SetRemark(v string) *UpdateCloudGtmAddressRemarkRequest
	GetRemark() *string
}

type UpdateCloudGtmAddressRemarkRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The token must be unique across requests and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new remarks for the address. To delete the remarks, leave this parameter empty.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateCloudGtmAddressRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressRemarkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressRemarkRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *UpdateCloudGtmAddressRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCloudGtmAddressRemarkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressRemarkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkRequest) SetAddressId(v string) *UpdateCloudGtmAddressRemarkRequest {
	s.AddressId = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkRequest) SetClientToken(v string) *UpdateCloudGtmAddressRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkRequest) SetRemark(v string) *UpdateCloudGtmAddressRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkRequest) Validate() error {
	return dara.Validate(s)
}
