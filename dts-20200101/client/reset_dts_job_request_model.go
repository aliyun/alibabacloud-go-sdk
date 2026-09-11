// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ResetDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ResetDtsJobRequest
	GetDtsJobId() *string
	SetRegionId(v string) *ResetDtsJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ResetDtsJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ResetDtsJobRequest
	GetSynchronizationDirection() *string
}

type ResetDtsJobRequest struct {
	// The synchronization or subscribe instance ID.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data synchronization or change tracking task.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzn4iqlbsm7hy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - You can set this parameter to **Reverse*	- to reset the reverse synchronization task only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s ResetDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ResetDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ResetDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ResetDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ResetDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ResetDtsJobRequest) SetDtsInstanceId(v string) *ResetDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ResetDtsJobRequest) SetDtsJobId(v string) *ResetDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *ResetDtsJobRequest) SetRegionId(v string) *ResetDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDtsJobRequest) SetResourceGroupId(v string) *ResetDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ResetDtsJobRequest) SetSynchronizationDirection(v string) *ResetDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ResetDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
