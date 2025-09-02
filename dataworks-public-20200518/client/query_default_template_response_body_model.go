// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDefaultTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *QueryDefaultTemplateResponseBody
	GetData() interface{}
	SetErrorCode(v string) *QueryDefaultTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryDefaultTemplateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *QueryDefaultTemplateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryDefaultTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDefaultTemplateResponseBody
	GetSuccess() *bool
}

type QueryDefaultTemplateResponseBody struct {
	// The returned data about the default data category and data sensitivity level template. The data is in the JSON array format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"gmtModified":1709022365000,"fileName":"default file","isDelete":false,"isDefaultTemplate":true}]
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9990030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// This parameter is required.
	//
	// example:
	//
	// Parameter error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDefaultTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDefaultTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDefaultTemplateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QueryDefaultTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDefaultTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDefaultTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryDefaultTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDefaultTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDefaultTemplateResponseBody) SetData(v interface{}) *QueryDefaultTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryDefaultTemplateResponseBody) SetErrorCode(v string) *QueryDefaultTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryDefaultTemplateResponseBody) SetErrorMessage(v string) *QueryDefaultTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDefaultTemplateResponseBody) SetHttpStatusCode(v int32) *QueryDefaultTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryDefaultTemplateResponseBody) SetRequestId(v string) *QueryDefaultTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDefaultTemplateResponseBody) SetSuccess(v bool) *QueryDefaultTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDefaultTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
