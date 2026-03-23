// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceMetricsResponseBody
	GetDBInstanceName() *string
	SetItems(v []*DescribeDBInstanceMetricsResponseBodyItems) *DescribeDBInstanceMetricsResponseBody
	GetItems() []*DescribeDBInstanceMetricsResponseBodyItems
	SetRequestId(v string) *DescribeDBInstanceMetricsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstanceMetricsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstanceMetricsResponseBody struct {
	DBInstanceName   *string                                       `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items            []*DescribeDBInstanceMetricsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                        `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceMetricsResponseBody) GetItems() []*DescribeDBInstanceMetricsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceMetricsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstanceMetricsResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceMetricsResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetItems(v []*DescribeDBInstanceMetricsResponseBodyItems) *DescribeDBInstanceMetricsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetRequestId(v string) *DescribeDBInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstanceMetricsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) Validate() error {
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

type DescribeDBInstanceMetricsResponseBodyItems struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Dimension       *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	GroupKey        *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	GroupKeyType    *string `json:"GroupKeyType,omitempty" xml:"GroupKeyType,omitempty"`
	Method          *string `json:"Method,omitempty" xml:"Method,omitempty"`
	MetricsKey      *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	MetricsKeyAlias *string `json:"MetricsKeyAlias,omitempty" xml:"MetricsKeyAlias,omitempty"`
	SortRule        *int32  `json:"SortRule,omitempty" xml:"SortRule,omitempty"`
	Unit            *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBInstanceMetricsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetGroupKey() *string {
	return s.GroupKey
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetGroupKeyType() *string {
	return s.GroupKeyType
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMethod() *string {
	return s.Method
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMetricsKey() *string {
	return s.MetricsKey
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMetricsKeyAlias() *string {
	return s.MetricsKeyAlias
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetSortRule() *int32 {
	return s.SortRule
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetDescription(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetDimension(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Dimension = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetGroupKey(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.GroupKey = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetGroupKeyType(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.GroupKeyType = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMethod(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Method = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMetricsKey(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.MetricsKey = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMetricsKeyAlias(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.MetricsKeyAlias = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetSortRule(v int32) *DescribeDBInstanceMetricsResponseBodyItems {
	s.SortRule = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetUnit(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Unit = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
