// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogisticsOrderResult interface {
	dara.Model
	String() string
	GoString() string
	SetDataProvider(v string) *LogisticsOrderResult
	GetDataProvider() *string
	SetDataProviderTitle(v string) *LogisticsOrderResult
	GetDataProviderTitle() *string
	SetGoods(v []*Good) *LogisticsOrderResult
	GetGoods() []*Good
	SetLogisticsCompanyCode(v string) *LogisticsOrderResult
	GetLogisticsCompanyCode() *string
	SetLogisticsCompanyName(v string) *LogisticsOrderResult
	GetLogisticsCompanyName() *string
	SetLogisticsDetailList(v []*LogisticsDetail) *LogisticsOrderResult
	GetLogisticsDetailList() []*LogisticsDetail
	SetMailNo(v string) *LogisticsOrderResult
	GetMailNo() *string
}

type LogisticsOrderResult struct {
	DataProvider      *string `json:"dataProvider,omitempty" xml:"dataProvider,omitempty"`
	DataProviderTitle *string `json:"dataProviderTitle,omitempty" xml:"dataProviderTitle,omitempty"`
	Goods             []*Good `json:"goods,omitempty" xml:"goods,omitempty" type:"Repeated"`
	// example:
	//
	// SF
	LogisticsCompanyCode *string            `json:"logisticsCompanyCode,omitempty" xml:"logisticsCompanyCode,omitempty"`
	LogisticsCompanyName *string            `json:"logisticsCompanyName,omitempty" xml:"logisticsCompanyName,omitempty"`
	LogisticsDetailList  []*LogisticsDetail `json:"logisticsDetailList,omitempty" xml:"logisticsDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// SF234***2345
	MailNo *string `json:"mailNo,omitempty" xml:"mailNo,omitempty"`
}

func (s LogisticsOrderResult) String() string {
	return dara.Prettify(s)
}

func (s LogisticsOrderResult) GoString() string {
	return s.String()
}

func (s *LogisticsOrderResult) GetDataProvider() *string {
	return s.DataProvider
}

func (s *LogisticsOrderResult) GetDataProviderTitle() *string {
	return s.DataProviderTitle
}

func (s *LogisticsOrderResult) GetGoods() []*Good {
	return s.Goods
}

func (s *LogisticsOrderResult) GetLogisticsCompanyCode() *string {
	return s.LogisticsCompanyCode
}

func (s *LogisticsOrderResult) GetLogisticsCompanyName() *string {
	return s.LogisticsCompanyName
}

func (s *LogisticsOrderResult) GetLogisticsDetailList() []*LogisticsDetail {
	return s.LogisticsDetailList
}

func (s *LogisticsOrderResult) GetMailNo() *string {
	return s.MailNo
}

func (s *LogisticsOrderResult) SetDataProvider(v string) *LogisticsOrderResult {
	s.DataProvider = &v
	return s
}

func (s *LogisticsOrderResult) SetDataProviderTitle(v string) *LogisticsOrderResult {
	s.DataProviderTitle = &v
	return s
}

func (s *LogisticsOrderResult) SetGoods(v []*Good) *LogisticsOrderResult {
	s.Goods = v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsCompanyCode(v string) *LogisticsOrderResult {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsCompanyName(v string) *LogisticsOrderResult {
	s.LogisticsCompanyName = &v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsDetailList(v []*LogisticsDetail) *LogisticsOrderResult {
	s.LogisticsDetailList = v
	return s
}

func (s *LogisticsOrderResult) SetMailNo(v string) *LogisticsOrderResult {
	s.MailNo = &v
	return s
}

func (s *LogisticsOrderResult) Validate() error {
	if s.Goods != nil {
		for _, item := range s.Goods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogisticsDetailList != nil {
		for _, item := range s.LogisticsDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
