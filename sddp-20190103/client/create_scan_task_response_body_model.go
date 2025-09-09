// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *CreateScanTaskResponseBody
	GetId() *int32
	SetRequestId(v string) *CreateScanTaskResponseBody
	GetRequestId() *string
}

type CreateScanTaskResponseBody struct {
	// The ID of the custom scan task.
	//
	// example:
	//
	// 100
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B1F2BB1F-04EC-5D36-B136-B4DE17FD8DE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScanTaskResponseBody) GetId() *int32 {
	return s.Id
}

func (s *CreateScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScanTaskResponseBody) SetId(v int32) *CreateScanTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScanTaskResponseBody) SetRequestId(v string) *CreateScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
