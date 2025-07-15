// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeFailoverTestJobRequest
	GetClientToken() *string
	SetJobId(v string) *DescribeFailoverTestJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *DescribeFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeFailoverTestJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeFailoverTestJobRequest
	GetResourceOwnerAccount() *string
}

type DescribeFailoverTestJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the failover test.
	//
	// This parameter is required.
	//
	// example:
	//
	// ftj-xxxxxxxxx
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the failover test.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFailoverTestJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFailoverTestJobRequest) SetClientToken(v string) *DescribeFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) SetJobId(v string) *DescribeFailoverTestJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) SetOwnerAccount(v string) *DescribeFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) SetOwnerId(v int64) *DescribeFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) SetRegionId(v string) *DescribeFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) SetResourceOwnerAccount(v string) *DescribeFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
