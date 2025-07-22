// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v string) *DeleteInstanceLabelsRequest
	GetLabelKeys() *string
}

type DeleteInstanceLabelsRequest struct {
	// The keys of the tags that you want to delete. Separate multiple tags with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// labelKey1,labelKey2,labelKey3
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteInstanceLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceLabelsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *DeleteInstanceLabelsRequest) SetLabelKeys(v string) *DeleteInstanceLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *DeleteInstanceLabelsRequest) Validate() error {
	return dara.Validate(s)
}
