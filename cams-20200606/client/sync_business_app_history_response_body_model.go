// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncBusinessAppHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncBusinessAppHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SyncBusinessAppHistoryResponseBody
	GetCode() *string
	SetMessage(v string) *SyncBusinessAppHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncBusinessAppHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncBusinessAppHistoryResponseBody
	GetSuccess() *bool
}

type SyncBusinessAppHistoryResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncBusinessAppHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncBusinessAppHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *SyncBusinessAppHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncBusinessAppHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncBusinessAppHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncBusinessAppHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncBusinessAppHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncBusinessAppHistoryResponseBody) SetAccessDeniedDetail(v string) *SyncBusinessAppHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncBusinessAppHistoryResponseBody) SetCode(v string) *SyncBusinessAppHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *SyncBusinessAppHistoryResponseBody) SetMessage(v string) *SyncBusinessAppHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *SyncBusinessAppHistoryResponseBody) SetRequestId(v string) *SyncBusinessAppHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncBusinessAppHistoryResponseBody) SetSuccess(v bool) *SyncBusinessAppHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *SyncBusinessAppHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}
