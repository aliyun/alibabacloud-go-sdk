// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetUploadInfoRequest
	GetName() *string
	SetRegionId(v string) *GetUploadInfoRequest
	GetRegionId() *string
	SetResourceType(v string) *GetUploadInfoRequest
	GetResourceType() *string
}

type GetUploadInfoRequest struct {
	// Upload name.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUploadInfoRequest) GetName() *string {
	return s.Name
}

func (s *GetUploadInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUploadInfoRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetUploadInfoRequest) SetName(v string) *GetUploadInfoRequest {
	s.Name = &v
	return s
}

func (s *GetUploadInfoRequest) SetRegionId(v string) *GetUploadInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetUploadInfoRequest) SetResourceType(v string) *GetUploadInfoRequest {
	s.ResourceType = &v
	return s
}

func (s *GetUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}
