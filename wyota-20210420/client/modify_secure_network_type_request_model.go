// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecureNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecureNetworkType(v string) *ModifySecureNetworkTypeRequest
	GetSecureNetworkType() *string
	SetSerialNo(v string) *ModifySecureNetworkTypeRequest
	GetSerialNo() *string
}

type ModifySecureNetworkTypeRequest struct {
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s ModifySecureNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecureNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeRequest) GetSecureNetworkType() *string {
	return s.SecureNetworkType
}

func (s *ModifySecureNetworkTypeRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ModifySecureNetworkTypeRequest) SetSecureNetworkType(v string) *ModifySecureNetworkTypeRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *ModifySecureNetworkTypeRequest) SetSerialNo(v string) *ModifySecureNetworkTypeRequest {
	s.SerialNo = &v
	return s
}

func (s *ModifySecureNetworkTypeRequest) Validate() error {
	return dara.Validate(s)
}
