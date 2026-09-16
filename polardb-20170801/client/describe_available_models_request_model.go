// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKubeType(v string) *DescribeAvailableModelsRequest
	GetKubeType() *string
	SetModelType(v string) *DescribeAvailableModelsRequest
	GetModelType() *string
	SetRegionId(v string) *DescribeAvailableModelsRequest
	GetRegionId() *string
}

type DescribeAvailableModelsRequest struct {
	// aideploy
	//
	// example:
	//
	// aideploy
	KubeType *string `json:"KubeType,omitempty" xml:"KubeType,omitempty"`
	// The model type. Valid values: custom or public. If this parameter is not specified, all models are returned.
	//
	// example:
	//
	// custom
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableModelsRequest) GetKubeType() *string {
	return s.KubeType
}

func (s *DescribeAvailableModelsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *DescribeAvailableModelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableModelsRequest) SetKubeType(v string) *DescribeAvailableModelsRequest {
	s.KubeType = &v
	return s
}

func (s *DescribeAvailableModelsRequest) SetModelType(v string) *DescribeAvailableModelsRequest {
	s.ModelType = &v
	return s
}

func (s *DescribeAvailableModelsRequest) SetRegionId(v string) *DescribeAvailableModelsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableModelsRequest) Validate() error {
	return dara.Validate(s)
}
