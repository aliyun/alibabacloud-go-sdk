// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ImportNameListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ImportNameListResponseBody
	GetResultObject() *bool
}

type ImportNameListResponseBody struct {
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

func (s ImportNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportNameListResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportNameListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ImportNameListResponseBody) SetRequestId(v string) *ImportNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNameListResponseBody) SetResultObject(v bool) *ImportNameListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ImportNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
