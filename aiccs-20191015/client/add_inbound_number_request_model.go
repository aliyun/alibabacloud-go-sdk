// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInboundNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *AddInboundNumberRequest
	GetApplicationCode() *string
	SetInboundNumbers(v []*string) *AddInboundNumberRequest
	GetInboundNumbers() []*string
	SetInboundType(v int64) *AddInboundNumberRequest
	GetInboundType() *int64
	SetLineCode(v string) *AddInboundNumberRequest
	GetLineCode() *string
	SetOwnerId(v int64) *AddInboundNumberRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddInboundNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddInboundNumberRequest
	GetResourceOwnerId() *int64
}

type AddInboundNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DLWERWLD
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// This parameter is required.
	InboundNumbers []*string `json:"InboundNumbers,omitempty" xml:"InboundNumbers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 8
	InboundType *int64 `json:"InboundType,omitempty" xml:"InboundType,omitempty"`
	// example:
	//
	// JILIANG_*****_TEST_NET
	LineCode             *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddInboundNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInboundNumberRequest) GoString() string {
	return s.String()
}

func (s *AddInboundNumberRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *AddInboundNumberRequest) GetInboundNumbers() []*string {
	return s.InboundNumbers
}

func (s *AddInboundNumberRequest) GetInboundType() *int64 {
	return s.InboundType
}

func (s *AddInboundNumberRequest) GetLineCode() *string {
	return s.LineCode
}

func (s *AddInboundNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddInboundNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddInboundNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddInboundNumberRequest) SetApplicationCode(v string) *AddInboundNumberRequest {
	s.ApplicationCode = &v
	return s
}

func (s *AddInboundNumberRequest) SetInboundNumbers(v []*string) *AddInboundNumberRequest {
	s.InboundNumbers = v
	return s
}

func (s *AddInboundNumberRequest) SetInboundType(v int64) *AddInboundNumberRequest {
	s.InboundType = &v
	return s
}

func (s *AddInboundNumberRequest) SetLineCode(v string) *AddInboundNumberRequest {
	s.LineCode = &v
	return s
}

func (s *AddInboundNumberRequest) SetOwnerId(v int64) *AddInboundNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *AddInboundNumberRequest) SetResourceOwnerAccount(v string) *AddInboundNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddInboundNumberRequest) SetResourceOwnerId(v int64) *AddInboundNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddInboundNumberRequest) Validate() error {
	return dara.Validate(s)
}
