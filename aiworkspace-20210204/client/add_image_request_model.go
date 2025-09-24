// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *AddImageRequest
	GetAccessibility() *string
	SetDescription(v string) *AddImageRequest
	GetDescription() *string
	SetImageId(v string) *AddImageRequest
	GetImageId() *string
	SetImageUri(v string) *AddImageRequest
	GetImageUri() *string
	SetLabels(v []*AddImageRequestLabels) *AddImageRequest
	GetLabels() []*AddImageRequestLabels
	SetName(v string) *AddImageRequest
	GetName() *string
	SetSize(v int32) *AddImageRequest
	GetSize() *int32
	SetSourceId(v string) *AddImageRequest
	GetSourceId() *string
	SetSourceType(v string) *AddImageRequest
	GetSourceType() *string
	SetWorkspaceId(v string) *AddImageRequest
	GetWorkspaceId() *string
}

type AddImageRequest struct {
	// The accessibility of the image. Valid values:
	//
	// 	- PUBLIC: The image is accessible to all members in the workspace.
	//
	// 	- PRIVATE: The image is accessible only to the image creator.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The image description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID. If you do not specify this parameter, the system automatically generates an image ID. The image ID must start with image- followed by 18 characters in letters or digits.
	//
	// example:
	//
	// image-k83*****cv
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The URL of the image, which can be repeated. You can call [ListImage](https://help.aliyun.com/document_detail/449118.html) to view the image URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-compression/nlp:gpu
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The image tag, which is an array. Each element in the array contains a key-value pair. Alibaba Cloud images have the system.official=true tag. You can add the following keys to an image:
	//
	// 	- system.chipType
	//
	// 	- system.dsw.cudaVersion
	//
	// 	- system.dsw.fromImageId
	//
	// 	- system.dsw.fromInstanceId
	//
	// 	- system.dsw.id
	//
	// 	- system.dsw.os
	//
	// 	- system.dsw.osVersion
	//
	// 	- system.dsw.resourceType
	//
	// 	- system.dsw.rootImageId
	//
	// 	- system.dsw.stage
	//
	// 	- system.dsw.tag
	//
	// 	- system.dsw.type
	//
	// 	- system.framework
	//
	// 	- system.origin
	//
	// 	- system.pythonVersion
	//
	// 	- system.source
	//
	// 	- system.supported.dlc
	//
	// 	- system.supported.dsw
	Labels []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The image name. The name must meet the following requirements:
	//
	// 	- The name must be 1 to 50 characters in length.
	//
	// 	- The name can contain lowercase letters, digits, and hyphens (-). The name must start with a lowercase letter.
	//
	// 	- The name must be unique in a workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// nlp-compression
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the image. Unit: GB.
	//
	// example:
	//
	// 2
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SourceId   *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 15******45
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddImageRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *AddImageRequest) GetDescription() *string {
	return s.Description
}

func (s *AddImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AddImageRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *AddImageRequest) GetLabels() []*AddImageRequestLabels {
	return s.Labels
}

func (s *AddImageRequest) GetName() *string {
	return s.Name
}

func (s *AddImageRequest) GetSize() *int32 {
	return s.Size
}

func (s *AddImageRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *AddImageRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *AddImageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddImageRequest) SetAccessibility(v string) *AddImageRequest {
	s.Accessibility = &v
	return s
}

func (s *AddImageRequest) SetDescription(v string) *AddImageRequest {
	s.Description = &v
	return s
}

func (s *AddImageRequest) SetImageId(v string) *AddImageRequest {
	s.ImageId = &v
	return s
}

func (s *AddImageRequest) SetImageUri(v string) *AddImageRequest {
	s.ImageUri = &v
	return s
}

func (s *AddImageRequest) SetLabels(v []*AddImageRequestLabels) *AddImageRequest {
	s.Labels = v
	return s
}

func (s *AddImageRequest) SetName(v string) *AddImageRequest {
	s.Name = &v
	return s
}

func (s *AddImageRequest) SetSize(v int32) *AddImageRequest {
	s.Size = &v
	return s
}

func (s *AddImageRequest) SetSourceId(v string) *AddImageRequest {
	s.SourceId = &v
	return s
}

func (s *AddImageRequest) SetSourceType(v string) *AddImageRequest {
	s.SourceType = &v
	return s
}

func (s *AddImageRequest) SetWorkspaceId(v string) *AddImageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddImageRequest) Validate() error {
	return dara.Validate(s)
}

type AddImageRequestLabels struct {
	// The tag key.
	//
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// GPU
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageRequestLabels) GetKey() *string {
	return s.Key
}

func (s *AddImageRequestLabels) GetValue() *string {
	return s.Value
}

func (s *AddImageRequestLabels) SetKey(v string) *AddImageRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageRequestLabels) SetValue(v string) *AddImageRequestLabels {
	s.Value = &v
	return s
}

func (s *AddImageRequestLabels) Validate() error {
	return dara.Validate(s)
}
