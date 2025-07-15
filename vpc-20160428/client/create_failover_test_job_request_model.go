// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFailoverTestJobRequest
	GetClientToken() *string
	SetDescription(v string) *CreateFailoverTestJobRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFailoverTestJobRequest
	GetDryRun() *bool
	SetJobDuration(v int32) *CreateFailoverTestJobRequest
	GetJobDuration() *int32
	SetJobType(v string) *CreateFailoverTestJobRequest
	GetJobType() *string
	SetName(v string) *CreateFailoverTestJobRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateFailoverTestJobRequest
	GetRegionId() *string
	SetResourceId(v []*string) *CreateFailoverTestJobRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *CreateFailoverTestJobRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *CreateFailoverTestJobRequest
	GetResourceType() *string
}

type CreateFailoverTestJobRequest struct {
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
	// The description of the failover test.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// If you set the value to true, the system performs only a dry run without actually performing the actual request. If you set the value to false, the system performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration of the failover test. Unit: minutes. Valid values: **1 to 4320**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	JobDuration *int32 `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The type of the failover test. Valid values:
	//
	// 	- **StartNow**
	//
	// 	- **StartLater**
	//
	// This parameter is required.
	//
	// example:
	//
	// StartNow
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the failover test.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the failover test.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of failover test resources. You can add at most 16 resources.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type of the failover test. Set the value to **PHYSICALCONNECTION**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFailoverTestJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFailoverTestJobRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFailoverTestJobRequest) GetJobDuration() *int32 {
	return s.JobDuration
}

func (s *CreateFailoverTestJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateFailoverTestJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFailoverTestJobRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *CreateFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFailoverTestJobRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateFailoverTestJobRequest) SetClientToken(v string) *CreateFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetDescription(v string) *CreateFailoverTestJobRequest {
	s.Description = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetDryRun(v bool) *CreateFailoverTestJobRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetJobDuration(v int32) *CreateFailoverTestJobRequest {
	s.JobDuration = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetJobType(v string) *CreateFailoverTestJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetName(v string) *CreateFailoverTestJobRequest {
	s.Name = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetOwnerAccount(v string) *CreateFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetOwnerId(v int64) *CreateFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetRegionId(v string) *CreateFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetResourceId(v []*string) *CreateFailoverTestJobRequest {
	s.ResourceId = v
	return s
}

func (s *CreateFailoverTestJobRequest) SetResourceOwnerAccount(v string) *CreateFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFailoverTestJobRequest) SetResourceType(v string) *CreateFailoverTestJobRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
