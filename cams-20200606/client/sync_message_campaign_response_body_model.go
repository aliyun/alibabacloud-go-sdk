// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessageCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncMessageCampaignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SyncMessageCampaignResponseBody
	GetCode() *string
	SetMessage(v string) *SyncMessageCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncMessageCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncMessageCampaignResponseBody
	GetSuccess() *bool
}

type SyncMessageCampaignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 233**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncMessageCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncMessageCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMessageCampaignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncMessageCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncMessageCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncMessageCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncMessageCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncMessageCampaignResponseBody) SetAccessDeniedDetail(v string) *SyncMessageCampaignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncMessageCampaignResponseBody) SetCode(v string) *SyncMessageCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *SyncMessageCampaignResponseBody) SetMessage(v string) *SyncMessageCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *SyncMessageCampaignResponseBody) SetRequestId(v string) *SyncMessageCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMessageCampaignResponseBody) SetSuccess(v bool) *SyncMessageCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *SyncMessageCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
