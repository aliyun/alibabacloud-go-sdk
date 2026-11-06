// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *DescribeApplicationAvailableModelsRequest
	GetApplicationType() *string
	SetRegionId(v string) *DescribeApplicationAvailableModelsRequest
	GetRegionId() *string
}

type DescribeApplicationAvailableModelsRequest struct {
	// The application engine filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// mem0
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The region ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApplicationAvailableModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationAvailableModelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAvailableModelsRequest) SetApplicationType(v string) *DescribeApplicationAvailableModelsRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationAvailableModelsRequest) SetRegionId(v string) *DescribeApplicationAvailableModelsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAvailableModelsRequest) Validate() error {
	return dara.Validate(s)
}
