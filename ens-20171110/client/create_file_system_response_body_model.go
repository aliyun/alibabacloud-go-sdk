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
	SetBizStatusCode(v string) *CreateFileSystemResponseBody
	GetBizStatusCode() *string
	SetRequestId(v string) *CreateFileSystemResponseBody
	GetRequestId() *string
	SetUnAllocationId(v []*string) *CreateFileSystemResponseBody
	GetUnAllocationId() []*string
}

type CreateFileSystemResponseBody struct {
	// The information about the file system that was created.
	AllocationId []*string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty" type:"Repeated"`
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
