// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *UpdateStackGroupResponseBody
	GetOperationId() *string
	SetRequestId(v string) *UpdateStackGroupResponseBody
	GetRequestId() *string
}

type UpdateStackGroupResponseBody struct {
	// The ID of the operation.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStackGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *UpdateStackGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStackGroupResponseBody) SetOperationId(v string) *UpdateStackGroupResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateStackGroupResponseBody) SetRequestId(v string) *UpdateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
