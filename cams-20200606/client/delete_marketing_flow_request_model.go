// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMarketingFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *DeleteMarketingFlowRequest
	GetActivityCode() *string
	SetActivityId(v string) *DeleteMarketingFlowRequest
	GetActivityId() *string
	SetOwnerId(v int64) *DeleteMarketingFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMarketingFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMarketingFlowRequest
	GetResourceOwnerId() *int64
}

type DeleteMarketingFlowRequest struct {
	// example:
	//
	// 1111
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// example:
	//
	// 示例值示例值
	ActivityId           *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMarketingFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMarketingFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteMarketingFlowRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *DeleteMarketingFlowRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *DeleteMarketingFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMarketingFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMarketingFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMarketingFlowRequest) SetActivityCode(v string) *DeleteMarketingFlowRequest {
	s.ActivityCode = &v
	return s
}

func (s *DeleteMarketingFlowRequest) SetActivityId(v string) *DeleteMarketingFlowRequest {
	s.ActivityId = &v
	return s
}

func (s *DeleteMarketingFlowRequest) SetOwnerId(v int64) *DeleteMarketingFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMarketingFlowRequest) SetResourceOwnerAccount(v string) *DeleteMarketingFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMarketingFlowRequest) SetResourceOwnerId(v int64) *DeleteMarketingFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMarketingFlowRequest) Validate() error {
	return dara.Validate(s)
}
