// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeSpaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeSpaceId(v string) *DescribeKnowledgeSpaceAttributeRequest
	GetKnowledgeSpaceId() *string
	SetRegionId(v string) *DescribeKnowledgeSpaceAttributeRequest
	GetRegionId() *string
}

type DescribeKnowledgeSpaceAttributeRequest struct {
	// The unique identifier of the knowledge space.
	//
	// This parameter is required.
	//
	// example:
	//
	// pks-xxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeRequest) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeSpaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeSpaceAttributeRequest) SetKnowledgeSpaceId(v string) *DescribeKnowledgeSpaceAttributeRequest {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeRequest) SetRegionId(v string) *DescribeKnowledgeSpaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
