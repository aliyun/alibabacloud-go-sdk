// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeModelServicesResponseBodyItems) *DescribeModelServicesResponseBody
	GetItems() []*DescribeModelServicesResponseBodyItems
	SetPageNumber(v int32) *DescribeModelServicesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeModelServicesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeModelServicesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeModelServicesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeModelServicesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeModelServicesResponseBody struct {
	Items []*DescribeModelServicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeModelServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelServicesResponseBody) GetItems() []*DescribeModelServicesResponseBodyItems {
	return s.Items
}

func (s *DescribeModelServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelServicesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeModelServicesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelServicesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeModelServicesResponseBody) SetItems(v []*DescribeModelServicesResponseBodyItems) *DescribeModelServicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeModelServicesResponseBody) SetPageNumber(v int32) *DescribeModelServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelServicesResponseBody) SetPageRecordCount(v int32) *DescribeModelServicesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeModelServicesResponseBody) SetPageSize(v int32) *DescribeModelServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelServicesResponseBody) SetRequestId(v string) *DescribeModelServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelServicesResponseBody) SetTotalRecordCount(v int32) *DescribeModelServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeModelServicesResponseBody) Validate() error {
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

type DescribeModelServicesResponseBodyItems struct {
	// example:
	//
	// xxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// https://xxxxxx
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// example:
	//
	// 2026-03-31T14:40:48Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 10
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 10
	RequestCostPoints *string `json:"RequestCostPoints,omitempty" xml:"RequestCostPoints,omitempty"`
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeModelServicesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServicesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeModelServicesResponseBodyItems) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeModelServicesResponseBodyItems) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *DescribeModelServicesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeModelServicesResponseBodyItems) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *DescribeModelServicesResponseBodyItems) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *DescribeModelServicesResponseBodyItems) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeModelServicesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeModelServicesResponseBodyItems) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *DescribeModelServicesResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeModelServicesResponseBodyItems) GetRequestCostPoints() *string {
	return s.RequestCostPoints
}

func (s *DescribeModelServicesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelServicesResponseBodyItems) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeModelServicesResponseBodyItems) SetApiKey(v string) *DescribeModelServicesResponseBodyItems {
	s.ApiKey = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetBaseUrl(v string) *DescribeModelServicesResponseBodyItems {
	s.BaseUrl = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetGmtCreated(v string) *DescribeModelServicesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetInputCostPointsPerMillion(v string) *DescribeModelServicesResponseBodyItems {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetModelCategory(v string) *DescribeModelServicesResponseBodyItems {
	s.ModelCategory = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetModelServiceId(v string) *DescribeModelServicesResponseBodyItems {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetName(v string) *DescribeModelServicesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetOutputCostPointsPerMillion(v string) *DescribeModelServicesResponseBodyItems {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetProtocol(v string) *DescribeModelServicesResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetRequestCostPoints(v string) *DescribeModelServicesResponseBodyItems {
	s.RequestCostPoints = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetStatus(v string) *DescribeModelServicesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) SetVendor(v string) *DescribeModelServicesResponseBodyItems {
	s.Vendor = &v
	return s
}

func (s *DescribeModelServicesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
