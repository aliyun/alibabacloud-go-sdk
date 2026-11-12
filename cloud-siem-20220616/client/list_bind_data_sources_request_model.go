// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListBindDataSourcesRequest
	GetAccountId() *string
	SetCloudCode(v string) *ListBindDataSourcesRequest
	GetCloudCode() *string
	SetRegionId(v string) *ListBindDataSourcesRequest
	GetRegionId() *string
}

type ListBindDataSourcesRequest struct {
	// The cloud account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the multi-cloud provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region where the threat analysis data management center resides. Select the management center region based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBindDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListBindDataSourcesRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListBindDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBindDataSourcesRequest) SetAccountId(v string) *ListBindDataSourcesRequest {
	s.AccountId = &v
	return s
}

func (s *ListBindDataSourcesRequest) SetCloudCode(v string) *ListBindDataSourcesRequest {
	s.CloudCode = &v
	return s
}

func (s *ListBindDataSourcesRequest) SetRegionId(v string) *ListBindDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBindDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
