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

func (s *DescribeAvailableModelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableModelsRequest) SetKubeType(v string) *DescribeAvailableModelsRequest {
	s.KubeType = &v
	return s
}

func (s *DescribeAvailableModelsRequest) SetRegionId(v string) *DescribeAvailableModelsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableModelsRequest) Validate() error {
	return dara.Validate(s)
}
