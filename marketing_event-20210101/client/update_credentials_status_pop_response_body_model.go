// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsStatusPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateCredentialsStatusPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateCredentialsStatusPopResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateCredentialsStatusPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateCredentialsStatusPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateCredentialsStatusPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateCredentialsStatusPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCredentialsStatusPopResponseBody
	GetSuccess() *bool
}

type UpdateCredentialsStatusPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCredentialsStatusPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsStatusPopResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateCredentialsStatusPopResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateCredentialsStatusPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateCredentialsStatusPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateCredentialsStatusPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCredentialsStatusPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialsStatusPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCredentialsStatusPopResponseBody) SetAccessDeniedDetail(v string) *UpdateCredentialsStatusPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetData(v bool) *UpdateCredentialsStatusPopResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetErrCode(v string) *UpdateCredentialsStatusPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetErrMessage(v string) *UpdateCredentialsStatusPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetHttpStatusCode(v int32) *UpdateCredentialsStatusPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetRequestId(v string) *UpdateCredentialsStatusPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) SetSuccess(v bool) *UpdateCredentialsStatusPopResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponseBody) Validate() error {
	return dara.Validate(s)
}
