// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTypeDetermineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHistoryShrink(v string) *ModelTypeDetermineShrinkRequest
	GetHistoryShrink() *string
	SetInputText(v string) *ModelTypeDetermineShrinkRequest
	GetInputText() *string
}

type ModelTypeDetermineShrinkRequest struct {
	HistoryShrink *string `json:"history,omitempty" xml:"history,omitempty"`
	// This parameter is required.
	InputText *string `json:"inputText,omitempty" xml:"inputText,omitempty"`
}

func (s ModelTypeDetermineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineShrinkRequest) GetHistoryShrink() *string {
	return s.HistoryShrink
}

func (s *ModelTypeDetermineShrinkRequest) GetInputText() *string {
	return s.InputText
}

func (s *ModelTypeDetermineShrinkRequest) SetHistoryShrink(v string) *ModelTypeDetermineShrinkRequest {
	s.HistoryShrink = &v
	return s
}

func (s *ModelTypeDetermineShrinkRequest) SetInputText(v string) *ModelTypeDetermineShrinkRequest {
	s.InputText = &v
	return s
}

func (s *ModelTypeDetermineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
