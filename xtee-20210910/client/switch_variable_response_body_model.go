// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SwitchVariableResponseBody
	GetResultObject() *bool
}

type SwitchVariableResponseBody struct {
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

func (s SwitchVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchVariableResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SwitchVariableResponseBody) SetRequestId(v string) *SwitchVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchVariableResponseBody) SetResultObject(v bool) *SwitchVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SwitchVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
