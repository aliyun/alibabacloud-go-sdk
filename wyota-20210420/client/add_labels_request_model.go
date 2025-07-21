// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContents(v string) *AddLabelsRequest
	GetLabelContents() *string
}

type AddLabelsRequest struct {
	LabelContents *string `json:"LabelContents,omitempty" xml:"LabelContents,omitempty"`
}

func (s AddLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddLabelsRequest) GetLabelContents() *string {
	return s.LabelContents
}

func (s *AddLabelsRequest) SetLabelContents(v string) *AddLabelsRequest {
	s.LabelContents = &v
	return s
}

func (s *AddLabelsRequest) Validate() error {
	return dara.Validate(s)
}
