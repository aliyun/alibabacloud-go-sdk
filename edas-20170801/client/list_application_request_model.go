// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *ListApplicationRequest
	GetAppIds() *string
	SetAppName(v string) *ListApplicationRequest
	GetAppName() *string
	SetClusterId(v string) *ListApplicationRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *ListApplicationRequest
	GetCurrentPage() *int32
	SetLogicalRegionId(v string) *ListApplicationRequest
	GetLogicalRegionId() *string
	SetLogicalRegionIdFilter(v string) *ListApplicationRequest
	GetLogicalRegionIdFilter() *string
	SetPageSize(v int32) *ListApplicationRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListApplicationRequest
	GetResourceGroupId() *string
}

type ListApplicationRequest struct {
	// The application IDs.
	//
	// example:
	//
	// [
	//
	//       "5657d271-****-4f03-9bb2-431f942886bb",
	//
	//       "5657d271-****-4f03-9bb2-431f942bbddd"
	//
	// ]
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The name of the application. Specify this parameter if you want to filter applications by application name.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID. Specify this parameter if you want to filter applications by cluster.
	//
	// example:
	//
	// c37aec2a-bcca-4ec1-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The namespace ID. Specify this parameter if you want to filter applications by namespace.
	//
	// example:
	//
	// cn-beijing:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The ID of the namespace that you use in the exact search to filter applications.
	//
	// example:
	//
	// cn-beijing:test
	LogicalRegionIdFilter *string `json:"LogicalRegionIdFilter,omitempty" xml:"LogicalRegionIdFilter,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. Specify this parameter if you want to filter applications by resource group.
	//
	// example:
	//
	// rg-aek24j4s4b*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *ListApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListApplicationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListApplicationRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListApplicationRequest) GetLogicalRegionIdFilter() *string {
	return s.LogicalRegionIdFilter
}

func (s *ListApplicationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationRequest) SetAppIds(v string) *ListApplicationRequest {
	s.AppIds = &v
	return s
}

func (s *ListApplicationRequest) SetAppName(v string) *ListApplicationRequest {
	s.AppName = &v
	return s
}

func (s *ListApplicationRequest) SetClusterId(v string) *ListApplicationRequest {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationRequest) SetCurrentPage(v int32) *ListApplicationRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationRequest) SetLogicalRegionId(v string) *ListApplicationRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListApplicationRequest) SetLogicalRegionIdFilter(v string) *ListApplicationRequest {
	s.LogicalRegionIdFilter = &v
	return s
}

func (s *ListApplicationRequest) SetPageSize(v int32) *ListApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationRequest) SetResourceGroupId(v string) *ListApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationRequest) Validate() error {
	return dara.Validate(s)
}
