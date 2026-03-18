// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientDTO) *ModelRouterUpdateClientResponseBody
	GetData() *ClientDTO
	SetErrCode(v string) *ModelRouterUpdateClientResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateClientResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateClientResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateClientResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateClientResponseBody struct {
	// example:
	//
	// []
	Data *ClientDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterUpdateClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateClientResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateClientResponseBody) GetData() *ClientDTO {
	return s.Data
}

func (s *ModelRouterUpdateClientResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateClientResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateClientResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateClientResponseBody) SetData(v *ClientDTO) *ModelRouterUpdateClientResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) SetErrCode(v string) *ModelRouterUpdateClientResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) SetErrMessage(v string) *ModelRouterUpdateClientResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateClientResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) SetRequestId(v string) *ModelRouterUpdateClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) SetSuccess(v bool) *ModelRouterUpdateClientResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
