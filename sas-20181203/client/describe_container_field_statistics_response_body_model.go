// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerFieldStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerGroupedFields(v *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) *DescribeContainerFieldStatisticsResponseBody
	GetContainerGroupedFields() *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields
	SetRequestId(v string) *DescribeContainerFieldStatisticsResponseBody
	GetRequestId() *string
}

type DescribeContainerFieldStatisticsResponseBody struct {
	// The statistical information about containers.
	ContainerGroupedFields *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields `json:"ContainerGroupedFields,omitempty" xml:"ContainerGroupedFields,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerFieldStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerFieldStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerFieldStatisticsResponseBody) GetContainerGroupedFields() *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	return s.ContainerGroupedFields
}

func (s *DescribeContainerFieldStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerFieldStatisticsResponseBody) SetContainerGroupedFields(v *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) *DescribeContainerFieldStatisticsResponseBody {
	s.ContainerGroupedFields = v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBody) SetRequestId(v string) *DescribeContainerFieldStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields struct {
	// The number of applications.
	//
	// example:
	//
	// 3
	AppCount *int32 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// The number of clusters.
	//
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The number of containers.
	//
	// example:
	//
	// 1
	ContainerCount *int32 `json:"ContainerCount,omitempty" xml:"ContainerCount,omitempty"`
	// The number of images.
	//
	// example:
	//
	// 3
	ImageCount *int32 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The number of namespaces.
	//
	// example:
	//
	// 3
	NamespaceCount *int32 `json:"NamespaceCount,omitempty" xml:"NamespaceCount,omitempty"`
	// The number of pods.
	//
	// example:
	//
	// 1
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// The number of the applications on which risks are detected.
	//
	// example:
	//
	// 1
	RiskAppCount *int32 `json:"RiskAppCount,omitempty" xml:"RiskAppCount,omitempty"`
	// The number of the clusters on which risks are detected.
	//
	// example:
	//
	// 1
	RiskClusterCount *int32 `json:"RiskClusterCount,omitempty" xml:"RiskClusterCount,omitempty"`
	// The number of the containers on which risks are detected.
	//
	// example:
	//
	// 1
	RiskContainerCount *int32 `json:"RiskContainerCount,omitempty" xml:"RiskContainerCount,omitempty"`
	// The number of the images on which risks are detected.
	//
	// example:
	//
	// 1
	RiskImageCount *int32 `json:"RiskImageCount,omitempty" xml:"RiskImageCount,omitempty"`
	// The number of the instances on which risks are detected.
	//
	// example:
	//
	// 3
	RiskInstanceCount *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
	// The number of the pods on which risks are detected.
	//
	// example:
	//
	// 2
	RiskPodCount *int32 `json:"RiskPodCount,omitempty" xml:"RiskPodCount,omitempty"`
}

func (s DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GoString() string {
	return s.String()
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetAppCount() *int32 {
	return s.AppCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetContainerCount() *int32 {
	return s.ContainerCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetNamespaceCount() *int32 {
	return s.NamespaceCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetPodCount() *int32 {
	return s.PodCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskAppCount() *int32 {
	return s.RiskAppCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskClusterCount() *int32 {
	return s.RiskClusterCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskContainerCount() *int32 {
	return s.RiskContainerCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskImageCount() *int32 {
	return s.RiskImageCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskInstanceCount() *int32 {
	return s.RiskInstanceCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) GetRiskPodCount() *int32 {
	return s.RiskPodCount
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetAppCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.AppCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetClusterCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.ClusterCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetContainerCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.ContainerCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetImageCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.ImageCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetInstanceCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.InstanceCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetNamespaceCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.NamespaceCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetPodCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.PodCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskAppCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskAppCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskClusterCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskClusterCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskContainerCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskContainerCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskImageCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskImageCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskInstanceCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) SetRiskPodCount(v int32) *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields {
	s.RiskPodCount = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponseBodyContainerGroupedFields) Validate() error {
	return dara.Validate(s)
}
