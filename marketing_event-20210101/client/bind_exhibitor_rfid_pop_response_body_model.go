// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindExhibitorRfidPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindExhibitorRfidPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *BindExhibitorRfidPopResponseBody
	GetData() *bool
	SetErrCode(v string) *BindExhibitorRfidPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *BindExhibitorRfidPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *BindExhibitorRfidPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *BindExhibitorRfidPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindExhibitorRfidPopResponseBody
	GetSuccess() *bool
}

type BindExhibitorRfidPopResponseBody struct {
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

func (s BindExhibitorRfidPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindExhibitorRfidPopResponseBody) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindExhibitorRfidPopResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindExhibitorRfidPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *BindExhibitorRfidPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *BindExhibitorRfidPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindExhibitorRfidPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindExhibitorRfidPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindExhibitorRfidPopResponseBody) SetAccessDeniedDetail(v string) *BindExhibitorRfidPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetData(v bool) *BindExhibitorRfidPopResponseBody {
	s.Data = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetErrCode(v string) *BindExhibitorRfidPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetErrMessage(v string) *BindExhibitorRfidPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetHttpStatusCode(v int32) *BindExhibitorRfidPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetRequestId(v string) *BindExhibitorRfidPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) SetSuccess(v bool) *BindExhibitorRfidPopResponseBody {
	s.Success = &v
	return s
}

func (s *BindExhibitorRfidPopResponseBody) Validate() error {
	return dara.Validate(s)
}
