// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupField(v string) *DescribeContainerCriteriaRequest
	GetGroupField() *string
	SetValue(v string) *DescribeContainerCriteriaRequest
	GetValue() *string
}

type DescribeContainerCriteriaRequest struct {
	// The search field. Valid values:
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
	// example:
	//
	// clusterId
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The value of the search field. Fuzzy match is supported for application names, node names, namespaces, cluster names, public IP addresses, pod addresses, regions, pods, instance IDs, cluster IDs, and container IDs.
	//
	// example:
	//
	// cfb41a869c71e4678a97021582dd8a****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerCriteriaRequest) GetGroupField() *string {
	return s.GroupField
}

func (s *DescribeContainerCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeContainerCriteriaRequest) SetGroupField(v string) *DescribeContainerCriteriaRequest {
	s.GroupField = &v
	return s
}

func (s *DescribeContainerCriteriaRequest) SetValue(v string) *DescribeContainerCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeContainerCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
