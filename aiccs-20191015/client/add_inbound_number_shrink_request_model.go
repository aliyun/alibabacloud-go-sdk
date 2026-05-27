// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInboundNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *AddInboundNumberShrinkRequest
	GetApplicationCode() *string
	SetInboundNumbersShrink(v string) *AddInboundNumberShrinkRequest
	GetInboundNumbersShrink() *string
	SetInboundType(v int64) *AddInboundNumberShrinkRequest
	GetInboundType() *int64
	SetLineCode(v string) *AddInboundNumberShrinkRequest
	GetLineCode() *string
	SetOwnerId(v int64) *AddInboundNumberShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddInboundNumberShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddInboundNumberShrinkRequest
	GetResourceOwnerId() *int64
}

type AddInboundNumberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DLWERWLD
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// This parameter is required.
	InboundNumbersShrink *string `json:"InboundNumbers,omitempty" xml:"InboundNumbers,omitempty"`
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

func (s AddInboundNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInboundNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddInboundNumberShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *AddInboundNumberShrinkRequest) GetInboundNumbersShrink() *string {
	return s.InboundNumbersShrink
}

func (s *AddInboundNumberShrinkRequest) GetInboundType() *int64 {
	return s.InboundType
}

func (s *AddInboundNumberShrinkRequest) GetLineCode() *string {
	return s.LineCode
}

func (s *AddInboundNumberShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddInboundNumberShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddInboundNumberShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddInboundNumberShrinkRequest) SetApplicationCode(v string) *AddInboundNumberShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetInboundNumbersShrink(v string) *AddInboundNumberShrinkRequest {
	s.InboundNumbersShrink = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetInboundType(v int64) *AddInboundNumberShrinkRequest {
	s.InboundType = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetLineCode(v string) *AddInboundNumberShrinkRequest {
	s.LineCode = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetOwnerId(v int64) *AddInboundNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetResourceOwnerAccount(v string) *AddInboundNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) SetResourceOwnerId(v int64) *AddInboundNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddInboundNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
