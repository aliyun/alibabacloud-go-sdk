// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSelfBindVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSelfBindVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteSelfBindVariableResponseBody
	GetResultObject() *bool
}

type DeleteSelfBindVariableResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DeleteSelfBindVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSelfBindVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSelfBindVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSelfBindVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteSelfBindVariableResponseBody) SetRequestId(v string) *DeleteSelfBindVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSelfBindVariableResponseBody) SetResultObject(v bool) *DeleteSelfBindVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteSelfBindVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
