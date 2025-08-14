// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressionVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExpressionVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteExpressionVariableResponseBody
	GetResultObject() *bool
}

type DeleteExpressionVariableResponseBody struct {
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

func (s DeleteExpressionVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressionVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressionVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressionVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteExpressionVariableResponseBody) SetRequestId(v string) *DeleteExpressionVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressionVariableResponseBody) SetResultObject(v bool) *DeleteExpressionVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteExpressionVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
