// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *StartDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *StartDtsJobRequest
	GetDtsJobId() *string
	SetRegionId(v string) *StartDtsJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartDtsJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *StartDtsJobRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *StartDtsJobRequest
	GetZeroEtlJob() *bool
}

type StartDtsJobRequest struct {
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// > You can query the instance ID by invoking [Query the list of data migration instances and details of each instance](https://help.aliyun.com/document_detail/208139.html), [Query the list of change tracking instances and details of each instance](https://help.aliyun.com/document_detail/49442.html), or [Query the list of data synchronization instances and details of each instance](https://help.aliyun.com/document_detail/49454.html).
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2ilvoxlrdcby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward. This is the default value.
	//
	// - **Reverse**: reverse.
	//
	// > You can set this parameter to **Reverse*	- to start the reverse synchronization link only when the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Specifies whether the node is a zero-ETL node. Valid values:
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s StartDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobRequest) GoString() string {
	return s.String()
}

func (s *StartDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *StartDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *StartDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *StartDtsJobRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *StartDtsJobRequest) SetDtsInstanceId(v string) *StartDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *StartDtsJobRequest) SetDtsJobId(v string) *StartDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *StartDtsJobRequest) SetRegionId(v string) *StartDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartDtsJobRequest) SetResourceGroupId(v string) *StartDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartDtsJobRequest) SetSynchronizationDirection(v string) *StartDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StartDtsJobRequest) SetZeroEtlJob(v bool) *StartDtsJobRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *StartDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
