// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioDetail(v *GetPortfolioResponseBodyPortfolioDetail) *GetPortfolioResponseBody
	GetPortfolioDetail() *GetPortfolioResponseBodyPortfolioDetail
	SetRequestId(v string) *GetPortfolioResponseBody
	GetRequestId() *string
	SetTagOptions(v []*GetPortfolioResponseBodyTagOptions) *GetPortfolioResponseBody
	GetTagOptions() []*GetPortfolioResponseBodyTagOptions
}

type GetPortfolioResponseBody struct {
	// The details of the product portfolio.
	PortfolioDetail *GetPortfolioResponseBodyPortfolioDetail `json:"PortfolioDetail,omitempty" xml:"PortfolioDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag options associated with the service portfolio.
	TagOptions []*GetPortfolioResponseBodyTagOptions `json:"TagOptions,omitempty" xml:"TagOptions,omitempty" type:"Repeated"`
}

func (s GetPortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBody) GetPortfolioDetail() *GetPortfolioResponseBodyPortfolioDetail {
	return s.PortfolioDetail
}

func (s *GetPortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPortfolioResponseBody) GetTagOptions() []*GetPortfolioResponseBodyTagOptions {
	return s.TagOptions
}

func (s *GetPortfolioResponseBody) SetPortfolioDetail(v *GetPortfolioResponseBodyPortfolioDetail) *GetPortfolioResponseBody {
	s.PortfolioDetail = v
	return s
}

func (s *GetPortfolioResponseBody) SetRequestId(v string) *GetPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPortfolioResponseBody) SetTagOptions(v []*GetPortfolioResponseBodyTagOptions) *GetPortfolioResponseBody {
	s.TagOptions = v
	return s
}

func (s *GetPortfolioResponseBody) Validate() error {
	if s.PortfolioDetail != nil {
		if err := s.PortfolioDetail.Validate(); err != nil {
			return err
		}
	}
	if s.TagOptions != nil {
		for _, item := range s.TagOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPortfolioResponseBodyPortfolioDetail struct {
	// The time when the product portfolio was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-12T06:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product portfolio.
	//
	// example:
	//
	// The description of the product portfolio.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product portfolio.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:portfolio/port-bp1yt7582g****
	PortfolioArn *string `json:"PortfolioArn,omitempty" xml:"PortfolioArn,omitempty"`
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The name of the product portfolio.
	//
	// example:
	//
	// DEMO-IT services
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// The provider of the product portfolio.
	//
	// example:
	//
	// IT team
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetPortfolioResponseBodyPortfolioDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPortfolioResponseBodyPortfolioDetail) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetDescription() *string {
	return s.Description
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetPortfolioArn() *string {
	return s.PortfolioArn
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetPortfolioName() *string {
	return s.PortfolioName
}

func (s *GetPortfolioResponseBodyPortfolioDetail) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetCreateTime(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.CreateTime = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetDescription(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.Description = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioArn(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioArn = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioId(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioName(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioName = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetProviderName(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.ProviderName = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) Validate() error {
	return dara.Validate(s)
}

type GetPortfolioResponseBodyTagOptions struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the owner of the tag option.
	//
	// example:
	//
	// 133413081827****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	//
	// example:
	//
	// tag-bp1q65xd3j****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPortfolioResponseBodyTagOptions) String() string {
	return dara.Prettify(s)
}

func (s GetPortfolioResponseBodyTagOptions) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBodyTagOptions) GetActive() *bool {
	return s.Active
}

func (s *GetPortfolioResponseBodyTagOptions) GetKey() *string {
	return s.Key
}

func (s *GetPortfolioResponseBodyTagOptions) GetOwner() *string {
	return s.Owner
}

func (s *GetPortfolioResponseBodyTagOptions) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *GetPortfolioResponseBodyTagOptions) GetValue() *string {
	return s.Value
}

func (s *GetPortfolioResponseBodyTagOptions) SetActive(v bool) *GetPortfolioResponseBodyTagOptions {
	s.Active = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetKey(v string) *GetPortfolioResponseBodyTagOptions {
	s.Key = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetOwner(v string) *GetPortfolioResponseBodyTagOptions {
	s.Owner = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetTagOptionId(v string) *GetPortfolioResponseBodyTagOptions {
	s.TagOptionId = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetValue(v string) *GetPortfolioResponseBodyTagOptions {
	s.Value = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) Validate() error {
	return dara.Validate(s)
}
