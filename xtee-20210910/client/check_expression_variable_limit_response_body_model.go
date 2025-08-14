// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpressionVariableLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckExpressionVariableLimitResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckExpressionVariableLimitResponseBody
	GetResultObject() *bool
}

type CheckExpressionVariableLimitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CheckExpressionVariableLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckExpressionVariableLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CheckExpressionVariableLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckExpressionVariableLimitResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckExpressionVariableLimitResponseBody) SetRequestId(v string) *CheckExpressionVariableLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckExpressionVariableLimitResponseBody) SetResultObject(v bool) *CheckExpressionVariableLimitResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckExpressionVariableLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
