// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuggestCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCna(v string) *SuggestCategoryRequest
	GetCna() *string
	SetContent(v string) *SuggestCategoryRequest
	GetContent() *string
	SetDistributionChannel(v string) *SuggestCategoryRequest
	GetDistributionChannel() *string
	SetReferrer(v string) *SuggestCategoryRequest
	GetReferrer() *string
	SetSubDistributionChannel(v string) *SuggestCategoryRequest
	GetSubDistributionChannel() *string
	SetTicketId(v string) *SuggestCategoryRequest
	GetTicketId() *string
	SetTopN(v int32) *SuggestCategoryRequest
	GetTopN() *int32
	SetUseHotSuggestCatchAll(v bool) *SuggestCategoryRequest
	GetUseHotSuggestCatchAll() *bool
	SetXGatewayExtendInfo(v string) *SuggestCategoryRequest
	GetXGatewayExtendInfo() *string
}

type SuggestCategoryRequest struct {
	Cna                    *string `json:"Cna,omitempty" xml:"Cna,omitempty"`
	Content                *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DistributionChannel    *string `json:"DistributionChannel,omitempty" xml:"DistributionChannel,omitempty"`
	Referrer               *string `json:"Referrer,omitempty" xml:"Referrer,omitempty"`
	SubDistributionChannel *string `json:"SubDistributionChannel,omitempty" xml:"SubDistributionChannel,omitempty"`
	TicketId               *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TopN                   *int32  `json:"TopN,omitempty" xml:"TopN,omitempty"`
	UseHotSuggestCatchAll  *bool   `json:"UseHotSuggestCatchAll,omitempty" xml:"UseHotSuggestCatchAll,omitempty"`
	XGatewayExtendInfo     *string `json:"XGatewayExtendInfo,omitempty" xml:"XGatewayExtendInfo,omitempty"`
}

func (s SuggestCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SuggestCategoryRequest) GoString() string {
	return s.String()
}

func (s *SuggestCategoryRequest) GetCna() *string {
	return s.Cna
}

func (s *SuggestCategoryRequest) GetContent() *string {
	return s.Content
}

func (s *SuggestCategoryRequest) GetDistributionChannel() *string {
	return s.DistributionChannel
}

func (s *SuggestCategoryRequest) GetReferrer() *string {
	return s.Referrer
}

func (s *SuggestCategoryRequest) GetSubDistributionChannel() *string {
	return s.SubDistributionChannel
}

func (s *SuggestCategoryRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *SuggestCategoryRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *SuggestCategoryRequest) GetUseHotSuggestCatchAll() *bool {
	return s.UseHotSuggestCatchAll
}

func (s *SuggestCategoryRequest) GetXGatewayExtendInfo() *string {
	return s.XGatewayExtendInfo
}

func (s *SuggestCategoryRequest) SetCna(v string) *SuggestCategoryRequest {
	s.Cna = &v
	return s
}

func (s *SuggestCategoryRequest) SetContent(v string) *SuggestCategoryRequest {
	s.Content = &v
	return s
}

func (s *SuggestCategoryRequest) SetDistributionChannel(v string) *SuggestCategoryRequest {
	s.DistributionChannel = &v
	return s
}

func (s *SuggestCategoryRequest) SetReferrer(v string) *SuggestCategoryRequest {
	s.Referrer = &v
	return s
}

func (s *SuggestCategoryRequest) SetSubDistributionChannel(v string) *SuggestCategoryRequest {
	s.SubDistributionChannel = &v
	return s
}

func (s *SuggestCategoryRequest) SetTicketId(v string) *SuggestCategoryRequest {
	s.TicketId = &v
	return s
}

func (s *SuggestCategoryRequest) SetTopN(v int32) *SuggestCategoryRequest {
	s.TopN = &v
	return s
}

func (s *SuggestCategoryRequest) SetUseHotSuggestCatchAll(v bool) *SuggestCategoryRequest {
	s.UseHotSuggestCatchAll = &v
	return s
}

func (s *SuggestCategoryRequest) SetXGatewayExtendInfo(v string) *SuggestCategoryRequest {
	s.XGatewayExtendInfo = &v
	return s
}

func (s *SuggestCategoryRequest) Validate() error {
	return dara.Validate(s)
}
