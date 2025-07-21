// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v string) *DeleteLabelRequest
	GetForce() *string
	SetLabelContent(v string) *DeleteLabelRequest
	GetLabelContent() *string
	SetLabelId(v string) *DeleteLabelRequest
	GetLabelId() *string
}

type DeleteLabelRequest struct {
	// This parameter is required.
	Force        *string `json:"Force,omitempty" xml:"Force,omitempty"`
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s DeleteLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteLabelRequest) GetForce() *string {
	return s.Force
}

func (s *DeleteLabelRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *DeleteLabelRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *DeleteLabelRequest) SetForce(v string) *DeleteLabelRequest {
	s.Force = &v
	return s
}

func (s *DeleteLabelRequest) SetLabelContent(v string) *DeleteLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *DeleteLabelRequest) SetLabelId(v string) *DeleteLabelRequest {
	s.LabelId = &v
	return s
}

func (s *DeleteLabelRequest) Validate() error {
	return dara.Validate(s)
}
