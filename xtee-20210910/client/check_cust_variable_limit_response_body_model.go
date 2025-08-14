// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCustVariableLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckCustVariableLimitResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckCustVariableLimitResponseBody
	GetResultObject() *bool
}

type CheckCustVariableLimitResponseBody struct {
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

func (s CheckCustVariableLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCustVariableLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCustVariableLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCustVariableLimitResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckCustVariableLimitResponseBody) SetRequestId(v string) *CheckCustVariableLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCustVariableLimitResponseBody) SetResultObject(v bool) *CheckCustVariableLimitResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckCustVariableLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
