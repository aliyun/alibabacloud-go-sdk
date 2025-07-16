// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDocContentTakIdResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *GetDocContentTakIdResponseBody
	GetTaskId() *int64
	SetVendorRequestId(v string) *GetDocContentTakIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDocContentTakIdResponseBody
	GetVendorType() *string
}

type GetDocContentTakIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 72652830001
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

func (s GetDocContentTakIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocContentTakIdResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDocContentTakIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDocContentTakIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDocContentTakIdResponseBody) SetRequestId(v string) *GetDocContentTakIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocContentTakIdResponseBody) SetTaskId(v int64) *GetDocContentTakIdResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDocContentTakIdResponseBody) SetVendorRequestId(v string) *GetDocContentTakIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDocContentTakIdResponseBody) SetVendorType(v string) *GetDocContentTakIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDocContentTakIdResponseBody) Validate() error {
	return dara.Validate(s)
}
