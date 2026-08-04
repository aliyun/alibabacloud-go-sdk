// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureMemberBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterConfigureMemberBalanceResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterConfigureMemberBalanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterConfigureMemberBalanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterConfigureMemberBalanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterConfigureMemberBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterConfigureMemberBalanceResponseBody
	GetSuccess() *bool
}

type ModelRouterConfigureMemberBalanceResponseBody struct {
	// example:
	//
	// {}
	Data *ClientBalanceDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterConfigureMemberBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureMemberBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterConfigureMemberBalanceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetErrCode(v string) *ModelRouterConfigureMemberBalanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetErrMessage(v string) *ModelRouterConfigureMemberBalanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetHttpStatusCode(v int32) *ModelRouterConfigureMemberBalanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetRequestId(v string) *ModelRouterConfigureMemberBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) SetSuccess(v bool) *ModelRouterConfigureMemberBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
