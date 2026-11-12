// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePlansModificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstancePlansModificationRequest
	GetInstanceId() *string
	SetRegionId(v string) *ListInstancePlansModificationRequest
	GetRegionId() *string
}

type ListInstancePlansModificationRequest struct {
	// The ID of the Simple Application Server instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePlansModificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePlansModificationRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancePlansModificationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancePlansModificationRequest) SetInstanceId(v string) *ListInstancePlansModificationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePlansModificationRequest) SetRegionId(v string) *ListInstancePlansModificationRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePlansModificationRequest) Validate() error {
	return dara.Validate(s)
}
