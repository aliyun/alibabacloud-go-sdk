// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ListImageLabelsRequest
	GetImageId() *string
	SetLabelFilter(v string) *ListImageLabelsRequest
	GetLabelFilter() *string
	SetLabelKeys(v string) *ListImageLabelsRequest
	GetLabelKeys() *string
	SetRegion(v string) *ListImageLabelsRequest
	GetRegion() *string
	SetWorkspaceId(v string) *ListImageLabelsRequest
	GetWorkspaceId() *string
}

type ListImageLabelsRequest struct {
	// The image ID. For more information about how to obtain an image ID, see [ListImages](https://help.aliyun.com/document_detail/449118.html).
	//
	// example:
	//
	// image-4c62******53uor
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The filter conditions for labels. Separate multiple conditions with commas.
	//
	// The format for a single condition is `key=value`.
	//
	// This parameter works independently of the LabelKeys parameter.
	//
	// example:
	//
	// system.framework=XGBoost 1.6.0,system.official=true
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	// The list of tag keys. Separate multiple keys with commas (,).
	//
	// System tags start with "system". This parameter works independently of the LabelFilter parameter.
	//
	// example:
	//
	// system.framework,system.official
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// The region where the image resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The workspace ID. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImageLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListImageLabelsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageLabelsRequest) GetLabelFilter() *string {
	return s.LabelFilter
}

func (s *ListImageLabelsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *ListImageLabelsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListImageLabelsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListImageLabelsRequest) SetImageId(v string) *ListImageLabelsRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelFilter(v string) *ListImageLabelsRequest {
	s.LabelFilter = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelKeys(v string) *ListImageLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *ListImageLabelsRequest) SetRegion(v string) *ListImageLabelsRequest {
	s.Region = &v
	return s
}

func (s *ListImageLabelsRequest) SetWorkspaceId(v string) *ListImageLabelsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListImageLabelsRequest) Validate() error {
	return dara.Validate(s)
}
