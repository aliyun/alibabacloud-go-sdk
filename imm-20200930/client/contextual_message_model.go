// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualMessage interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ContextualMessage
	GetContent() *string
	SetFiles(v []*ContextualFile) *ContextualMessage
	GetFiles() []*ContextualFile
	SetRole(v string) *ContextualMessage
	GetRole() *string
}

type ContextualMessage struct {
	// example:
	//
	// 你好
	Content *string           `json:"Content,omitempty" xml:"Content,omitempty"`
	Files   []*ContextualFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ContextualMessage) String() string {
	return dara.Prettify(s)
}

func (s ContextualMessage) GoString() string {
	return s.String()
}

func (s *ContextualMessage) GetContent() *string {
	return s.Content
}

func (s *ContextualMessage) GetFiles() []*ContextualFile {
	return s.Files
}

func (s *ContextualMessage) GetRole() *string {
	return s.Role
}

func (s *ContextualMessage) SetContent(v string) *ContextualMessage {
	s.Content = &v
	return s
}

func (s *ContextualMessage) SetFiles(v []*ContextualFile) *ContextualMessage {
	s.Files = v
	return s
}

func (s *ContextualMessage) SetRole(v string) *ContextualMessage {
	s.Role = &v
	return s
}

func (s *ContextualMessage) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
