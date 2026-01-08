// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteMessageCampaignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteMessageCampaignResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMessageCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMessageCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMessageCampaignResponseBody
	GetSuccess() *bool
}

type DeleteMessageCampaignResponseBody struct {
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
	// 239***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMessageCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageCampaignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteMessageCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMessageCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMessageCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMessageCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMessageCampaignResponseBody) SetAccessDeniedDetail(v string) *DeleteMessageCampaignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteMessageCampaignResponseBody) SetCode(v string) *DeleteMessageCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMessageCampaignResponseBody) SetMessage(v string) *DeleteMessageCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMessageCampaignResponseBody) SetRequestId(v string) *DeleteMessageCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageCampaignResponseBody) SetSuccess(v bool) *DeleteMessageCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMessageCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
