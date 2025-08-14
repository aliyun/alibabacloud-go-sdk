// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressionVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressionVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModifyExpressionVariableResponseBody
	GetResultObject() *bool
}

type ModifyExpressionVariableResponseBody struct {
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

func (s ModifyExpressionVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressionVariableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressionVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressionVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModifyExpressionVariableResponseBody) SetRequestId(v string) *ModifyExpressionVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressionVariableResponseBody) SetResultObject(v bool) *ModifyExpressionVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModifyExpressionVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
