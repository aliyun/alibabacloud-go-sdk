// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDConversionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *UAIDConversionRequest
	GetAuthCode() *string
	SetCarrier(v string) *UAIDConversionRequest
	GetCarrier() *string
	SetOutId(v string) *UAIDConversionRequest
	GetOutId() *string
	SetOwnerId(v int64) *UAIDConversionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UAIDConversionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UAIDConversionRequest
	GetResourceOwnerId() *int64
	SetUaidList(v string) *UAIDConversionRequest
	GetUaidList() *string
}

type UAIDConversionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	UaidList *string `json:"UaidList,omitempty" xml:"UaidList,omitempty"`
}

func (s UAIDConversionRequest) String() string {
	return dara.Prettify(s)
}

func (s UAIDConversionRequest) GoString() string {
	return s.String()
}

func (s *UAIDConversionRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *UAIDConversionRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *UAIDConversionRequest) GetOutId() *string {
	return s.OutId
}

func (s *UAIDConversionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UAIDConversionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UAIDConversionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UAIDConversionRequest) GetUaidList() *string {
	return s.UaidList
}

func (s *UAIDConversionRequest) SetAuthCode(v string) *UAIDConversionRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDConversionRequest) SetCarrier(v string) *UAIDConversionRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDConversionRequest) SetOutId(v string) *UAIDConversionRequest {
	s.OutId = &v
	return s
}

func (s *UAIDConversionRequest) SetOwnerId(v int64) *UAIDConversionRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDConversionRequest) SetResourceOwnerAccount(v string) *UAIDConversionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDConversionRequest) SetResourceOwnerId(v int64) *UAIDConversionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDConversionRequest) SetUaidList(v string) *UAIDConversionRequest {
	s.UaidList = &v
	return s
}

func (s *UAIDConversionRequest) Validate() error {
	return dara.Validate(s)
}
