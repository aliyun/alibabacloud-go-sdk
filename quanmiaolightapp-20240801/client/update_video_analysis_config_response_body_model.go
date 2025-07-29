// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVideoAnalysisConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateVideoAnalysisConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVideoAnalysisConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVideoAnalysisConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVideoAnalysisConfigResponseBody
	GetSuccess() *bool
}

type UpdateVideoAnalysisConfigResponseBody struct {
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
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoAnalysisConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoAnalysisConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVideoAnalysisConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVideoAnalysisConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoAnalysisConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetCode(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetHttpStatusCode(v int32) *UpdateVideoAnalysisConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetMessage(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetRequestId(v string) *UpdateVideoAnalysisConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) SetSuccess(v bool) *UpdateVideoAnalysisConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
