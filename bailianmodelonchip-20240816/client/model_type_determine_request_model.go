// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTypeDetermineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHistory(v []*ModelTypeDetermineRequestHistory) *ModelTypeDetermineRequest
	GetHistory() []*ModelTypeDetermineRequestHistory
	SetInputText(v string) *ModelTypeDetermineRequest
	GetInputText() *string
}

type ModelTypeDetermineRequest struct {
	History []*ModelTypeDetermineRequestHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	// This parameter is required.
	InputText *string `json:"inputText,omitempty" xml:"inputText,omitempty"`
}

func (s ModelTypeDetermineRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineRequest) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineRequest) GetHistory() []*ModelTypeDetermineRequestHistory {
	return s.History
}

func (s *ModelTypeDetermineRequest) GetInputText() *string {
	return s.InputText
}

func (s *ModelTypeDetermineRequest) SetHistory(v []*ModelTypeDetermineRequestHistory) *ModelTypeDetermineRequest {
	s.History = v
	return s
}

func (s *ModelTypeDetermineRequest) SetInputText(v string) *ModelTypeDetermineRequest {
	s.InputText = &v
	return s
}

func (s *ModelTypeDetermineRequest) Validate() error {
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelTypeDetermineRequestHistory struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ModelTypeDetermineRequestHistory) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineRequestHistory) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineRequestHistory) GetContent() *string {
	return s.Content
}

func (s *ModelTypeDetermineRequestHistory) GetRole() *string {
	return s.Role
}

func (s *ModelTypeDetermineRequestHistory) SetContent(v string) *ModelTypeDetermineRequestHistory {
	s.Content = &v
	return s
}

func (s *ModelTypeDetermineRequestHistory) SetRole(v string) *ModelTypeDetermineRequestHistory {
	s.Role = &v
	return s
}

func (s *ModelTypeDetermineRequestHistory) Validate() error {
	return dara.Validate(s)
}
