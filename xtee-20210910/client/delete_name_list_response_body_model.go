// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNameListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteNameListResponseBody
	GetResultObject() *bool
}

type DeleteNameListResponseBody struct {
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

func (s DeleteNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNameListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteNameListResponseBody) SetRequestId(v string) *DeleteNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNameListResponseBody) SetResultObject(v bool) *DeleteNameListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
