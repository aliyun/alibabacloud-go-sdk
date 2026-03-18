// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientDTO) *ModelRouterCreateClientResponseBody
	GetData() *ClientDTO
	SetErrCode(v string) *ModelRouterCreateClientResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateClientResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateClientResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateClientResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateClientResponseBody struct {
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

func (s ModelRouterCreateClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateClientResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateClientResponseBody) GetData() *ClientDTO {
	return s.Data
}

func (s *ModelRouterCreateClientResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateClientResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateClientResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateClientResponseBody) SetData(v *ClientDTO) *ModelRouterCreateClientResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateClientResponseBody) SetErrCode(v string) *ModelRouterCreateClientResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateClientResponseBody) SetErrMessage(v string) *ModelRouterCreateClientResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateClientResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateClientResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateClientResponseBody) SetRequestId(v string) *ModelRouterCreateClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateClientResponseBody) SetSuccess(v bool) *ModelRouterCreateClientResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
