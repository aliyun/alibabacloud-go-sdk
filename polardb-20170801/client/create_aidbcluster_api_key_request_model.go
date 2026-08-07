// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAIDBClusterApiKeyRequest
	GetDescription() *string
	SetModelSpaceName(v string) *CreateAIDBClusterApiKeyRequest
	GetModelSpaceName() *string
	SetRegionId(v string) *CreateAIDBClusterApiKeyRequest
	GetRegionId() *string
}

type CreateAIDBClusterApiKeyRequest struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// pms-xxx
	ModelSpaceName *string `json:"ModelSpaceName,omitempty" xml:"ModelSpaceName,omitempty"`
	// The region ID.
	//
	// > 	- You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions of all clusters under the specified account.
	//
	// > 	- If you leave this parameter empty, scheduled tasks across all regions under the current account are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAIDBClusterApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAIDBClusterApiKeyRequest) GetModelSpaceName() *string {
	return s.ModelSpaceName
}

func (s *CreateAIDBClusterApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAIDBClusterApiKeyRequest) SetDescription(v string) *CreateAIDBClusterApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateAIDBClusterApiKeyRequest) SetModelSpaceName(v string) *CreateAIDBClusterApiKeyRequest {
	s.ModelSpaceName = &v
	return s
}

func (s *CreateAIDBClusterApiKeyRequest) SetRegionId(v string) *CreateAIDBClusterApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAIDBClusterApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
