// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v string) *DeleteModelLabelsRequest
	GetLabelKeys() *string
}

type DeleteModelLabelsRequest struct {
	// The label key to be deleted. To delete multiple label keys, separate them with commas (,).
	//
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteModelLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *DeleteModelLabelsRequest) SetLabelKeys(v string) *DeleteModelLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *DeleteModelLabelsRequest) Validate() error {
	return dara.Validate(s)
}
