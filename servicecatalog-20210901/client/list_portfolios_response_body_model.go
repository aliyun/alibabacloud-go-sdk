// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortfoliosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPortfoliosResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPortfoliosResponseBody
	GetPageSize() *int32
	SetPortfolioDetails(v []*ListPortfoliosResponseBodyPortfolioDetails) *ListPortfoliosResponseBody
	GetPortfolioDetails() []*ListPortfoliosResponseBodyPortfolioDetails
	SetRequestId(v string) *ListPortfoliosResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPortfoliosResponseBody
	GetTotalCount() *int32
}

type ListPortfoliosResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize         *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PortfolioDetails []*ListPortfoliosResponseBodyPortfolioDetails `json:"PortfolioDetails,omitempty" xml:"PortfolioDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPortfoliosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPortfoliosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortfoliosResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortfoliosResponseBody) GetPortfolioDetails() []*ListPortfoliosResponseBodyPortfolioDetails {
	return s.PortfolioDetails
}

func (s *ListPortfoliosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPortfoliosResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPortfoliosResponseBody) SetPageNumber(v int32) *ListPortfoliosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetPageSize(v int32) *ListPortfoliosResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetPortfolioDetails(v []*ListPortfoliosResponseBodyPortfolioDetails) *ListPortfoliosResponseBody {
	s.PortfolioDetails = v
	return s
}

func (s *ListPortfoliosResponseBody) SetRequestId(v string) *ListPortfoliosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetTotalCount(v int32) *ListPortfoliosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPortfoliosResponseBody) Validate() error {
	if s.PortfolioDetails != nil {
		for _, item := range s.PortfolioDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPortfoliosResponseBodyPortfolioDetails struct {
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 2022-04-12T06:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 产品组合描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:portfolio/port-bp1yt7582g****
	PortfolioArn *string `json:"PortfolioArn,omitempty" xml:"PortfolioArn,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// 代表资源名称的资源属性字段
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// 产品组合提供方
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s ListPortfoliosResponseBodyPortfolioDetails) String() string {
	return dara.Prettify(s)
}

func (s ListPortfoliosResponseBodyPortfolioDetails) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetDescription() *string {
	return s.Description
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetPortfolioArn() *string {
	return s.PortfolioArn
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetPortfolioName() *string {
	return s.PortfolioName
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetCreateTime(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.CreateTime = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetDescription(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.Description = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioArn(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioArn = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioId(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioName(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioName = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetProviderName(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.ProviderName = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) Validate() error {
	return dara.Validate(s)
}
