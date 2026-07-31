// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelGroupDTO) *ModelRouterQueryModelGroupResponseBody
	GetData() *ModelGroupDTO
	SetErrCode(v string) *ModelRouterQueryModelGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryModelGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupResponseBody struct {
	// The data object.
	//
	// example:
	//
	// []
	Data *ModelGroupDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupResponseBody) GetData() *ModelGroupDTO {
	return s.Data
}

func (s *ModelRouterQueryModelGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupResponseBody) SetData(v *ModelGroupDTO) *ModelRouterQueryModelGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
