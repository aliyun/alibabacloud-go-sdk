// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCustVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateCustVariableResponseBody
	GetResultObject() *bool
}

type CreateCustVariableResponseBody struct {
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

func (s CreateCustVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateCustVariableResponseBody) SetRequestId(v string) *CreateCustVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustVariableResponseBody) SetResultObject(v bool) *CreateCustVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateCustVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
