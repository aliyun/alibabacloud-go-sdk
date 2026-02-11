// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAllProdsResponseBodyData) *ListAllProdsResponseBody
	GetData() *ListAllProdsResponseBodyData
	SetRequestId(v string) *ListAllProdsResponseBody
	GetRequestId() *string
}

type ListAllProdsResponseBody struct {
	// The data returned.
	Data *ListAllProdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAllProdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllProdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBody) GetData() *ListAllProdsResponseBodyData {
	return s.Data
}

func (s *ListAllProdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllProdsResponseBody) SetData(v *ListAllProdsResponseBodyData) *ListAllProdsResponseBody {
	s.Data = v
	return s
}

func (s *ListAllProdsResponseBody) SetRequestId(v string) *ListAllProdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllProdsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllProdsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud services.
	//
	// example:
	//
	// 1
	ProdList []*ListAllProdsResponseBodyDataProdList `json:"ProdList,omitempty" xml:"ProdList,omitempty" type:"Repeated"`
	// The total number of logs.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAllProdsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllProdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAllProdsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAllProdsResponseBodyData) GetProdList() []*ListAllProdsResponseBodyDataProdList {
	return s.ProdList
}

func (s *ListAllProdsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAllProdsResponseBodyData) SetCurrentPage(v int32) *ListAllProdsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAllProdsResponseBodyData) SetPageSize(v int32) *ListAllProdsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAllProdsResponseBodyData) SetProdList(v []*ListAllProdsResponseBodyDataProdList) *ListAllProdsResponseBodyData {
	s.ProdList = v
	return s
}

func (s *ListAllProdsResponseBodyData) SetTotalCount(v int32) *ListAllProdsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAllProdsResponseBodyData) Validate() error {
	if s.ProdList != nil {
		for _, item := range s.ProdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllProdsResponseBodyDataProdList struct {
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- hcloud: Huawei Cloud.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The number of logs within the cloud service that are added to the threat analysis feature.
	//
	// example:
	//
	// 10
	ImportedLogCount *int32 `json:"ImportedLogCount,omitempty" xml:"ImportedLogCount,omitempty"`
	// The time when the logs within the cloud service were last added to the threat analysis feature.
	//
	// example:
	//
	// 2023-11-23 12:12:12
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// sas
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The total number of logs within the cloud service.
	//
	// example:
	//
	// 19
	TotalLogCount *int32 `json:"TotalLogCount,omitempty" xml:"TotalLogCount,omitempty"`
}

func (s ListAllProdsResponseBodyDataProdList) String() string {
	return dara.Prettify(s)
}

func (s ListAllProdsResponseBodyDataProdList) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBodyDataProdList) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListAllProdsResponseBodyDataProdList) GetImportedLogCount() *int32 {
	return s.ImportedLogCount
}

func (s *ListAllProdsResponseBodyDataProdList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAllProdsResponseBodyDataProdList) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListAllProdsResponseBodyDataProdList) GetTotalLogCount() *int32 {
	return s.TotalLogCount
}

func (s *ListAllProdsResponseBodyDataProdList) SetCloudCode(v string) *ListAllProdsResponseBodyDataProdList {
	s.CloudCode = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetImportedLogCount(v int32) *ListAllProdsResponseBodyDataProdList {
	s.ImportedLogCount = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetModifyTime(v string) *ListAllProdsResponseBodyDataProdList {
	s.ModifyTime = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetProdCode(v string) *ListAllProdsResponseBodyDataProdList {
	s.ProdCode = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetTotalLogCount(v int32) *ListAllProdsResponseBodyDataProdList {
	s.TotalLogCount = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) Validate() error {
	return dara.Validate(s)
}
