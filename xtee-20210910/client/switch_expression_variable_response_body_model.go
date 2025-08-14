// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchExpressionVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchExpressionVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SwitchExpressionVariableResponseBody
	GetResultObject() *bool
}

type SwitchExpressionVariableResponseBody struct {
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

func (s SwitchExpressionVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchExpressionVariableResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchExpressionVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchExpressionVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SwitchExpressionVariableResponseBody) SetRequestId(v string) *SwitchExpressionVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchExpressionVariableResponseBody) SetResultObject(v bool) *SwitchExpressionVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SwitchExpressionVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
