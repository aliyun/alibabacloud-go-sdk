// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsageVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckUsageVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckUsageVariableResponseBody
	GetResultObject() *bool
}

type CheckUsageVariableResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether it was successful
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CheckUsageVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUsageVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsageVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUsageVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckUsageVariableResponseBody) SetRequestId(v string) *CheckUsageVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsageVariableResponseBody) SetResultObject(v bool) *CheckUsageVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckUsageVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
