// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelDTO) *ModelRouterCreateModelResponseBody
	GetData() *ModelDTO
	SetErrCode(v string) *ModelRouterCreateModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateModelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateModelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateModelResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateModelResponseBody struct {
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

func (s ModelRouterCreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelResponseBody) GetData() *ModelDTO {
	return s.Data
}

func (s *ModelRouterCreateModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateModelResponseBody) SetData(v *ModelDTO) *ModelRouterCreateModelResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateModelResponseBody) SetErrCode(v string) *ModelRouterCreateModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateModelResponseBody) SetErrMessage(v string) *ModelRouterCreateModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateModelResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateModelResponseBody) SetRequestId(v string) *ModelRouterCreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateModelResponseBody) SetSuccess(v bool) *ModelRouterCreateModelResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
