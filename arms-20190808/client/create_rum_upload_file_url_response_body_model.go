// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumUploadFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateRumUploadFileUrlResponseBody
	GetCode() *int32
	SetData(v string) *CreateRumUploadFileUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateRumUploadFileUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateRumUploadFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRumUploadFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRumUploadFileUrlResponseBody
	GetSuccess() *bool
}

type CreateRumUploadFileUrlResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file upload URL.
	//
	// example:
	//
	// http://arms-rum-v2.oss-cn-hangzhou.aliyuncs.com/113197164xxxxx28/b590lhguqs%40f93xxxxxbf31d3/1.0.0-robots.txt?Expires=1713847079&OSSAccessKeyId=STS.NT6XvoxkyqA&Signature=6ptYX4OTjLMrsleTlA2t97
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRumUploadFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRumUploadFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRumUploadFileUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateRumUploadFileUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateRumUploadFileUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRumUploadFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRumUploadFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRumUploadFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRumUploadFileUrlResponseBody) SetCode(v int32) *CreateRumUploadFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) SetData(v string) *CreateRumUploadFileUrlResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) SetHttpStatusCode(v int32) *CreateRumUploadFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) SetMessage(v string) *CreateRumUploadFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) SetRequestId(v string) *CreateRumUploadFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) SetSuccess(v bool) *CreateRumUploadFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRumUploadFileUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
