// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v string) *DeleteModelVersionLabelsRequest
	GetLabelKeys() *string
}

type DeleteModelVersionLabelsRequest struct {
	// The key of the tag to be deleted. Separate multiple tag keys with commas (,).
	//
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteModelVersionLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *DeleteModelVersionLabelsRequest) SetLabelKeys(v string) *DeleteModelVersionLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *DeleteModelVersionLabelsRequest) Validate() error {
	return dara.Validate(s)
}
