// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *CreateStackInstancesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *CreateStackInstancesResponseBody
	GetRequestId() *string
}

type CreateStackInstancesResponseBody struct {
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

func (s CreateStackInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateStackInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStackInstancesResponseBody) SetOperationId(v string) *CreateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateStackInstancesResponseBody) SetRequestId(v string) *CreateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
