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
	// The code of the third-party cloud service.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
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
