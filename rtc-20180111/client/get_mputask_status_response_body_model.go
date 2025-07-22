// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMPUTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMPUTaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetMPUTaskStatusResponseBody
	GetStatus() *int32
}

type GetMPUTaskStatusResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMPUTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMPUTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMPUTaskStatusResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetMPUTaskStatusResponseBody) SetRequestId(v string) *GetMPUTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) SetStatus(v int32) *GetMPUTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
