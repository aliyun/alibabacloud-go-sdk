// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetData() []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData
	SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
	GetSuccess() *bool
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetData() []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	return s.Data
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Success = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData struct {
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"FreeConcurrencyCount,omitempty" xml:"FreeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GetFreeConcurrencyCount() *int32 {
	return s.FreeConcurrencyCount
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GetFreeCount() *int32 {
	return s.FreeCount
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeConcurrencyCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceName(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
