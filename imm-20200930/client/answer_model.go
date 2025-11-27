// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswer interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *Answer
	GetContent() *string
	SetReferences(v []*ReferenceFile) *Answer
	GetReferences() []*ReferenceFile
}

type Answer struct {
	// example:
	//
	// 你好
	Content    *string          `json:"Content,omitempty" xml:"Content,omitempty"`
	References []*ReferenceFile `json:"References,omitempty" xml:"References,omitempty" type:"Repeated"`
}

func (s Answer) String() string {
	return dara.Prettify(s)
}

func (s Answer) GoString() string {
	return s.String()
}

func (s *Answer) GetContent() *string {
	return s.Content
}

func (s *Answer) GetReferences() []*ReferenceFile {
	return s.References
}

func (s *Answer) SetContent(v string) *Answer {
	s.Content = &v
	return s
}

func (s *Answer) SetReferences(v []*ReferenceFile) *Answer {
	s.References = v
	return s
}

func (s *Answer) Validate() error {
	if s.References != nil {
		for _, item := range s.References {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
