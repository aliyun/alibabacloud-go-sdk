// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchQueryVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchQueryVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SwitchQueryVariableResponseBody
	GetResultObject() *bool
}

type SwitchQueryVariableResponseBody struct {
	// Request ID
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

func (s SwitchQueryVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchQueryVariableResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchQueryVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchQueryVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SwitchQueryVariableResponseBody) SetRequestId(v string) *SwitchQueryVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchQueryVariableResponseBody) SetResultObject(v bool) *SwitchQueryVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SwitchQueryVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
