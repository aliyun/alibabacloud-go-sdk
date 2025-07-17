// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumUploadFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteRumUploadFileResponseBody
	GetCode() *int32
	SetData(v string) *DeleteRumUploadFileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteRumUploadFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRumUploadFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRumUploadFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRumUploadFileResponseBody
	GetSuccess() *bool
}

type DeleteRumUploadFileResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the file is deleted. If the file is deleted, "success" is returned.
	//
	// example:
	//
	// success
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
	// Internal error, please contact the administrator.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s DeleteRumUploadFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRumUploadFileResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteRumUploadFileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteRumUploadFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRumUploadFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRumUploadFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRumUploadFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRumUploadFileResponseBody) SetCode(v int32) *DeleteRumUploadFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) SetData(v string) *DeleteRumUploadFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) SetHttpStatusCode(v int32) *DeleteRumUploadFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) SetMessage(v string) *DeleteRumUploadFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) SetRequestId(v string) *DeleteRumUploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) SetSuccess(v bool) *DeleteRumUploadFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRumUploadFileResponseBody) Validate() error {
	return dara.Validate(s)
}
