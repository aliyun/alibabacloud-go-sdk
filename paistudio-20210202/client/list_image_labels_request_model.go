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
}

type ListImageLabelsRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// a=b,c=d
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	// example:
	//
	// ImageType,Framework,Platform
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
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

func (s *ListImageLabelsRequest) Validate() error {
	return dara.Validate(s)
}
