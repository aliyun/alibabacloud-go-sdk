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
	SetSize(v int64) *AddImageRequest
	GetSize() *int64
	SetSourceId(v string) *AddImageRequest
	GetSourceId() *string
	SetSourceType(v string) *AddImageRequest
	GetSourceType() *string
	SetUserId(v string) *AddImageRequest
	GetUserId() *string
	SetWorkspaceId(v string) *AddImageRequest
	GetWorkspaceId() *string
}

type AddImageRequest struct {
	// The visibility of the image. Valid values:
	//
	// - PUBLIC: All members of the workspace can perform operations on the image.
	//
	// - PRIVATE: Only the creator can perform operations on the image.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// NLP model compression training image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image. If you leave this parameter empty, the system automatically generates an ID.
	//
	// The format is \\`image-\\` followed by 18 uppercase letters, lowercase letters, or digits.
	//
	// example:
	//
	// image-k83*****cv
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The URI of the image. The URI can be reused. For more information, see [ListImage](https://help.aliyun.com/document_detail/449118.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-compression/nlp:gpu
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The labels of the image. This is an array where each item contains a key and a value.
	//
	// Official images have the following label: system.official=true
	//
	// The following keys are supported:
	//
	// - system.chipType
	//
	// - system.dsw\\.cudaVersion
	//
	// - system.dsw\\.fromImageId
	//
	// - system.dsw\\.fromInstanceId
	//
	// - system.dsw\\.id
	//
	// - system.dsw\\.os
	//
	// - system.dsw\\.osVersion
	//
	// - system.dsw\\.resourceType
	//
	// - system.dsw\\.rootImageId
	//
	// - system.dsw\\.stage
	//
	// - system.dsw\\.tag
	//
	// - system.dsw\\.type
	//
	// - system.framework
	//
	// - system.origin
	//
	// - system.pythonVersion
	//
	// - system.source
	//
	// - system.supported.dlc
	//
	// - system.supported.dsw
	Labels []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The image name. The naming convention is as follows:
	//
	// - The name must be 1 to 50 characters long.
	//
	// - The name can contain lowercase letters, digits, and hyphens (-). It must start with a letter.
	//
	// - The name must be unique within the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// nlp-compression
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the image, in GB.
	//
	// example:
	//
	// 2
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The source ID of the image. If the source type is Build, this ID corresponds to the image build ID.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the image. Valid values:
	//
	// Import
	//
	// Build
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace to which the image belongs. For more information, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *AddImageRequest) GetSize() *int64 {
	return s.Size
}

func (s *AddImageRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *AddImageRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *AddImageRequest) GetUserId() *string {
	return s.UserId
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

func (s *AddImageRequest) SetSize(v int64) *AddImageRequest {
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

func (s *AddImageRequest) SetUserId(v string) *AddImageRequest {
	s.UserId = &v
	return s
}

func (s *AddImageRequest) SetWorkspaceId(v string) *AddImageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddImageRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddImageRequestLabels struct {
	// The key of the label.
	//
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label.
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
