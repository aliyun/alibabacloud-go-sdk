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
	// The data source.
	//
	// example:
	//
	// Cainiao
	DataProvider *string `json:"dataProvider,omitempty" xml:"dataProvider,omitempty"`
	// The data text.
	//
	// example:
	//
	// Cainiao Guoguo
	DataProviderTitle *string `json:"dataProviderTitle,omitempty" xml:"dataProviderTitle,omitempty"`
	// The product details.
	Goods []*Good `json:"goods,omitempty" xml:"goods,omitempty" type:"Repeated"`
	// The logistics company code.
	//
	// > Valid values:
	//
	// > - ZTKY: China Railway Logistics
	//
	// > - POST: China Post
	//
	// > - DBKD: Deppon Express
	//
	// > - JT: J&T Express
	//
	// > - QFKD: Quanfeng Express
	//
	// > - EYB: China Post E-commerce Express EYB
	//
	// > - STO: STO Express
	//
	// > - SF: SF Express
	//
	// > - ZTO: ZTO Express
	//
	// > - YTO: YTO Express
	//
	// > - TTKDEX: Tiantian Express
	//
	// > - JDLEx: JD Express
	//
	// > - ETICKET: Electronic vouchers
	//
	// > - HTKY: Best Express
	//
	// > - SHQ: Huaqiang Logistics
	//
	// > - TAOBAO: Taobao Logistics
	//
	// > - YUNDA: Yunda Express
	//
	// > - ZJS: ZJS Express
	//
	// > - FEDEX: FedEx
	//
	// > - EMS: EMS
	//
	// > - POSTB: China Post Parcel
	//
	// > - OTHER: Other
	//
	// > - CNDJWL: Cainiao Large Item Logistics
	//
	// > - TN: Teneng
	//
	// > - ZMKM: Cainiao Express
	//
	// example:
	//
	// SF
	LogisticsCompanyCode *string `json:"logisticsCompanyCode,omitempty" xml:"logisticsCompanyCode,omitempty"`
	// The logistics company name.
	//
	// example:
	//
	// SF Express
	LogisticsCompanyName *string `json:"logisticsCompanyName,omitempty" xml:"logisticsCompanyName,omitempty"`
	// The logistics details.
	LogisticsDetailList []*LogisticsDetail `json:"logisticsDetailList,omitempty" xml:"logisticsDetailList,omitempty" type:"Repeated"`
	// The logistics tracking number.
	//
	// > The logistics tracking number for electronic vouchers is a fixed value: *.
	//
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
