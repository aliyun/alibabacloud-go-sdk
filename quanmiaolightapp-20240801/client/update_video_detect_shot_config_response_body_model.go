// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVideoDetectShotConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateVideoDetectShotConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVideoDetectShotConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVideoDetectShotConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVideoDetectShotConfigResponseBody
	GetSuccess() *bool
}

type UpdateVideoDetectShotConfigResponseBody struct {
	// example:
	//
	// xx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoDetectShotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoDetectShotConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVideoDetectShotConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVideoDetectShotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoDetectShotConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVideoDetectShotConfigResponseBody) SetCode(v string) *UpdateVideoDetectShotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponseBody) SetHttpStatusCode(v int32) *UpdateVideoDetectShotConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponseBody) SetMessage(v string) *UpdateVideoDetectShotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponseBody) SetRequestId(v string) *UpdateVideoDetectShotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponseBody) SetSuccess(v bool) *UpdateVideoDetectShotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVideoDetectShotConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
