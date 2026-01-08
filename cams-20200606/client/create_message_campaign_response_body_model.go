// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateMessageCampaignResponseBody
	GetAccessDeniedDetail() *string
	SetCampaignId(v string) *CreateMessageCampaignResponseBody
	GetCampaignId() *string
	SetCode(v string) *CreateMessageCampaignResponseBody
	GetCode() *string
	SetMessage(v string) *CreateMessageCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMessageCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMessageCampaignResponseBody
	GetSuccess() *bool
}

type CreateMessageCampaignResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 39***
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
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
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMessageCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageCampaignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateMessageCampaignResponseBody) GetCampaignId() *string {
	return s.CampaignId
}

func (s *CreateMessageCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMessageCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMessageCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMessageCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMessageCampaignResponseBody) SetAccessDeniedDetail(v string) *CreateMessageCampaignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) SetCampaignId(v string) *CreateMessageCampaignResponseBody {
	s.CampaignId = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) SetCode(v string) *CreateMessageCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) SetMessage(v string) *CreateMessageCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) SetRequestId(v string) *CreateMessageCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) SetSuccess(v bool) *CreateMessageCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMessageCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
