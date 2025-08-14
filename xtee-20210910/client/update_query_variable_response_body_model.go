// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueryVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateQueryVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateQueryVariableResponseBody
	GetResultObject() *bool
}

type UpdateQueryVariableResponseBody struct {
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

func (s UpdateQueryVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueryVariableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueryVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQueryVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateQueryVariableResponseBody) SetRequestId(v string) *UpdateQueryVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQueryVariableResponseBody) SetResultObject(v bool) *UpdateQueryVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateQueryVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
