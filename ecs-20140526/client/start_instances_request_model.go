// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StartInstancesRequest
	GetBatchOptimization() *string
	SetDryRun(v bool) *StartInstancesRequest
	GetDryRun() *bool
	SetInstanceId(v []*string) *StartInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartInstancesRequest
	GetResourceOwnerId() *int64
}

type StartInstancesRequest struct {
	// The batch operation pattern. Valid values:
	//
	// - AllTogether: In this pattern, if all instances are started, a success message is returned. If any instance fails validation, all instances fail to start and a failed message is returned.
	//
	// - SuccessFirst: In this pattern, each instance is started separately. The response contains the operation result for each instance.
	//
	// Default value: AllTogether.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding fault is returned. If the check succeeds, `DRYRUN.SUCCESS` is returned.
	//
	// > If the BatchOptimization parameter is set to `SuccessFirst`, the dry run with `DryRun=true` returns only `DRYRUN.SUCCESS`.
	//
	// - false: sends a Normal request. After the check succeeds, the instances are started.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of instance IDs. Valid values of the array length: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instances. You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StartInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartInstancesRequest) SetBatchOptimization(v string) *StartInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StartInstancesRequest) SetDryRun(v bool) *StartInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *StartInstancesRequest) SetInstanceId(v []*string) *StartInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StartInstancesRequest) SetOwnerAccount(v string) *StartInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetOwnerId(v int64) *StartInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StartInstancesRequest) SetRegionId(v string) *StartInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerAccount(v string) *StartInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerId(v int64) *StartInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartInstancesRequest) Validate() error {
	return dara.Validate(s)
}
