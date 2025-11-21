// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChangeResourceGroupResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ChangeResourceGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *ChangeResourceGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChangeResourceGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ChangeResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeResourceGroupResponseBody
	GetSuccess() *bool
}

type ChangeResourceGroupResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChangeResourceGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ChangeResourceGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChangeResourceGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChangeResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeResourceGroupResponseBody) SetAccessDeniedDetail(v string) *ChangeResourceGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v bool) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrCode(v string) *ChangeResourceGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrMessage(v string) *ChangeResourceGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetHttpStatusCode(v int32) *ChangeResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v bool) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
