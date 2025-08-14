// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressionVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateExpressionVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateExpressionVariableResponseBody
	GetResultObject() *bool
}

type CreateExpressionVariableResponseBody struct {
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

func (s CreateExpressionVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressionVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressionVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressionVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateExpressionVariableResponseBody) SetRequestId(v string) *CreateExpressionVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressionVariableResponseBody) SetResultObject(v bool) *CreateExpressionVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateExpressionVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
