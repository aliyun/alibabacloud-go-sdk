// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessengerSubscriptionTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncMessengerSubscriptionTokenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SyncMessengerSubscriptionTokenResponseBody
	GetCode() *string
	SetMessage(v string) *SyncMessengerSubscriptionTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncMessengerSubscriptionTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncMessengerSubscriptionTokenResponseBody
	GetSuccess() *bool
	SetTaskCode(v string) *SyncMessengerSubscriptionTokenResponseBody
	GetTaskCode() *string
}

type SyncMessengerSubscriptionTokenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 39ss**
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
}

func (s SyncMessengerSubscriptionTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncMessengerSubscriptionTokenResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncMessengerSubscriptionTokenResponseBody) GetTaskCode() *string {
	return s.TaskCode
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetAccessDeniedDetail(v string) *SyncMessengerSubscriptionTokenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetCode(v string) *SyncMessengerSubscriptionTokenResponseBody {
	s.Code = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetMessage(v string) *SyncMessengerSubscriptionTokenResponseBody {
	s.Message = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetRequestId(v string) *SyncMessengerSubscriptionTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetSuccess(v bool) *SyncMessengerSubscriptionTokenResponseBody {
	s.Success = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) SetTaskCode(v string) *SyncMessengerSubscriptionTokenResponseBody {
	s.TaskCode = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
