// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerGroupedFieldDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeContainerGroupedFieldDetailRequest
	GetCriteria() *string
	SetGroupField(v string) *DescribeContainerGroupedFieldDetailRequest
	GetGroupField() *string
}

type DescribeContainerGroupedFieldDetailRequest struct {
	// The search conditions that are used to query assets. The value of this parameter is in the JSON format. Separate multiple search conditions with commas (,). Example: `[{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]`.
	//
	// >  Supported search conditions include the instance ID, instance name, virtual private cloud (VPC) ID, region, and public IP address. You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"clusterId\\",\\"value\\":\\"cfd26658431084c73a48dd97328ba8acf\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The filter condition for a grouping and aggregation query. Valid values:
	//
	// 	- **pod**
	//
	// 	- **appName**
	//
	// 	- **clusterId**
	//
	// 	- **namespace**
	//
	// 	- **image**
	//
	// 	- **containerScan**
	//
	// This parameter is required.
	//
	// example:
	//
	// pod
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
}

func (s DescribeContainerGroupedFieldDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerGroupedFieldDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupedFieldDetailRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeContainerGroupedFieldDetailRequest) GetGroupField() *string {
	return s.GroupField
}

func (s *DescribeContainerGroupedFieldDetailRequest) SetCriteria(v string) *DescribeContainerGroupedFieldDetailRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailRequest) SetGroupField(v string) *DescribeContainerGroupedFieldDetailRequest {
	s.GroupField = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailRequest) Validate() error {
	return dara.Validate(s)
}
