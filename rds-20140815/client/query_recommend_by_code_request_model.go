// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecommendByCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRecommendByCodeRequest
	GetCode() *string
	SetOwnerId(v string) *QueryRecommendByCodeRequest
	GetOwnerId() *string
	SetResourceGroupId(v string) *QueryRecommendByCodeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *QueryRecommendByCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRecommendByCodeRequest
	GetResourceOwnerId() *int64
}

type QueryRecommendByCodeRequest struct {
	// The code.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_recommend
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRecommendByCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecommendByCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryRecommendByCodeRequest) GetCode() *string {
	return s.Code
}

func (s *QueryRecommendByCodeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryRecommendByCodeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryRecommendByCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRecommendByCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRecommendByCodeRequest) SetCode(v string) *QueryRecommendByCodeRequest {
	s.Code = &v
	return s
}

func (s *QueryRecommendByCodeRequest) SetOwnerId(v string) *QueryRecommendByCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecommendByCodeRequest) SetResourceGroupId(v string) *QueryRecommendByCodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryRecommendByCodeRequest) SetResourceOwnerAccount(v string) *QueryRecommendByCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRecommendByCodeRequest) SetResourceOwnerId(v int64) *QueryRecommendByCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRecommendByCodeRequest) Validate() error {
	return dara.Validate(s)
}
