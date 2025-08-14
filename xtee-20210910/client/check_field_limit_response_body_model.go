// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFieldLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckFieldLimitResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckFieldLimitResponseBody
	GetResultObject() *bool
}

type CheckFieldLimitResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the condition is met: -**true**: meets the condition-**false**: does not meet the condition
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CheckFieldLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckFieldLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFieldLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckFieldLimitResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckFieldLimitResponseBody) SetRequestId(v string) *CheckFieldLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFieldLimitResponseBody) SetResultObject(v bool) *CheckFieldLimitResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckFieldLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
