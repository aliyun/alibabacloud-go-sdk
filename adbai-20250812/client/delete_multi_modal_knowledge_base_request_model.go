// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiModalKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMultiModalKnowledgeBaseRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteMultiModalKnowledgeBaseRequest
	GetRegionId() *string
}

type DeleteMultiModalKnowledgeBaseRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMultiModalKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiModalKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiModalKnowledgeBaseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMultiModalKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMultiModalKnowledgeBaseRequest) SetDBClusterId(v string) *DeleteMultiModalKnowledgeBaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMultiModalKnowledgeBaseRequest) SetRegionId(v string) *DeleteMultiModalKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMultiModalKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
