// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContent(v string) *UpdateLabelRequest
	GetLabelContent() *string
	SetLabelId(v string) *UpdateLabelRequest
	GetLabelId() *string
}

type UpdateLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s UpdateLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateLabelRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *UpdateLabelRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *UpdateLabelRequest) SetLabelContent(v string) *UpdateLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *UpdateLabelRequest) SetLabelId(v string) *UpdateLabelRequest {
	s.LabelId = &v
	return s
}

func (s *UpdateLabelRequest) Validate() error {
	return dara.Validate(s)
}
