// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *UpdateStackInstancesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *UpdateStackInstancesResponseBody
	GetRequestId() *string
}

type UpdateStackInstancesResponseBody struct {
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

func (s UpdateStackInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *UpdateStackInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStackInstancesResponseBody) SetOperationId(v string) *UpdateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateStackInstancesResponseBody) SetRequestId(v string) *UpdateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
