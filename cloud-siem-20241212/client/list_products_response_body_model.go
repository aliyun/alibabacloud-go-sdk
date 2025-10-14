// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsResponseBody
	GetNextToken() *string
	SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody
	GetProducts() []*ListProductsResponseBodyProducts
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsResponseBody
	GetTotalCount() *int32
}

type ListProductsResponseBody struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Products  []*ListProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsResponseBody) GetProducts() []*ListProductsResponseBodyProducts {
	return s.Products
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyProducts struct {
	// example:
	//
	// 2
	AbnormalDataIngestionCount *int32 `json:"AbnormalDataIngestionCount,omitempty" xml:"AbnormalDataIngestionCount,omitempty"`
	// example:
	//
	// 1733269771123
	ActiveTime *int64 `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	// example:
	//
	// true
	AllowAddDataIngestion *bool `json:"AllowAddDataIngestion,omitempty" xml:"AllowAddDataIngestion,omitempty"`
	// example:
	//
	// 1733269771123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// enabled
	DataIngestionStatus *bool `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// example:
	//
	// 1
	EnabledDataIngestionCount *int32 `json:"EnabledDataIngestionCount,omitempty" xml:"EnabledDataIngestionCount,omitempty"`
	// example:
	//
	// alibaba_cloud_sas
	ProductAlias *string `json:"ProductAlias,omitempty" xml:"ProductAlias,omitempty"`
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// sas
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// preset
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 10
	TotalDataIngestionCount *int32 `json:"TotalDataIngestionCount,omitempty" xml:"TotalDataIngestionCount,omitempty"`
	// example:
	//
	// 1733269771123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// vd-qlsw5eocx94w9
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
	// example:
	//
	// 111
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) GetAbnormalDataIngestionCount() *int32 {
	return s.AbnormalDataIngestionCount
}

func (s *ListProductsResponseBodyProducts) GetActiveTime() *int64 {
	return s.ActiveTime
}

func (s *ListProductsResponseBodyProducts) GetAllowAddDataIngestion() *bool {
	return s.AllowAddDataIngestion
}

func (s *ListProductsResponseBodyProducts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListProductsResponseBodyProducts) GetDataIngestionStatus() *bool {
	return s.DataIngestionStatus
}

func (s *ListProductsResponseBodyProducts) GetEnabledDataIngestionCount() *int32 {
	return s.EnabledDataIngestionCount
}

func (s *ListProductsResponseBodyProducts) GetProductAlias() *string {
	return s.ProductAlias
}

func (s *ListProductsResponseBodyProducts) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductsResponseBodyProducts) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsResponseBodyProducts) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductsResponseBodyProducts) GetTotalDataIngestionCount() *int32 {
	return s.TotalDataIngestionCount
}

func (s *ListProductsResponseBodyProducts) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListProductsResponseBodyProducts) GetVendorId() *string {
	return s.VendorId
}

func (s *ListProductsResponseBodyProducts) GetVendorName() *string {
	return s.VendorName
}

func (s *ListProductsResponseBodyProducts) SetAbnormalDataIngestionCount(v int32) *ListProductsResponseBodyProducts {
	s.AbnormalDataIngestionCount = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetActiveTime(v int64) *ListProductsResponseBodyProducts {
	s.ActiveTime = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetAllowAddDataIngestion(v bool) *ListProductsResponseBodyProducts {
	s.AllowAddDataIngestion = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetCreateTime(v int64) *ListProductsResponseBodyProducts {
	s.CreateTime = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetDataIngestionStatus(v bool) *ListProductsResponseBodyProducts {
	s.DataIngestionStatus = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetEnabledDataIngestionCount(v int32) *ListProductsResponseBodyProducts {
	s.EnabledDataIngestionCount = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductAlias(v string) *ListProductsResponseBodyProducts {
	s.ProductAlias = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductId(v string) *ListProductsResponseBodyProducts {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductName(v string) *ListProductsResponseBodyProducts {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductType(v string) *ListProductsResponseBodyProducts {
	s.ProductType = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetTotalDataIngestionCount(v int32) *ListProductsResponseBodyProducts {
	s.TotalDataIngestionCount = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetUpdateTime(v int64) *ListProductsResponseBodyProducts {
	s.UpdateTime = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetVendorId(v string) *ListProductsResponseBodyProducts {
	s.VendorId = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetVendorName(v string) *ListProductsResponseBodyProducts {
	s.VendorName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}
