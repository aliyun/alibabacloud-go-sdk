// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *ListDataSourceTypesRequest
	GetCloudCode() *string
	SetRegionId(v string) *ListDataSourceTypesRequest
	GetRegionId() *string
}

type ListDataSourceTypesRequest struct {
	// The code of the multi-cloud provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets belong to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: The assets belong to regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDataSourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListDataSourceTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourceTypesRequest) SetCloudCode(v string) *ListDataSourceTypesRequest {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceTypesRequest) SetRegionId(v string) *ListDataSourceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
