// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AliyunConsoleServiceInfoDTO) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetData() []*AliyunConsoleServiceInfoDTO
	SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
	GetSuccess() *bool
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleServiceInfoDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetData() []*AliyunConsoleServiceInfoDTO {
	return s.Data
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetData(v []*AliyunConsoleServiceInfoDTO) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Success = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) Validate() error {
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
