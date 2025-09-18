// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v []*string) *CreateFileSystemResponseBody
	GetAllocationId() []*string
	SetAllocationIds(v []*CreateFileSystemResponseBodyAllocationIds) *CreateFileSystemResponseBody
	GetAllocationIds() []*CreateFileSystemResponseBodyAllocationIds
	SetBizStatusCode(v string) *CreateFileSystemResponseBody
	GetBizStatusCode() *string
	SetRequestId(v string) *CreateFileSystemResponseBody
	GetRequestId() *string
	SetUnAllocationId(v []*string) *CreateFileSystemResponseBody
	GetUnAllocationId() []*string
}

type CreateFileSystemResponseBody struct {
	// The information about the file system that was created.
	AllocationId  []*string                                    `json:"AllocationId,omitempty" xml:"AllocationId,omitempty" type:"Repeated"`
	AllocationIds []*CreateFileSystemResponseBodyAllocationIds `json:"AllocationIds,omitempty" xml:"AllocationIds,omitempty" type:"Repeated"`
	// The status code for successful operations. Valid values:
	//
	// 	- PartSuccess: The operation is partially successful.
	//
	// 	- AllSuccess: The operation is successful.
	//
	// example:
	//
	// PartSuccess
	BizStatusCode *string `json:"BizStatusCode,omitempty" xml:"BizStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the file system that failed to be created.
	UnAllocationId []*string `json:"UnAllocationId,omitempty" xml:"UnAllocationId,omitempty" type:"Repeated"`
}

func (s CreateFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) GetAllocationId() []*string {
	return s.AllocationId
}

func (s *CreateFileSystemResponseBody) GetAllocationIds() []*CreateFileSystemResponseBodyAllocationIds {
	return s.AllocationIds
}

func (s *CreateFileSystemResponseBody) GetBizStatusCode() *string {
	return s.BizStatusCode
}

func (s *CreateFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileSystemResponseBody) GetUnAllocationId() []*string {
	return s.UnAllocationId
}

func (s *CreateFileSystemResponseBody) SetAllocationId(v []*string) *CreateFileSystemResponseBody {
	s.AllocationId = v
	return s
}

func (s *CreateFileSystemResponseBody) SetAllocationIds(v []*CreateFileSystemResponseBodyAllocationIds) *CreateFileSystemResponseBody {
	s.AllocationIds = v
	return s
}

func (s *CreateFileSystemResponseBody) SetBizStatusCode(v string) *CreateFileSystemResponseBody {
	s.BizStatusCode = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetUnAllocationId(v []*string) *CreateFileSystemResponseBody {
	s.UnAllocationId = v
	return s
}

func (s *CreateFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateFileSystemResponseBodyAllocationIds struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateFileSystemResponseBodyAllocationIds) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemResponseBodyAllocationIds) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBodyAllocationIds) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateFileSystemResponseBodyAllocationIds) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFileSystemResponseBodyAllocationIds) SetEnsRegionId(v string) *CreateFileSystemResponseBodyAllocationIds {
	s.EnsRegionId = &v
	return s
}

func (s *CreateFileSystemResponseBodyAllocationIds) SetInstanceId(v string) *CreateFileSystemResponseBodyAllocationIds {
	s.InstanceId = &v
	return s
}

func (s *CreateFileSystemResponseBodyAllocationIds) Validate() error {
	return dara.Validate(s)
}
