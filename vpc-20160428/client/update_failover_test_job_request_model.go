// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFailoverTestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateFailoverTestJobRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateFailoverTestJobRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateFailoverTestJobRequest
	GetDryRun() *bool
	SetJobDuration(v int32) *UpdateFailoverTestJobRequest
	GetJobDuration() *int32
	SetJobId(v string) *UpdateFailoverTestJobRequest
	GetJobId() *string
	SetName(v string) *UpdateFailoverTestJobRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdateFailoverTestJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateFailoverTestJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateFailoverTestJobRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UpdateFailoverTestJobRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *UpdateFailoverTestJobRequest
	GetResourceOwnerAccount() *string
}

type UpdateFailoverTestJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system uses the value of **RequestId*	- as **ClientToken**. The value of **RequestId*	- for each API request is different.
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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration of the failover test. Unit: minutes. Valid values: **1*	- to **4320**.
	//
	// example:
	//
	// 60
	JobDuration *int32 `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The ID of the failover test.
	//
	// This parameter is required.
	//
	// example:
	//
	// ftj-xxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
	// The IDs of the failover test resources. You can add at most 16 resources.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s UpdateFailoverTestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFailoverTestJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateFailoverTestJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateFailoverTestJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFailoverTestJobRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateFailoverTestJobRequest) GetJobDuration() *int32 {
	return s.JobDuration
}

func (s *UpdateFailoverTestJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateFailoverTestJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFailoverTestJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateFailoverTestJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateFailoverTestJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateFailoverTestJobRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UpdateFailoverTestJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateFailoverTestJobRequest) SetClientToken(v string) *UpdateFailoverTestJobRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetDescription(v string) *UpdateFailoverTestJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetDryRun(v bool) *UpdateFailoverTestJobRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetJobDuration(v int32) *UpdateFailoverTestJobRequest {
	s.JobDuration = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetJobId(v string) *UpdateFailoverTestJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetName(v string) *UpdateFailoverTestJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetOwnerAccount(v string) *UpdateFailoverTestJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetOwnerId(v int64) *UpdateFailoverTestJobRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetRegionId(v string) *UpdateFailoverTestJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetResourceId(v []*string) *UpdateFailoverTestJobRequest {
	s.ResourceId = v
	return s
}

func (s *UpdateFailoverTestJobRequest) SetResourceOwnerAccount(v string) *UpdateFailoverTestJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateFailoverTestJobRequest) Validate() error {
	return dara.Validate(s)
}
