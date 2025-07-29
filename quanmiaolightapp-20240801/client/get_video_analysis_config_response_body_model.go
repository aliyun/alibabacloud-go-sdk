// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoAnalysisConfigResponseBody
	GetCode() *string
	SetData(v *GetVideoAnalysisConfigResponseBodyData) *GetVideoAnalysisConfigResponseBody
	GetData() *GetVideoAnalysisConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetVideoAnalysisConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoAnalysisConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoAnalysisConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVideoAnalysisConfigResponseBody
	GetSuccess() *bool
}

type GetVideoAnalysisConfigResponseBody struct {
	// example:
	//
	// xx
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoAnalysisConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoAnalysisConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoAnalysisConfigResponseBody) GetData() *GetVideoAnalysisConfigResponseBodyData {
	return s.Data
}

func (s *GetVideoAnalysisConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoAnalysisConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoAnalysisConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoAnalysisConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoAnalysisConfigResponseBody) SetCode(v string) *GetVideoAnalysisConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetData(v *GetVideoAnalysisConfigResponseBodyData) *GetVideoAnalysisConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetHttpStatusCode(v int32) *GetVideoAnalysisConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetMessage(v string) *GetVideoAnalysisConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetRequestId(v string) *GetVideoAnalysisConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) SetSuccess(v bool) *GetVideoAnalysisConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisConfigResponseBodyData struct {
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s GetVideoAnalysisConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponseBodyData) GetAsyncConcurrency() *int32 {
	return s.AsyncConcurrency
}

func (s *GetVideoAnalysisConfigResponseBodyData) SetAsyncConcurrency(v int32) *GetVideoAnalysisConfigResponseBodyData {
	s.AsyncConcurrency = &v
	return s
}

func (s *GetVideoAnalysisConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
