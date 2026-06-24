// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePackagesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribePackagesResponseBodyItems) *DescribePackagesResponseBody
	GetItems() []*DescribePackagesResponseBodyItems
	SetPageSize(v int32) *DescribePackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePackagesResponseBody
	GetTotalCount() *int32
}

type DescribePackagesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The information about the data asset packages.
	Items []*DescribePackagesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePackagesResponseBody) GetItems() []*DescribePackagesResponseBodyItems {
	return s.Items
}

func (s *DescribePackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePackagesResponseBody) SetCurrentPage(v int32) *DescribePackagesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackagesResponseBody) SetItems(v []*DescribePackagesResponseBodyItems) *DescribePackagesResponseBody {
	s.Items = v
	return s
}

func (s *DescribePackagesResponseBody) SetPageSize(v int32) *DescribePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackagesResponseBody) SetRequestId(v string) *DescribePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackagesResponseBody) SetTotalCount(v int32) *DescribePackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackagesResponseBody) Validate() error {
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

type DescribePackagesResponseBodyItems struct {
	// The time when the data asset package was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The unique ID of the data asset package.
	//
	// example:
	//
	// 111111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the asset instance to which the data asset package belongs.
	//
	// example:
	//
	// 223453332
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the data asset package.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account of the data asset package owner.
	//
	// example:
	//
	// cou-2221
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the risk level for the data asset package.
	//
	// - **1**: N/A: No sensitive data is detected.
	//
	// - **2**: S1: Level 1 sensitive data.
	//
	// - **3**: S2: Level 2 sensitive data.
	//
	// - **4**: S3: Level 3 sensitive data.
	//
	// - **5**: S4: Level 4 sensitive data.
	//
	// example:
	//
	// 4
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the risk level for the package.
	//
	// example:
	//
	// S3
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// Indicates whether the data asset package contains sensitive data.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive data entries in the data asset package. For example, the total number of sensitive tables in MaxCompute.
	//
	// example:
	//
	// 123
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The total number of data entries in the data asset package. For example, the total number of tables in MaxCompute.
	//
	// example:
	//
	// 321
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePackagesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePackagesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribePackagesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribePackagesResponseBodyItems) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribePackagesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribePackagesResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribePackagesResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribePackagesResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribePackagesResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DescribePackagesResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribePackagesResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePackagesResponseBodyItems) SetCreationTime(v int64) *DescribePackagesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetId(v int64) *DescribePackagesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetInstanceId(v int64) *DescribePackagesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetName(v string) *DescribePackagesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetOwner(v string) *DescribePackagesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelId(v int64) *DescribePackagesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelName(v string) *DescribePackagesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitive(v bool) *DescribePackagesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitiveCount(v int32) *DescribePackagesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetTotalCount(v int32) *DescribePackagesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
