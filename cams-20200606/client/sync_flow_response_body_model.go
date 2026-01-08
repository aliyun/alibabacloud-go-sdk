// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *SyncFlowResponseBody
	GetCode() *int64
	SetMessage(v string) *SyncFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncFlowResponseBody
	GetSuccess() *bool
}

type SyncFlowResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 172D586A-5E9D-11C8-BED3-7E4A0***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SyncFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncFlowResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SyncFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncFlowResponseBody) SetAccessDeniedDetail(v string) *SyncFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncFlowResponseBody) SetCode(v int64) *SyncFlowResponseBody {
	s.Code = &v
	return s
}

func (s *SyncFlowResponseBody) SetMessage(v string) *SyncFlowResponseBody {
	s.Message = &v
	return s
}

func (s *SyncFlowResponseBody) SetRequestId(v string) *SyncFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncFlowResponseBody) SetSuccess(v bool) *SyncFlowResponseBody {
	s.Success = &v
	return s
}

func (s *SyncFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
