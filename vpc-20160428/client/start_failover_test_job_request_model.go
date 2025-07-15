// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartFailoverTestJobRequest
	GetClientToken() *string
	SetJobId(v string) *StartFailoverTestJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *StartFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartFailoverTestJobRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartFailoverTestJobRequest
	GetResourceOwnerAccount() *string
}

type StartFailoverTestJobRequest struct {
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

func (s StartFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *StartFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartFailoverTestJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StartFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartFailoverTestJobRequest) SetClientToken(v string) *StartFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StartFailoverTestJobRequest) SetJobId(v string) *StartFailoverTestJobRequest {
	s.JobId = &v
	return s
}

func (s *StartFailoverTestJobRequest) SetOwnerAccount(v string) *StartFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartFailoverTestJobRequest) SetOwnerId(v int64) *StartFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartFailoverTestJobRequest) SetRegionId(v string) *StartFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartFailoverTestJobRequest) SetResourceOwnerAccount(v string) *StartFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
