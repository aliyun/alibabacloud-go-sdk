// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelDTO) *ModelRouterUpdateModelResponseBody
	GetData() *ModelDTO
	SetErrCode(v string) *ModelRouterUpdateModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateModelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateModelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateModelResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateModelResponseBody struct {
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

func (s ModelRouterUpdateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelResponseBody) GetData() *ModelDTO {
	return s.Data
}

func (s *ModelRouterUpdateModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateModelResponseBody) SetData(v *ModelDTO) *ModelRouterUpdateModelResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) SetErrCode(v string) *ModelRouterUpdateModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) SetErrMessage(v string) *ModelRouterUpdateModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) SetRequestId(v string) *ModelRouterUpdateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) SetSuccess(v bool) *ModelRouterUpdateModelResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
