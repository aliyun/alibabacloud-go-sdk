// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *RebootInstancesRequest
	GetBatchOptimization() *string
	SetDryRun(v bool) *RebootInstancesRequest
	GetDryRun() *bool
	SetForceReboot(v bool) *RebootInstancesRequest
	GetForceReboot() *bool
	SetInstanceId(v []*string) *RebootInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *RebootInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebootInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RebootInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RebootInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebootInstancesRequest
	GetResourceOwnerId() *int64
}

type RebootInstancesRequest struct {
	// The batch operation mode. Valid values:
	//
	// - AllTogether: In this mode, a success message is returned if all instances are restarted. If any instance fails validation, all instances fail to restart and an error message is returned.
	//
	// - SuccessFirst: In this mode, each instance is restarted separately. The response includes the operation result for each instance.
	//
	// Default value: AllTogether.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run without restarting the instance. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check passes, `DRYRUN.SUCCESS` is returned.
	//
	// > If the BatchOptimization parameter is set to `SuccessFirst`, the dry run result for `DryRun=true` returns only `DRYRUN.SUCCESS`.
	//
	// - false: performs a dry run and sends the request. After the check passes, the instance is restarted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to force restart the instance. Valid values:
	//
	// -   true: forces a restart. This is equivalent to a power-off operation. Cached data that has not been written to storage devices is lost.
	//
	// -   false: performs a normal restart.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceReboot *bool `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	// The instance ID array. Array length: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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

func (s RebootInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *RebootInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RebootInstancesRequest) GetForceReboot() *bool {
	return s.ForceReboot
}

func (s *RebootInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *RebootInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebootInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebootInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebootInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebootInstancesRequest) SetBatchOptimization(v string) *RebootInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *RebootInstancesRequest) SetDryRun(v bool) *RebootInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *RebootInstancesRequest) SetForceReboot(v bool) *RebootInstancesRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootInstancesRequest) SetInstanceId(v []*string) *RebootInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *RebootInstancesRequest) SetOwnerAccount(v string) *RebootInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootInstancesRequest) SetOwnerId(v int64) *RebootInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootInstancesRequest) SetRegionId(v string) *RebootInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebootInstancesRequest) SetResourceOwnerAccount(v string) *RebootInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootInstancesRequest) SetResourceOwnerId(v int64) *RebootInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebootInstancesRequest) Validate() error {
	return dara.Validate(s)
}
