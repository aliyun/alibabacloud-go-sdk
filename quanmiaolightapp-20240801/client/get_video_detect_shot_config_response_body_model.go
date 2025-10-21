// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetectShotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoDetectShotConfigResponseBody
	GetCode() *string
	SetData(v *GetVideoDetectShotConfigResponseBodyData) *GetVideoDetectShotConfigResponseBody
	GetData() *GetVideoDetectShotConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetVideoDetectShotConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoDetectShotConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoDetectShotConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVideoDetectShotConfigResponseBody
	GetSuccess() *bool
}

type GetVideoDetectShotConfigResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoDetectShotConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoDetectShotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoDetectShotConfigResponseBody) GetData() *GetVideoDetectShotConfigResponseBodyData {
	return s.Data
}

func (s *GetVideoDetectShotConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoDetectShotConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoDetectShotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoDetectShotConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoDetectShotConfigResponseBody) SetCode(v string) *GetVideoDetectShotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) SetData(v *GetVideoDetectShotConfigResponseBodyData) *GetVideoDetectShotConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) SetHttpStatusCode(v int32) *GetVideoDetectShotConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) SetMessage(v string) *GetVideoDetectShotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) SetRequestId(v string) *GetVideoDetectShotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) SetSuccess(v bool) *GetVideoDetectShotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoDetectShotConfigResponseBodyData struct {
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s GetVideoDetectShotConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotConfigResponseBodyData) GetAsyncConcurrency() *int32 {
	return s.AsyncConcurrency
}

func (s *GetVideoDetectShotConfigResponseBodyData) SetAsyncConcurrency(v int32) *GetVideoDetectShotConfigResponseBodyData {
	s.AsyncConcurrency = &v
	return s
}

func (s *GetVideoDetectShotConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
