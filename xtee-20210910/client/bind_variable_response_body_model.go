// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *BindVariableResponseBody
	GetResultObject() *bool
}

type BindVariableResponseBody struct {
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

func (s BindVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindVariableResponseBody) GoString() string {
	return s.String()
}

func (s *BindVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *BindVariableResponseBody) SetRequestId(v string) *BindVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindVariableResponseBody) SetResultObject(v bool) *BindVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *BindVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
