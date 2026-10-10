// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterRenewApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterRenewApiKeyResponseBodyData) *ModelRouterRenewApiKeyResponseBody
	GetData() *ModelRouterRenewApiKeyResponseBodyData
	SetErrCode(v string) *ModelRouterRenewApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterRenewApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterRenewApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterRenewApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterRenewApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterRenewApiKeyResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterRenewApiKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message encoding.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterRenewApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterRenewApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterRenewApiKeyResponseBody) GetData() *ModelRouterRenewApiKeyResponseBodyData {
	return s.Data
}

func (s *ModelRouterRenewApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterRenewApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterRenewApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterRenewApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterRenewApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterRenewApiKeyResponseBody) SetData(v *ModelRouterRenewApiKeyResponseBodyData) *ModelRouterRenewApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) SetErrCode(v string) *ModelRouterRenewApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) SetErrMessage(v string) *ModelRouterRenewApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterRenewApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) SetRequestId(v string) *ModelRouterRenewApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) SetSuccess(v bool) *ModelRouterRenewApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterRenewApiKeyResponseBodyData struct {
	// The expiration time in RFC 3339 format. A value of null indicates that the API key remains valid indefinitely.
	//
	// example:
	//
	// 2027-01-01T00:00:00+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// API Key ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The enabled or disabled status. The status remains unchanged after renewal.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterRenewApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterRenewApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterRenewApiKeyResponseBodyData) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ModelRouterRenewApiKeyResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterRenewApiKeyResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModelRouterRenewApiKeyResponseBodyData) SetExpireAt(v string) *ModelRouterRenewApiKeyResponseBodyData {
	s.ExpireAt = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBodyData) SetId(v int64) *ModelRouterRenewApiKeyResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBodyData) SetStatus(v string) *ModelRouterRenewApiKeyResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
