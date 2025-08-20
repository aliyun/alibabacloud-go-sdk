// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *DeleteStackInstancesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *DeleteStackInstancesResponseBody
	GetRequestId() *string
}

type DeleteStackInstancesResponseBody struct {
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

func (s DeleteStackInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *DeleteStackInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStackInstancesResponseBody) SetOperationId(v string) *DeleteStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *DeleteStackInstancesResponseBody) SetRequestId(v string) *DeleteStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStackInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
