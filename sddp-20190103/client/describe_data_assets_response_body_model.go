// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataAssetsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataAssetsResponseBodyItems) *DescribeDataAssetsResponseBody
	GetItems() []*DescribeDataAssetsResponseBodyItems
	SetPageSize(v int32) *DescribeDataAssetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataAssetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataAssetsResponseBody
	GetTotalCount() *int32
}

type DescribeDataAssetsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of data assets.
	Items []*DescribeDataAssetsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71064826-726F-4ADA-B879-05D8055476FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried data assets that contain sensitive data.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataAssetsResponseBody) GetItems() []*DescribeDataAssetsResponseBodyItems {
	return s.Items
}

func (s *DescribeDataAssetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataAssetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataAssetsResponseBody) SetCurrentPage(v int32) *DescribeDataAssetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetItems(v []*DescribeDataAssetsResponseBodyItems) *DescribeDataAssetsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetPageSize(v int32) *DescribeDataAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetRequestId(v string) *DescribeDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetTotalCount(v int32) *DescribeDataAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) Validate() error {
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

type DescribeDataAssetsResponseBodyItems struct {
	// The access control list (ACL) that controls the access permissions on the OSS bucket.
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **21**.
	//
	// example:
	//
	// private
	Acl *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	// The time when the data asset was created. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data type of the data asset.
	//
	// example:
	//
	// OSS_BUCKET
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The ID of the data asset.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sensitivity tag of the data. The value is fixed as **0**. **0**, **1**, **2**, or **3*	- is returned for this parameter only when you set the parameter **RangeId*	- to **1**.
	//
	// 	- **0**: unclassified
	//
	// 	- **1**: confidential
	//
	// 	- **2**: sensitive
	//
	// 	- **3**: highly sensitive
	//
	// example:
	//
	// 0
	Labelsec *bool `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	// The name of the data asset.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The key value of the OSS object.
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **22**.
	//
	// example:
	//
	// Internal
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The sensitivity level of the MaxCompute data asset. Valid values:
	//
	// 	- **S1**: low sensitivity level
	//
	// 	- **S2**: medium sensitivity level
	//
	// 	- **S3**: high sensitivity level
	//
	// 	- **S4**: highest sensitivity level
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **1**.
	//
	// example:
	//
	// S4
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The account that owns the data asset.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the service to which the data asset belongs.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Indicates whether the data protection mechanism is enabled for the data asset. The value is fixed as **false**. **true*	- or **false*	- is returned for this parameter only when you set the parameter **RangeId*	- to **1**.
	//
	// 	- **false**: The data protection mechanism is disabled.
	//
	// 	- **true**: The data protection mechanism is enabled. Only data inbound is supported. Data outbound is not supported.
	//
	// example:
	//
	// false
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
	// The sensitivity level of the data asset. A higher sensitivity level indicates that the identified data is more sensitive. Valid values:
	//
	// 	- **1**: No sensitive data is identified.
	//
	// 	- **2**: sensitive data at level 1.
	//
	// 	- **3**: sensitive data at level 2.
	//
	// 	- **3**: sensitive data at level 3.
	//
	// 	- **5**: sensitive data at level 4.
	//
	// 	- **6**: sensitive data at level 5.
	//
	// 	- **7**: sensitive data at level 6.
	//
	// 	- **8**: sensitive data at level 7.
	//
	// 	- **9**: sensitive data at level 8.
	//
	// 	- **10**: sensitive data at level 9.
	//
	// 	- **11**: sensitive data at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the data asset.
	//
	// example:
	//
	// Medium sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule that the data asset hits.
	//
	// example:
	//
	// \\*\\*\\	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the data asset contains sensitive data. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive data assets. For example, the value can be the total number of sensitive MaxCompute projects, packages, or tables, the total number of sensitive ApsaraDB RDS databases or tables, or the total number of sensitive OSS buckets or objects.
	//
	// example:
	//
	// 24
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The percentage of sensitive data in all data assets.
	//
	// example:
	//
	// 45%
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The total number of data assets. For example, the value can be the total number of MaxCompute projects, packages, or tables, the total number of ApsaraDB RDS databases or tables, or the total number of OSS buckets or objects.
	//
	// example:
	//
	// 432
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataAssetsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAssetsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponseBodyItems) GetAcl() *string {
	return s.Acl
}

func (s *DescribeDataAssetsResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeDataAssetsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataAssetsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeDataAssetsResponseBodyItems) GetLabelsec() *bool {
	return s.Labelsec
}

func (s *DescribeDataAssetsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeDataAssetsResponseBodyItems) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DescribeDataAssetsResponseBodyItems) GetOdpsRiskLevelName() *string {
	return s.OdpsRiskLevelName
}

func (s *DescribeDataAssetsResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribeDataAssetsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeDataAssetsResponseBodyItems) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeDataAssetsResponseBodyItems) GetProtection() *bool {
	return s.Protection
}

func (s *DescribeDataAssetsResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeDataAssetsResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeDataAssetsResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDataAssetsResponseBodyItems) GetSensitive() *bool {
	return s.Sensitive
}

func (s *DescribeDataAssetsResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribeDataAssetsResponseBodyItems) GetSensitiveRatio() *string {
	return s.SensitiveRatio
}

func (s *DescribeDataAssetsResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataAssetsResponseBodyItems) SetAcl(v string) *DescribeDataAssetsResponseBodyItems {
	s.Acl = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetCreationTime(v int64) *DescribeDataAssetsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetDataType(v string) *DescribeDataAssetsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetId(v string) *DescribeDataAssetsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetLabelsec(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetName(v string) *DescribeDataAssetsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetObjectKey(v string) *DescribeDataAssetsResponseBodyItems {
	s.ObjectKey = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOwner(v string) *DescribeDataAssetsResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductCode(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductId(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProtection(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRuleName(v string) *DescribeDataAssetsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitive(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveRatio(v string) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetTotalCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
