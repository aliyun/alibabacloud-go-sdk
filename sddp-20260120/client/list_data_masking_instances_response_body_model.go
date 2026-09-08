// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDataMaskingInstancesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListDataMaskingInstancesResponseBodyItems) *ListDataMaskingInstancesResponseBody
	GetItems() []*ListDataMaskingInstancesResponseBodyItems
	SetPageSize(v int32) *ListDataMaskingInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataMaskingInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataMaskingInstancesResponseBody
	GetTotalCount() *int32
}

type ListDataMaskingInstancesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListDataMaskingInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataMaskingInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataMaskingInstancesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataMaskingInstancesResponseBody) GetItems() []*ListDataMaskingInstancesResponseBodyItems {
	return s.Items
}

func (s *ListDataMaskingInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataMaskingInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataMaskingInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataMaskingInstancesResponseBody) SetCurrentPage(v int32) *ListDataMaskingInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBody) SetItems(v []*ListDataMaskingInstancesResponseBodyItems) *ListDataMaskingInstancesResponseBody {
	s.Items = v
	return s
}

func (s *ListDataMaskingInstancesResponseBody) SetPageSize(v int32) *ListDataMaskingInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBody) SetRequestId(v string) *ListDataMaskingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBody) SetTotalCount(v int32) *ListDataMaskingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataMaskingInstancesResponseBodyItems struct {
	// example:
	//
	// AES_256_GCM
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2
	FullAccessAccountCount *int32 `json:"FullAccessAccountCount,omitempty" xml:"FullAccessAccountCount,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 华北 3（张家口）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListDataMaskingInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetFullAccessAccountCount() *int32 {
	return s.FullAccessAccountCount
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataMaskingInstancesResponseBodyItems) GetRegionName() *string {
	return s.RegionName
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetEncryptionAlgorithm(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetEncryptionKeyId(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.EncryptionKeyId = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetEngineType(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetFullAccessAccountCount(v int32) *ListDataMaskingInstancesResponseBodyItems {
	s.FullAccessAccountCount = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetInstanceId(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetProductCode(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetProductId(v int64) *ListDataMaskingInstancesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetRegionId(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) SetRegionName(v string) *ListDataMaskingInstancesResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *ListDataMaskingInstancesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
