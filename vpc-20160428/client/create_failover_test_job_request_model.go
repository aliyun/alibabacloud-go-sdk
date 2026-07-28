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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the failover test job.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - **true**: sends the request without creating the failover test node. The system checks the request for potential issues, including whether the AccessKey is valid, the authorization of the Resource Access Management (RAM) user, and whether required parameters are specified. If the check fails, the corresponding error is returned. If the check passes, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a Normal request, and the failover test job is created after the check passes. A 2xx HTTP status code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration of the failover test job. Unit: minutes. Valid values: **1 to 4320**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	JobDuration *int32 `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The type of the failover test job. Valid values:
	//
	// - **StartNow**: The failover test starts immediately after the job is created.
	//
	// - **StartLater**: Only the job is created. The failover test does not start.
	//
	// This parameter is required.
	//
	// example:
	//
	// StartNow
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the failover test job.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the failover test job.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs to test. You can add up to 16 resources.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the resource to test. Valid values: **PHYSICALCONNECTION**: Express Connect circuit.
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
