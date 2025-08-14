// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteCustVariableResponseBody
	GetResultObject() *bool
}

type DeleteCustVariableResponseBody struct {
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

func (s DeleteCustVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteCustVariableResponseBody) SetRequestId(v string) *DeleteCustVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustVariableResponseBody) SetResultObject(v bool) *DeleteCustVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteCustVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
