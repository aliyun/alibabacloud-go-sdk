// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProductListResponseBody
	GetCode() *string
	SetDisplayInfo(v *GetProductListResponseBodyDisplayInfo) *GetProductListResponseBody
	GetDisplayInfo() *GetProductListResponseBodyDisplayInfo
	SetMaxResults(v int32) *GetProductListResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *GetProductListResponseBody
	GetMessage() *string
	SetNextToken(v string) *GetProductListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetProductListResponseBody
	GetRequestId() *string
	SetTraceId(v string) *GetProductListResponseBody
	GetTraceId() *string
}

type GetProductListResponseBody struct {
	Code        *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	DisplayInfo *GetProductListResponseBodyDisplayInfo `json:"DisplayInfo,omitempty" xml:"DisplayInfo,omitempty" type:"Struct"`
	MaxResults  *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken   *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId     *string                                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetProductListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProductListResponseBody) GetDisplayInfo() *GetProductListResponseBodyDisplayInfo {
	return s.DisplayInfo
}

func (s *GetProductListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetProductListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProductListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetProductListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetProductListResponseBody) SetCode(v string) *GetProductListResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductListResponseBody) SetDisplayInfo(v *GetProductListResponseBodyDisplayInfo) *GetProductListResponseBody {
	s.DisplayInfo = v
	return s
}

func (s *GetProductListResponseBody) SetMaxResults(v int32) *GetProductListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetProductListResponseBody) SetMessage(v string) *GetProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductListResponseBody) SetNextToken(v string) *GetProductListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetProductListResponseBody) SetRequestId(v string) *GetProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductListResponseBody) SetTraceId(v string) *GetProductListResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetProductListResponseBody) Validate() error {
	if s.DisplayInfo != nil {
		if err := s.DisplayInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductListResponseBodyDisplayInfo struct {
	ProductList []*GetProductListResponseBodyDisplayInfoProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s GetProductListResponseBodyDisplayInfo) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyDisplayInfo) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyDisplayInfo) GetProductList() []*GetProductListResponseBodyDisplayInfoProductList {
	return s.ProductList
}

func (s *GetProductListResponseBodyDisplayInfo) SetProductList(v []*GetProductListResponseBodyDisplayInfoProductList) *GetProductListResponseBodyDisplayInfo {
	s.ProductList = v
	return s
}

func (s *GetProductListResponseBodyDisplayInfo) Validate() error {
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProductListResponseBodyDisplayInfoProductList struct {
	DisplayAttribute  *string                                                    `json:"DisplayAttribute,omitempty" xml:"DisplayAttribute,omitempty"`
	DisplayConfig     *string                                                    `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	DynamicAttribute  *string                                                    `json:"DynamicAttribute,omitempty" xml:"DynamicAttribute,omitempty"`
	IsEnable          *bool                                                      `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	ModificationLevel *int32                                                     `json:"ModificationLevel,omitempty" xml:"ModificationLevel,omitempty"`
	ProductCode       *string                                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductDesc       *string                                                    `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	ProductGroup      *string                                                    `json:"ProductGroup,omitempty" xml:"ProductGroup,omitempty"`
	ProductName       *string                                                    `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType       *string                                                    `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SkuList           []*GetProductListResponseBodyDisplayInfoProductListSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
}

func (s GetProductListResponseBodyDisplayInfoProductList) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyDisplayInfoProductList) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetDisplayAttribute() *string {
	return s.DisplayAttribute
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetDisplayConfig() *string {
	return s.DisplayConfig
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetDynamicAttribute() *string {
	return s.DynamicAttribute
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetModificationLevel() *int32 {
	return s.ModificationLevel
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetProductDesc() *string {
	return s.ProductDesc
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetProductGroup() *string {
	return s.ProductGroup
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetProductName() *string {
	return s.ProductName
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetProductType() *string {
	return s.ProductType
}

func (s *GetProductListResponseBodyDisplayInfoProductList) GetSkuList() []*GetProductListResponseBodyDisplayInfoProductListSkuList {
	return s.SkuList
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetDisplayAttribute(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.DisplayAttribute = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetDisplayConfig(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.DisplayConfig = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetDynamicAttribute(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.DynamicAttribute = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetIsEnable(v bool) *GetProductListResponseBodyDisplayInfoProductList {
	s.IsEnable = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetModificationLevel(v int32) *GetProductListResponseBodyDisplayInfoProductList {
	s.ModificationLevel = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetProductCode(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.ProductCode = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetProductDesc(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.ProductDesc = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetProductGroup(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.ProductGroup = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetProductName(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.ProductName = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetProductType(v string) *GetProductListResponseBodyDisplayInfoProductList {
	s.ProductType = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) SetSkuList(v []*GetProductListResponseBodyDisplayInfoProductListSkuList) *GetProductListResponseBodyDisplayInfoProductList {
	s.SkuList = v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductList) Validate() error {
	if s.SkuList != nil {
		for _, item := range s.SkuList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProductListResponseBodyDisplayInfoProductListSkuList struct {
	AppleSkuCode *string `json:"AppleSkuCode,omitempty" xml:"AppleSkuCode,omitempty"`
	Attribute    *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	PurchaseMode *string `json:"PurchaseMode,omitempty" xml:"PurchaseMode,omitempty"`
	SkuCode      *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	SkuDesc      *string `json:"SkuDesc,omitempty" xml:"SkuDesc,omitempty"`
	SkuName      *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
}

func (s GetProductListResponseBodyDisplayInfoProductListSkuList) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyDisplayInfoProductListSkuList) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetAppleSkuCode() *string {
	return s.AppleSkuCode
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetAttribute() *string {
	return s.Attribute
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetPurchaseMode() *string {
	return s.PurchaseMode
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetSkuCode() *string {
	return s.SkuCode
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetSkuDesc() *string {
	return s.SkuDesc
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) GetSkuName() *string {
	return s.SkuName
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetAppleSkuCode(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.AppleSkuCode = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetAttribute(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.Attribute = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetPurchaseMode(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.PurchaseMode = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetSkuCode(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.SkuCode = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetSkuDesc(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.SkuDesc = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) SetSkuName(v string) *GetProductListResponseBodyDisplayInfoProductListSkuList {
	s.SkuName = &v
	return s
}

func (s *GetProductListResponseBodyDisplayInfoProductListSkuList) Validate() error {
	return dara.Validate(s)
}
