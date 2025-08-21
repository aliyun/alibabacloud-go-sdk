// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v []*CopySnapshotResponseBodyAllocationId) *CopySnapshotResponseBody
	GetAllocationId() []*CopySnapshotResponseBodyAllocationId
	SetBizStatusCode(v string) *CopySnapshotResponseBody
	GetBizStatusCode() *string
	SetRequestId(v string) *CopySnapshotResponseBody
	GetRequestId() *string
	SetUnAllocationId(v []*CopySnapshotResponseBodyUnAllocationId) *CopySnapshotResponseBody
	GetUnAllocationId() []*CopySnapshotResponseBodyUnAllocationId
}

type CopySnapshotResponseBody struct {
	// The list of created snapshots.
	AllocationId []*CopySnapshotResponseBodyAllocationId `json:"AllocationId,omitempty" xml:"AllocationId,omitempty" type:"Repeated"`
	// The success status code.
	//
	// 	- **PartSuccess**: partially succeeded.
	//
	// 	- **AllSuccess**: all succeeded.
	//
	// example:
	//
	// AllSuccess
	BizStatusCode *string `json:"BizStatusCode,omitempty" xml:"BizStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA3758E0-8899-17D3-9526-5F62CF33A586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of nodes that are not created.
	UnAllocationId []*CopySnapshotResponseBodyUnAllocationId `json:"UnAllocationId,omitempty" xml:"UnAllocationId,omitempty" type:"Repeated"`
}

func (s CopySnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CopySnapshotResponseBody) GetAllocationId() []*CopySnapshotResponseBodyAllocationId {
	return s.AllocationId
}

func (s *CopySnapshotResponseBody) GetBizStatusCode() *string {
	return s.BizStatusCode
}

func (s *CopySnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopySnapshotResponseBody) GetUnAllocationId() []*CopySnapshotResponseBodyUnAllocationId {
	return s.UnAllocationId
}

func (s *CopySnapshotResponseBody) SetAllocationId(v []*CopySnapshotResponseBodyAllocationId) *CopySnapshotResponseBody {
	s.AllocationId = v
	return s
}

func (s *CopySnapshotResponseBody) SetBizStatusCode(v string) *CopySnapshotResponseBody {
	s.BizStatusCode = &v
	return s
}

func (s *CopySnapshotResponseBody) SetRequestId(v string) *CopySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySnapshotResponseBody) SetUnAllocationId(v []*CopySnapshotResponseBodyUnAllocationId) *CopySnapshotResponseBody {
	s.UnAllocationId = v
	return s
}

func (s *CopySnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}

type CopySnapshotResponseBodyAllocationId struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of the instances.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CopySnapshotResponseBodyAllocationId) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotResponseBodyAllocationId) GoString() string {
	return s.String()
}

func (s *CopySnapshotResponseBodyAllocationId) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CopySnapshotResponseBodyAllocationId) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CopySnapshotResponseBodyAllocationId) SetEnsRegionId(v string) *CopySnapshotResponseBodyAllocationId {
	s.EnsRegionId = &v
	return s
}

func (s *CopySnapshotResponseBodyAllocationId) SetInstanceId(v []*string) *CopySnapshotResponseBodyAllocationId {
	s.InstanceId = v
	return s
}

func (s *CopySnapshotResponseBodyAllocationId) Validate() error {
	return dara.Validate(s)
}

type CopySnapshotResponseBodyUnAllocationId struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-26
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s CopySnapshotResponseBodyUnAllocationId) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotResponseBodyUnAllocationId) GoString() string {
	return s.String()
}

func (s *CopySnapshotResponseBodyUnAllocationId) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CopySnapshotResponseBodyUnAllocationId) SetEnsRegionId(v string) *CopySnapshotResponseBodyUnAllocationId {
	s.EnsRegionId = &v
	return s
}

func (s *CopySnapshotResponseBodyUnAllocationId) Validate() error {
	return dara.Validate(s)
}
