// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *GetOrgOrWebOpenDocContentTaskIdResponseBody
	GetTaskId() *int64
	SetVendorRequestId(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody
	GetVendorType() *string
}

type GetOrgOrWebOpenDocContentTaskIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 158740210521
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) SetRequestId(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) SetTaskId(v int64) *GetOrgOrWebOpenDocContentTaskIdResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) SetVendorRequestId(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) SetVendorType(v string) *GetOrgOrWebOpenDocContentTaskIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponseBody) Validate() error {
	return dara.Validate(s)
}
