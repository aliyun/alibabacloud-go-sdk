// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateModelResponseBody
	GetResultObject() *bool
}

type CreateModelResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	//
	// example:
	//
	// True
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetResultObject(v bool) *CreateModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
