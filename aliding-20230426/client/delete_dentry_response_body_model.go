// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *DeleteDentryResponseBody
	GetAsync() *bool
	SetRequestId(v string) *DeleteDentryResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteDentryResponseBody
	GetTaskId() *string
	SetVendorRequestId(v string) *DeleteDentryResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteDentryResponseBody
	GetVendorType() *string
}

type DeleteDentryResponseBody struct {
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task123abc
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s DeleteDentryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDentryResponseBody) GetAsync() *bool {
	return s.Async
}

func (s *DeleteDentryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDentryResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteDentryResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteDentryResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteDentryResponseBody) SetAsync(v bool) *DeleteDentryResponseBody {
	s.Async = &v
	return s
}

func (s *DeleteDentryResponseBody) SetRequestId(v string) *DeleteDentryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDentryResponseBody) SetTaskId(v string) *DeleteDentryResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteDentryResponseBody) SetVendorRequestId(v string) *DeleteDentryResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteDentryResponseBody) SetVendorType(v string) *DeleteDentryResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteDentryResponseBody) Validate() error {
	return dara.Validate(s)
}
