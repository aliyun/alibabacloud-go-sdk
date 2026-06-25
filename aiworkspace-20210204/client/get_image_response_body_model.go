// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetImageResponseBody
	GetAccessibility() *string
	SetDescription(v string) *GetImageResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetImageResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetImageResponseBody
	GetGmtModifiedTime() *string
	SetImageUri(v string) *GetImageResponseBody
	GetImageUri() *string
	SetLabels(v []*GetImageResponseBodyLabels) *GetImageResponseBody
	GetLabels() []*GetImageResponseBodyLabels
	SetName(v string) *GetImageResponseBody
	GetName() *string
	SetParentUserId(v string) *GetImageResponseBody
	GetParentUserId() *string
	SetRequestId(v string) *GetImageResponseBody
	GetRequestId() *string
	SetSize(v int64) *GetImageResponseBody
	GetSize() *int64
	SetSourceId(v string) *GetImageResponseBody
	GetSourceId() *string
	SetSourceType(v string) *GetImageResponseBody
	GetSourceType() *string
	SetUserId(v string) *GetImageResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetImageResponseBody
	GetWorkspaceId() *string
}

type GetImageResponseBody struct {
	// The visibility of the image. Valid values:
	//
	// - PUBLIC: All members in the current workspace can perform operations on the image.
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
	// The time when the image was created. The time is in UTC and the format is ISO 8601.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the image was last modified. The time is in UTC and the format is ISO 8601.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The URL of the image, including the version number.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.******ession/nlp:gpu
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// A list of image labels. This is an array. Each item in the array contains a Key and a Value field.
	//
	// Official images have the following label: the key is system.official and the value is true.
	Labels []*GetImageResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the image.
	//
	// example:
	//
	// nlp-compression
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account of the creator.
	//
	// example:
	//
	// 15577******8921
	ParentUserId *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The size of the image in bytes.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the image source.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the image source.
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The UID of the user who created the image.
	//
	// example:
	//
	// 15577******8921
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace to which the image belongs.
	//
	// example:
	//
	// 15945
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetImageResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetImageResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetImageResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetImageResponseBody) GetImageUri() *string {
	return s.ImageUri
}

func (s *GetImageResponseBody) GetLabels() []*GetImageResponseBodyLabels {
	return s.Labels
}

func (s *GetImageResponseBody) GetName() *string {
	return s.Name
}

func (s *GetImageResponseBody) GetParentUserId() *string {
	return s.ParentUserId
}

func (s *GetImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *GetImageResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetImageResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetImageResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetImageResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetImageResponseBody) SetAccessibility(v string) *GetImageResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetImageResponseBody) SetDescription(v string) *GetImageResponseBody {
	s.Description = &v
	return s
}

func (s *GetImageResponseBody) SetGmtCreateTime(v string) *GetImageResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetImageResponseBody) SetGmtModifiedTime(v string) *GetImageResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageUri(v string) *GetImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBody) SetLabels(v []*GetImageResponseBodyLabels) *GetImageResponseBody {
	s.Labels = v
	return s
}

func (s *GetImageResponseBody) SetName(v string) *GetImageResponseBody {
	s.Name = &v
	return s
}

func (s *GetImageResponseBody) SetParentUserId(v string) *GetImageResponseBody {
	s.ParentUserId = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSize(v int64) *GetImageResponseBody {
	s.Size = &v
	return s
}

func (s *GetImageResponseBody) SetSourceId(v string) *GetImageResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetImageResponseBody) SetSourceType(v string) *GetImageResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetImageResponseBody) SetUserId(v string) *GetImageResponseBody {
	s.UserId = &v
	return s
}

func (s *GetImageResponseBody) SetWorkspaceId(v string) *GetImageResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetImageResponseBody) Validate() error {
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

type GetImageResponseBodyLabels struct {
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

func (s GetImageResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *GetImageResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *GetImageResponseBodyLabels) SetKey(v string) *GetImageResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetImageResponseBodyLabels) SetValue(v string) *GetImageResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *GetImageResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
