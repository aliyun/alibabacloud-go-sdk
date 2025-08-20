// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportStacksToStackGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *ImportStacksToStackGroupResponseBody
	GetOperationId() *string
	SetRequestId(v string) *ImportStacksToStackGroupResponseBody
	GetRequestId() *string
}

type ImportStacksToStackGroupResponseBody struct {
	// Operation ID.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportStacksToStackGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportStacksToStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ImportStacksToStackGroupResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *ImportStacksToStackGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportStacksToStackGroupResponseBody) SetOperationId(v string) *ImportStacksToStackGroupResponseBody {
	s.OperationId = &v
	return s
}

func (s *ImportStacksToStackGroupResponseBody) SetRequestId(v string) *ImportStacksToStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportStacksToStackGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
