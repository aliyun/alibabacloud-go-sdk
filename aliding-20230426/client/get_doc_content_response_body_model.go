// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDocContentResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *GetDocContentResponseBody
	GetTaskId() *int64
	SetVendorRequestId(v string) *GetDocContentResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDocContentResponseBody
	GetVendorType() *string
}

type GetDocContentResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task123abc
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

func (s GetDocContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocContentResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDocContentResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDocContentResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDocContentResponseBody) SetRequestId(v string) *GetDocContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocContentResponseBody) SetTaskId(v int64) *GetDocContentResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDocContentResponseBody) SetVendorRequestId(v string) *GetDocContentResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDocContentResponseBody) SetVendorType(v string) *GetDocContentResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDocContentResponseBody) Validate() error {
	return dara.Validate(s)
}
