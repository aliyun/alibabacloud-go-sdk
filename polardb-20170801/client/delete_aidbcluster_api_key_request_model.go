// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DeleteAIDBClusterApiKeyRequest
	GetApiKey() *string
	SetModelSpaceName(v string) *DeleteAIDBClusterApiKeyRequest
	GetModelSpaceName() *string
	SetRegionId(v string) *DeleteAIDBClusterApiKeyRequest
	GetRegionId() *string
}

type DeleteAIDBClusterApiKeyRequest struct {
	// The API key of the model service.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The model space ID.
	//
	// example:
	//
	// pms-xxx
	ModelSpaceName *string `json:"ModelSpaceName,omitempty" xml:"ModelSpaceName,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAIDBClusterApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterApiKeyRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeleteAIDBClusterApiKeyRequest) GetModelSpaceName() *string {
	return s.ModelSpaceName
}

func (s *DeleteAIDBClusterApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAIDBClusterApiKeyRequest) SetApiKey(v string) *DeleteAIDBClusterApiKeyRequest {
	s.ApiKey = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyRequest) SetModelSpaceName(v string) *DeleteAIDBClusterApiKeyRequest {
	s.ModelSpaceName = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyRequest) SetRegionId(v string) *DeleteAIDBClusterApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
