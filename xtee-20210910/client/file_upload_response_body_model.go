// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FileUploadResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *FileUploadResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *FileUploadResponseBody
	GetMessage() *string
	SetRequestId(v string) *FileUploadResponseBody
	GetRequestId() *string
	SetResultObject(v string) *FileUploadResponseBody
	GetResultObject() *string
}

type FileUploadResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Information returned by the API request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s FileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *FileUploadResponseBody) GetCode() *string {
	return s.Code
}

func (s *FileUploadResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *FileUploadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FileUploadResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *FileUploadResponseBody) SetCode(v string) *FileUploadResponseBody {
	s.Code = &v
	return s
}

func (s *FileUploadResponseBody) SetHttpStatusCode(v string) *FileUploadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FileUploadResponseBody) SetMessage(v string) *FileUploadResponseBody {
	s.Message = &v
	return s
}

func (s *FileUploadResponseBody) SetRequestId(v string) *FileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *FileUploadResponseBody) SetResultObject(v string) *FileUploadResponseBody {
	s.ResultObject = &v
	return s
}

func (s *FileUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
