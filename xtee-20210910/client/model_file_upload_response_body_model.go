// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModelFileUploadResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModelFileUploadResponseBody
	GetResultObject() *bool
}

type ModelFileUploadResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s ModelFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *ModelFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelFileUploadResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModelFileUploadResponseBody) SetRequestId(v string) *ModelFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelFileUploadResponseBody) SetResultObject(v bool) *ModelFileUploadResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModelFileUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
