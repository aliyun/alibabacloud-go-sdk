// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopFailoverTestJobRequest
	GetClientToken() *string
	SetJobId(v string) *StopFailoverTestJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *StopFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopFailoverTestJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopFailoverTestJobRequest
	GetResourceOwnerAccount() *string
}

type StopFailoverTestJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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

func (s StopFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *StopFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopFailoverTestJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StopFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopFailoverTestJobRequest) SetClientToken(v string) *StopFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StopFailoverTestJobRequest) SetJobId(v string) *StopFailoverTestJobRequest {
	s.JobId = &v
	return s
}

func (s *StopFailoverTestJobRequest) SetOwnerAccount(v string) *StopFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopFailoverTestJobRequest) SetOwnerId(v int64) *StopFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopFailoverTestJobRequest) SetRegionId(v string) *StopFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopFailoverTestJobRequest) SetResourceOwnerAccount(v string) *StopFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
