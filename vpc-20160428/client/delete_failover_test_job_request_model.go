// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFailoverTestJobRequest
	GetClientToken() *string
	SetJobId(v string) *DeleteFailoverTestJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *DeleteFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteFailoverTestJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteFailoverTestJobRequest
	GetResourceOwnerAccount() *string
}

type DeleteFailoverTestJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the failover test job. This ID is returned by the CreateFailoverTestJob operation and can also be obtained by calling DescribeFailoverTestJobs. You cannot construct this value manually.
	//
	// This parameter is required.
	//
	// example:
	//
	// ftj-bp1yh6mvi13aq3g8w****
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the failover test job.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query region IDs.
	//
	// example:
	//
	// ch-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFailoverTestJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFailoverTestJobRequest) SetClientToken(v string) *DeleteFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) SetJobId(v string) *DeleteFailoverTestJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) SetOwnerAccount(v string) *DeleteFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) SetOwnerId(v int64) *DeleteFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) SetRegionId(v string) *DeleteFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) SetResourceOwnerAccount(v string) *DeleteFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
