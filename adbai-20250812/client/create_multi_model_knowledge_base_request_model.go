// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiModelKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetDBClusterId() *string
	SetRegionId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetRegionId() *string
}

type CreateMultiModelKnowledgeBaseRequest struct {
	// The instance cluster ID.
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
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMultiModelKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiModelKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetDBClusterId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetRegionId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
