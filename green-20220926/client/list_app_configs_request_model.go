// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *ListAppConfigsRequest
	GetClassify() *string
	SetRegionId(v string) *ListAppConfigsRequest
	GetRegionId() *string
	SetResourceType(v string) *ListAppConfigsRequest
	GetResourceType() *string
}

type ListAppConfigsRequest struct {
	// The classification.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAppConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAppConfigsRequest) GetClassify() *string {
	return s.Classify
}

func (s *ListAppConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppConfigsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppConfigsRequest) SetClassify(v string) *ListAppConfigsRequest {
	s.Classify = &v
	return s
}

func (s *ListAppConfigsRequest) SetRegionId(v string) *ListAppConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAppConfigsRequest) SetResourceType(v string) *ListAppConfigsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppConfigsRequest) Validate() error {
	return dara.Validate(s)
}
