// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInboundCallIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallDate(v string) *QueryInboundCallIdRequest
	GetCallDate() *string
	SetOutId(v string) *QueryInboundCallIdRequest
	GetOutId() *string
	SetOwnerId(v int64) *QueryInboundCallIdRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryInboundCallIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryInboundCallIdRequest
	GetResourceOwnerId() *int64
}

type QueryInboundCallIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryInboundCallIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInboundCallIdRequest) GoString() string {
	return s.String()
}

func (s *QueryInboundCallIdRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *QueryInboundCallIdRequest) GetOutId() *string {
	return s.OutId
}

func (s *QueryInboundCallIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryInboundCallIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryInboundCallIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryInboundCallIdRequest) SetCallDate(v string) *QueryInboundCallIdRequest {
	s.CallDate = &v
	return s
}

func (s *QueryInboundCallIdRequest) SetOutId(v string) *QueryInboundCallIdRequest {
	s.OutId = &v
	return s
}

func (s *QueryInboundCallIdRequest) SetOwnerId(v int64) *QueryInboundCallIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInboundCallIdRequest) SetResourceOwnerAccount(v string) *QueryInboundCallIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryInboundCallIdRequest) SetResourceOwnerId(v int64) *QueryInboundCallIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryInboundCallIdRequest) Validate() error {
	return dara.Validate(s)
}
