// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsight interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v string) *VideoInsight
	GetCaption() *string
	SetDescription(v string) *VideoInsight
	GetDescription() *string
}

type VideoInsight struct {
	// if can be null:
	// true
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s VideoInsight) String() string {
	return dara.Prettify(s)
}

func (s VideoInsight) GoString() string {
	return s.String()
}

func (s *VideoInsight) GetCaption() *string {
	return s.Caption
}

func (s *VideoInsight) GetDescription() *string {
	return s.Description
}

func (s *VideoInsight) SetCaption(v string) *VideoInsight {
	s.Caption = &v
	return s
}

func (s *VideoInsight) SetDescription(v string) *VideoInsight {
	s.Description = &v
	return s
}

func (s *VideoInsight) Validate() error {
	return dara.Validate(s)
}
