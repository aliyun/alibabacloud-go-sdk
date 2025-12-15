// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableImagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryServiceShrink(v string) *ListAvailableImagesShrinkRequest
	GetDirectoryServiceShrink() *string
	SetEnableHT(v bool) *ListAvailableImagesShrinkRequest
	GetEnableHT() *bool
	SetHPCInterConnect(v string) *ListAvailableImagesShrinkRequest
	GetHPCInterConnect() *string
	SetImageOwnerAlias(v string) *ListAvailableImagesShrinkRequest
	GetImageOwnerAlias() *string
	SetInstanceType(v string) *ListAvailableImagesShrinkRequest
	GetInstanceType() *string
	SetIsPublic(v bool) *ListAvailableImagesShrinkRequest
	GetIsPublic() *bool
	SetPageNumber(v int32) *ListAvailableImagesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAvailableImagesShrinkRequest
	GetPageSize() *int32
	SetSchedulerShrink(v string) *ListAvailableImagesShrinkRequest
	GetSchedulerShrink() *string
}

type ListAvailableImagesShrinkRequest struct {
	// The information about the domain account service.
	DirectoryServiceShrink *string `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty"`
	// Specifies whether to return images in which hyper-threading is enabled.
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The network type of the images that you want to query.
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: system image.
	//
	// 	- self: custom image.
	//
	// 	- others: shared image.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The instance type.
	//
	// >  By default, if you do not specify an instance type, the list of images that are supported by all instance types are queried. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to return published community images. Valid values:
	//
	// 	- true: returns published community images. If you set the value of this parameter to `true`, the `ImageOwnerAlias` parameter must be set to `others`.
	//
	// 	- false: returns non-community images. The value of the `ImageOwnerAlias` parameter prevails.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scheduler information about the images that you want to query.
	SchedulerShrink *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s ListAvailableImagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesShrinkRequest) GetDirectoryServiceShrink() *string {
	return s.DirectoryServiceShrink
}

func (s *ListAvailableImagesShrinkRequest) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *ListAvailableImagesShrinkRequest) GetHPCInterConnect() *string {
	return s.HPCInterConnect
}

func (s *ListAvailableImagesShrinkRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *ListAvailableImagesShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAvailableImagesShrinkRequest) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *ListAvailableImagesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAvailableImagesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvailableImagesShrinkRequest) GetSchedulerShrink() *string {
	return s.SchedulerShrink
}

func (s *ListAvailableImagesShrinkRequest) SetDirectoryServiceShrink(v string) *ListAvailableImagesShrinkRequest {
	s.DirectoryServiceShrink = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetEnableHT(v bool) *ListAvailableImagesShrinkRequest {
	s.EnableHT = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetHPCInterConnect(v string) *ListAvailableImagesShrinkRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetImageOwnerAlias(v string) *ListAvailableImagesShrinkRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetInstanceType(v string) *ListAvailableImagesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetIsPublic(v bool) *ListAvailableImagesShrinkRequest {
	s.IsPublic = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetPageNumber(v int32) *ListAvailableImagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetPageSize(v int32) *ListAvailableImagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetSchedulerShrink(v string) *ListAvailableImagesShrinkRequest {
	s.SchedulerShrink = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
