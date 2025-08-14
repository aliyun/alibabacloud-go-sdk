// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateQueryVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateQueryVariableResponseBody
	GetResultObject() *bool
}

type CreateQueryVariableResponseBody struct {
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

func (s CreateQueryVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueryVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueryVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateQueryVariableResponseBody) SetRequestId(v string) *CreateQueryVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueryVariableResponseBody) SetResultObject(v bool) *CreateQueryVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateQueryVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
