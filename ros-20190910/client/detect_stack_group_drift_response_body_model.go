// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackGroupDriftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *DetectStackGroupDriftResponseBody
	GetOperationId() *string
	SetRequestId(v string) *DetectStackGroupDriftResponseBody
	GetRequestId() *string
}

type DetectStackGroupDriftResponseBody struct {
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
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectStackGroupDriftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectStackGroupDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *DetectStackGroupDriftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectStackGroupDriftResponseBody) SetOperationId(v string) *DetectStackGroupDriftResponseBody {
	s.OperationId = &v
	return s
}

func (s *DetectStackGroupDriftResponseBody) SetRequestId(v string) *DetectStackGroupDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackGroupDriftResponseBody) Validate() error {
	return dara.Validate(s)
}
