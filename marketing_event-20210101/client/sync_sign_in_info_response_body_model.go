// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSignInInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncSignInInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SyncSignInInfoResponseBody
	GetCode() *string
	SetData(v int32) *SyncSignInInfoResponseBody
	GetData() *int32
	SetHttpStatusCode(v string) *SyncSignInInfoResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SyncSignInInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncSignInInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncSignInInfoResponseBody
	GetSuccess() *bool
}

type SyncSignInInfoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode     *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncSignInInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncSignInInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncSignInInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncSignInInfoResponseBody) GetData() *int32 {
	return s.Data
}

func (s *SyncSignInInfoResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SyncSignInInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncSignInInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncSignInInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncSignInInfoResponseBody) SetAccessDeniedDetail(v string) *SyncSignInInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetCode(v string) *SyncSignInInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetData(v int32) *SyncSignInInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetHttpStatusCode(v string) *SyncSignInInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetMessage(v string) *SyncSignInInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetRequestId(v string) *SyncSignInInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncSignInInfoResponseBody) SetSuccess(v bool) *SyncSignInInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SyncSignInInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
