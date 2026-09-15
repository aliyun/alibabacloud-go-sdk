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
	// The conditions for searching assets. This parameter is in JSON format. Separate multiple conditions with commas (,). Example: `[{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]`.
	//
	// > You can search for assets by conditions such as instance ID, instance name, VPC ID, region, and public IP address. Call [DescribeCriteria](~~DescribeCriteria~~) to query the supported search conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"clusterId\\",\\"value\\":\\"cfd26658431084c73a48dd97328ba8acf\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The search item. Valid values:
	//
	// - **pod**: pod.
	//
	// - **appName**: application name.
	//
	// - **clusterId**: cluster ID.
	//
	// - **namespace**: namespace.
	//
	// - **image**: image.
	//
	// - **containerScan**: container scan.
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
