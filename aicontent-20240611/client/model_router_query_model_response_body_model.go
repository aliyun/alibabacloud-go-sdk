// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelDTO) *ModelRouterQueryModelResponseBody
	GetData() *ModelDTO
	SetErrCode(v string) *ModelRouterQueryModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelResponseBody struct {
	// example:
	//
	// []
	Data *ModelDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelResponseBody) GetData() *ModelDTO {
	return s.Data
}

func (s *ModelRouterQueryModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelResponseBody) SetData(v *ModelDTO) *ModelRouterQueryModelResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelResponseBody) SetErrCode(v string) *ModelRouterQueryModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelResponseBody) SetErrMessage(v string) *ModelRouterQueryModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelResponseBody) SetRequestId(v string) *ModelRouterQueryModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelResponseBody) SetSuccess(v bool) *ModelRouterQueryModelResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
