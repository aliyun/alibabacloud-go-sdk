// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueryVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQueryVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteQueryVariableResponseBody
	GetResultObject() *bool
}

type DeleteQueryVariableResponseBody struct {
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

func (s DeleteQueryVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueryVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueryVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQueryVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteQueryVariableResponseBody) SetRequestId(v string) *DeleteQueryVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueryVariableResponseBody) SetResultObject(v bool) *DeleteQueryVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteQueryVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
