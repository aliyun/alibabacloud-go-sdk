// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryService(v *ListAvailableImagesRequestDirectoryService) *ListAvailableImagesRequest
	GetDirectoryService() *ListAvailableImagesRequestDirectoryService
	SetEnableHT(v bool) *ListAvailableImagesRequest
	GetEnableHT() *bool
	SetHPCInterConnect(v string) *ListAvailableImagesRequest
	GetHPCInterConnect() *string
	SetImageOwnerAlias(v string) *ListAvailableImagesRequest
	GetImageOwnerAlias() *string
	SetInstanceType(v string) *ListAvailableImagesRequest
	GetInstanceType() *string
	SetIsPublic(v bool) *ListAvailableImagesRequest
	GetIsPublic() *bool
	SetPageNumber(v int32) *ListAvailableImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAvailableImagesRequest
	GetPageSize() *int32
	SetScheduler(v *ListAvailableImagesRequestScheduler) *ListAvailableImagesRequest
	GetScheduler() *ListAvailableImagesRequestScheduler
}

type ListAvailableImagesRequest struct {
	// The information about the domain account service.
	DirectoryService *ListAvailableImagesRequestDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
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
	Scheduler *ListAvailableImagesRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s ListAvailableImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequest) GetDirectoryService() *ListAvailableImagesRequestDirectoryService {
	return s.DirectoryService
}

func (s *ListAvailableImagesRequest) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *ListAvailableImagesRequest) GetHPCInterConnect() *string {
	return s.HPCInterConnect
}

func (s *ListAvailableImagesRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *ListAvailableImagesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAvailableImagesRequest) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *ListAvailableImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAvailableImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvailableImagesRequest) GetScheduler() *ListAvailableImagesRequestScheduler {
	return s.Scheduler
}

func (s *ListAvailableImagesRequest) SetDirectoryService(v *ListAvailableImagesRequestDirectoryService) *ListAvailableImagesRequest {
	s.DirectoryService = v
	return s
}

func (s *ListAvailableImagesRequest) SetEnableHT(v bool) *ListAvailableImagesRequest {
	s.EnableHT = &v
	return s
}

func (s *ListAvailableImagesRequest) SetHPCInterConnect(v string) *ListAvailableImagesRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *ListAvailableImagesRequest) SetImageOwnerAlias(v string) *ListAvailableImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesRequest) SetInstanceType(v string) *ListAvailableImagesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAvailableImagesRequest) SetIsPublic(v bool) *ListAvailableImagesRequest {
	s.IsPublic = &v
	return s
}

func (s *ListAvailableImagesRequest) SetPageNumber(v int32) *ListAvailableImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesRequest) SetPageSize(v int32) *ListAvailableImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesRequest) SetScheduler(v *ListAvailableImagesRequestScheduler) *ListAvailableImagesRequest {
	s.Scheduler = v
	return s
}

func (s *ListAvailableImagesRequest) Validate() error {
	if s.DirectoryService != nil {
		if err := s.DirectoryService.Validate(); err != nil {
			return err
		}
	}
	if s.Scheduler != nil {
		if err := s.Scheduler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAvailableImagesRequestDirectoryService struct {
	// The type of the domain account.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAvailableImagesRequestDirectoryService) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesRequestDirectoryService) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequestDirectoryService) GetType() *string {
	return s.Type
}

func (s *ListAvailableImagesRequestDirectoryService) GetVersion() *string {
	return s.Version
}

func (s *ListAvailableImagesRequestDirectoryService) SetType(v string) *ListAvailableImagesRequestDirectoryService {
	s.Type = &v
	return s
}

func (s *ListAvailableImagesRequestDirectoryService) SetVersion(v string) *ListAvailableImagesRequestDirectoryService {
	s.Version = &v
	return s
}

func (s *ListAvailableImagesRequestDirectoryService) Validate() error {
	return dara.Validate(s)
}

type ListAvailableImagesRequestScheduler struct {
	// The scheduler type.
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAvailableImagesRequestScheduler) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesRequestScheduler) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequestScheduler) GetType() *string {
	return s.Type
}

func (s *ListAvailableImagesRequestScheduler) GetVersion() *string {
	return s.Version
}

func (s *ListAvailableImagesRequestScheduler) SetType(v string) *ListAvailableImagesRequestScheduler {
	s.Type = &v
	return s
}

func (s *ListAvailableImagesRequestScheduler) SetVersion(v string) *ListAvailableImagesRequestScheduler {
	s.Version = &v
	return s
}

func (s *ListAvailableImagesRequestScheduler) Validate() error {
	return dara.Validate(s)
}
