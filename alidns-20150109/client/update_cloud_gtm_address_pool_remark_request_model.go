// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolRemarkRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolRemarkRequest
	GetAddressPoolId() *string
	SetClientToken(v string) *UpdateCloudGtmAddressPoolRemarkRequest
	GetClientToken() *string
	SetRemark(v string) *UpdateCloudGtmAddressPoolRemarkRequest
	GetRemark() *string
}

type UpdateCloudGtmAddressPoolRemarkRequest struct {
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
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
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

func (s UpdateCloudGtmAddressPoolRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolRemarkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolRemarkRequest {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) SetClientToken(v string) *UpdateCloudGtmAddressPoolRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) SetRemark(v string) *UpdateCloudGtmAddressPoolRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkRequest) Validate() error {
	return dara.Validate(s)
}
