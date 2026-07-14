// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseMarketingFLowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *PauseMarketingFLowRequest
	GetActivityCode() *string
	SetActivityId(v string) *PauseMarketingFLowRequest
	GetActivityId() *string
	SetOwnerId(v int64) *PauseMarketingFLowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PauseMarketingFLowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PauseMarketingFLowRequest
	GetResourceOwnerId() *int64
}

type PauseMarketingFLowRequest struct {
	// example:
	//
	// 439859845**234
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// example:
	//
	// N/A
	ActivityId           *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PauseMarketingFLowRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseMarketingFLowRequest) GoString() string {
	return s.String()
}

func (s *PauseMarketingFLowRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *PauseMarketingFLowRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *PauseMarketingFLowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PauseMarketingFLowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PauseMarketingFLowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PauseMarketingFLowRequest) SetActivityCode(v string) *PauseMarketingFLowRequest {
	s.ActivityCode = &v
	return s
}

func (s *PauseMarketingFLowRequest) SetActivityId(v string) *PauseMarketingFLowRequest {
	s.ActivityId = &v
	return s
}

func (s *PauseMarketingFLowRequest) SetOwnerId(v int64) *PauseMarketingFLowRequest {
	s.OwnerId = &v
	return s
}

func (s *PauseMarketingFLowRequest) SetResourceOwnerAccount(v string) *PauseMarketingFLowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PauseMarketingFLowRequest) SetResourceOwnerId(v int64) *PauseMarketingFLowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PauseMarketingFLowRequest) Validate() error {
	return dara.Validate(s)
}
