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
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89528023225442****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// A client token to ensure the idempotence of the request. Generate a unique value from your client. The token can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new remarks for the address pool. If you leave this parameter empty, the remarks are deleted.
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
