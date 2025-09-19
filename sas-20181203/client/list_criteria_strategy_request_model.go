// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCriteriaStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListCriteriaStrategyRequest
	GetClusterId() *string
	SetImageName(v string) *ListCriteriaStrategyRequest
	GetImageName() *string
	SetLabel(v string) *ListCriteriaStrategyRequest
	GetLabel() *string
	SetNamespace(v string) *ListCriteriaStrategyRequest
	GetNamespace() *string
	SetStrategyName(v string) *ListCriteriaStrategyRequest
	GetStrategyName() *string
}

type ListCriteriaStrategyRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// c4af4fdf38a98496a9b63c2be5dae****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the image.
	//
	// >  You can call the [GetOpaClusterImageList](~~GetOpaClusterImageList~~) operation to query the names of images.
	//
	// example:
	//
	// testImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The tag that is added to the container.
	//
	// >  You can call the [GetOpaClusterLabelList](~~GetOpaClusterLabelList~~) operation to query the tags that are added to containers.
	//
	// example:
	//
	// testlabel
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The namespace of the cluster.
	//
	// >  You can call the [GetOpaClusterNamespaceList](~~GetOpaClusterNamespaceList~~) operation to query the namespaces of clusters.
	//
	// example:
	//
	// test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s ListCriteriaStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCriteriaStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListCriteriaStrategyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCriteriaStrategyRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListCriteriaStrategyRequest) GetLabel() *string {
	return s.Label
}

func (s *ListCriteriaStrategyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListCriteriaStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListCriteriaStrategyRequest) SetClusterId(v string) *ListCriteriaStrategyRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCriteriaStrategyRequest) SetImageName(v string) *ListCriteriaStrategyRequest {
	s.ImageName = &v
	return s
}

func (s *ListCriteriaStrategyRequest) SetLabel(v string) *ListCriteriaStrategyRequest {
	s.Label = &v
	return s
}

func (s *ListCriteriaStrategyRequest) SetNamespace(v string) *ListCriteriaStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *ListCriteriaStrategyRequest) SetStrategyName(v string) *ListCriteriaStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *ListCriteriaStrategyRequest) Validate() error {
	return dara.Validate(s)
}
