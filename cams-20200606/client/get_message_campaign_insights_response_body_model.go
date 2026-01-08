// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCampaignInsightsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetMessageCampaignInsightsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetMessageCampaignInsightsResponseBody
	GetCode() *string
	SetData(v []*GetMessageCampaignInsightsResponseBodyData) *GetMessageCampaignInsightsResponseBody
	GetData() []*GetMessageCampaignInsightsResponseBodyData
	SetMessage(v string) *GetMessageCampaignInsightsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMessageCampaignInsightsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMessageCampaignInsightsResponseBody
	GetSuccess() *bool
}

type GetMessageCampaignInsightsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetMessageCampaignInsightsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// s39**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMessageCampaignInsightsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCampaignInsightsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCampaignInsightsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetMessageCampaignInsightsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMessageCampaignInsightsResponseBody) GetData() []*GetMessageCampaignInsightsResponseBodyData {
	return s.Data
}

func (s *GetMessageCampaignInsightsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMessageCampaignInsightsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageCampaignInsightsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMessageCampaignInsightsResponseBody) SetAccessDeniedDetail(v string) *GetMessageCampaignInsightsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) SetCode(v string) *GetMessageCampaignInsightsResponseBody {
	s.Code = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) SetData(v []*GetMessageCampaignInsightsResponseBodyData) *GetMessageCampaignInsightsResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) SetMessage(v string) *GetMessageCampaignInsightsResponseBody {
	s.Message = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) SetRequestId(v string) *GetMessageCampaignInsightsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) SetSuccess(v bool) *GetMessageCampaignInsightsResponseBody {
	s.Success = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMessageCampaignInsightsResponseBodyData struct {
	// example:
	//
	// 2025-07-01
	DateStart *string `json:"DateStart,omitempty" xml:"DateStart,omitempty"`
	// example:
	//
	// 2025-07-29
	DateStop *string `json:"DateStop,omitempty" xml:"DateStop,omitempty"`
	// example:
	//
	// 2755
	MarketingMessagesCostPerDelivered *string `json:"MarketingMessagesCostPerDelivered,omitempty" xml:"MarketingMessagesCostPerDelivered,omitempty"`
	// example:
	//
	// 268
	MarketingMessagesCostPerLinkBtnClick *string `json:"MarketingMessagesCostPerLinkBtnClick,omitempty" xml:"MarketingMessagesCostPerLinkBtnClick,omitempty"`
	// example:
	//
	// 2755
	MarketingMessagesDelivered *string `json:"MarketingMessagesDelivered,omitempty" xml:"MarketingMessagesDelivered,omitempty"`
	// example:
	//
	// 79.419238
	MarketingMessagesDeliveryRate *string `json:"MarketingMessagesDeliveryRate,omitempty" xml:"MarketingMessagesDeliveryRate,omitempty"`
	// example:
	//
	// 268
	MarketingMessagesLinkBtnClick *string `json:"MarketingMessagesLinkBtnClick,omitempty" xml:"MarketingMessagesLinkBtnClick,omitempty"`
	// example:
	//
	// 79.419238
	MarketingMessagesLinkBtnClickRate *string `json:"MarketingMessagesLinkBtnClickRate,omitempty" xml:"MarketingMessagesLinkBtnClickRate,omitempty"`
	// example:
	//
	// 79.419238
	MarketingMessagesReadRate *string `json:"MarketingMessagesReadRate,omitempty" xml:"MarketingMessagesReadRate,omitempty"`
	// example:
	//
	// 38.87
	MarketingMessagesSpend *string `json:"MarketingMessagesSpend,omitempty" xml:"MarketingMessagesSpend,omitempty"`
}

func (s GetMessageCampaignInsightsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCampaignInsightsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetDateStart() *string {
	return s.DateStart
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetDateStop() *string {
	return s.DateStop
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesCostPerDelivered() *string {
	return s.MarketingMessagesCostPerDelivered
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesCostPerLinkBtnClick() *string {
	return s.MarketingMessagesCostPerLinkBtnClick
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesDelivered() *string {
	return s.MarketingMessagesDelivered
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesDeliveryRate() *string {
	return s.MarketingMessagesDeliveryRate
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesLinkBtnClick() *string {
	return s.MarketingMessagesLinkBtnClick
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesLinkBtnClickRate() *string {
	return s.MarketingMessagesLinkBtnClickRate
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesReadRate() *string {
	return s.MarketingMessagesReadRate
}

func (s *GetMessageCampaignInsightsResponseBodyData) GetMarketingMessagesSpend() *string {
	return s.MarketingMessagesSpend
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetDateStart(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.DateStart = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetDateStop(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.DateStop = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesCostPerDelivered(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesCostPerDelivered = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesCostPerLinkBtnClick(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesCostPerLinkBtnClick = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesDelivered(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesDelivered = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesDeliveryRate(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesDeliveryRate = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesLinkBtnClick(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesLinkBtnClick = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesLinkBtnClickRate(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesLinkBtnClickRate = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesReadRate(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesReadRate = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) SetMarketingMessagesSpend(v string) *GetMessageCampaignInsightsResponseBodyData {
	s.MarketingMessagesSpend = &v
	return s
}

func (s *GetMessageCampaignInsightsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
