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
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the address. This ID uniquely identifies the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input parameter serves as the updated note; if an empty value is passed, the note will be deleted.
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
