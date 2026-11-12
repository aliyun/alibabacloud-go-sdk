// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *ListImageBaseInfoRequest
	GetBizType() *int32
	SetImageId(v string) *ListImageBaseInfoRequest
	GetImageId() *string
	SetImageName(v string) *ListImageBaseInfoRequest
	GetImageName() *string
	SetImageType(v string) *ListImageBaseInfoRequest
	GetImageType() *string
	SetOsType(v string) *ListImageBaseInfoRequest
	GetOsType() *string
	SetPackageType(v string) *ListImageBaseInfoRequest
	GetPackageType() *string
	SetPageNumber(v int32) *ListImageBaseInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageBaseInfoRequest
	GetPageSize() *int32
	SetProductType(v string) *ListImageBaseInfoRequest
	GetProductType() *string
	SetProtocolType(v string) *ListImageBaseInfoRequest
	GetProtocolType() *string
	SetStatus(v string) *ListImageBaseInfoRequest
	GetStatus() *string
}

type ListImageBaseInfoRequest struct {
	// The business line code. Default value: 1.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The image ID. You can specify multiple IDs separated by commas (,). The query matches any of the specified IDs.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image query keyword. Fuzzy matching is performed on both the image name and the image ID.
	//
	// example:
	//
	// test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image type. Valid values: System (system image), User (custom image), Shared (shared image), and Community (community image).
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image operating system type. Valid values: Windows (Windows operating system), Linux (Linux operating system), Android (Android operating system), and Unknown (unknown operating system).
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The image package type. Valid values: Ecs (ECS image), Vhd (VHD image), Container (container image), Ecs_Container (ECS and container composite image), and K8s_Node (Kubernetes node image).
	//
	// example:
	//
	// Ecs
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The page number. Pages start from page 1. If you do not specify this parameter or specify a value less than 1, the first page is returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100. If the value is less than 1, 10 entries are returned. If the value is greater than 100, 100 entries are returned.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type, which is used to filter the products to which the images belong. Examples: CloudDesktop (WUYING Workspace), CloudApp (WUYING Cloud App), and WuyingServer (WUYING Workstation).
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The image protocol type. Valid values: ASP (Alibaba Cloud self-developed ASP protocol) and HDX (HDX protocol).
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The image status. Valid values: INIT (initialization), CREATING (creating), AVAILABLE (available), and CREATE_FAILED (creation failed). If you do not specify this parameter, all non-deleted images are queried. Deleted images are not returned.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListImageBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *ListImageBaseInfoRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListImageBaseInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageBaseInfoRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListImageBaseInfoRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListImageBaseInfoRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListImageBaseInfoRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *ListImageBaseInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageBaseInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageBaseInfoRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListImageBaseInfoRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListImageBaseInfoRequest) GetStatus() *string {
	return s.Status
}

func (s *ListImageBaseInfoRequest) SetBizType(v int32) *ListImageBaseInfoRequest {
	s.BizType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetImageId(v string) *ListImageBaseInfoRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetImageName(v string) *ListImageBaseInfoRequest {
	s.ImageName = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetImageType(v string) *ListImageBaseInfoRequest {
	s.ImageType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetOsType(v string) *ListImageBaseInfoRequest {
	s.OsType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetPackageType(v string) *ListImageBaseInfoRequest {
	s.PackageType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetPageNumber(v int32) *ListImageBaseInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetPageSize(v int32) *ListImageBaseInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetProductType(v string) *ListImageBaseInfoRequest {
	s.ProductType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetProtocolType(v string) *ListImageBaseInfoRequest {
	s.ProtocolType = &v
	return s
}

func (s *ListImageBaseInfoRequest) SetStatus(v string) *ListImageBaseInfoRequest {
	s.Status = &v
	return s
}

func (s *ListImageBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
