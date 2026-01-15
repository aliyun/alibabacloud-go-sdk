// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHottopicNews interface {
	dara.Model
	String() string
	GoString() string
	SetComments(v []*HottopicNewsComments) *HottopicNews
	GetComments() []*HottopicNewsComments
	SetContent(v string) *HottopicNews
	GetContent() *string
	SetTitle(v string) *HottopicNews
	GetTitle() *string
	SetUrl(v string) *HottopicNews
	GetUrl() *string
}

type HottopicNews struct {
	Comments []*HottopicNewsComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	Content  *string                 `json:"Content,omitempty" xml:"Content,omitempty"`
	Title    *string                 `json:"Title,omitempty" xml:"Title,omitempty"`
	Url      *string                 `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s HottopicNews) String() string {
	return dara.Prettify(s)
}

func (s HottopicNews) GoString() string {
	return s.String()
}

func (s *HottopicNews) GetComments() []*HottopicNewsComments {
	return s.Comments
}

func (s *HottopicNews) GetContent() *string {
	return s.Content
}

func (s *HottopicNews) GetTitle() *string {
	return s.Title
}

func (s *HottopicNews) GetUrl() *string {
	return s.Url
}

func (s *HottopicNews) SetComments(v []*HottopicNewsComments) *HottopicNews {
	s.Comments = v
	return s
}

func (s *HottopicNews) SetContent(v string) *HottopicNews {
	s.Content = &v
	return s
}

func (s *HottopicNews) SetTitle(v string) *HottopicNews {
	s.Title = &v
	return s
}

func (s *HottopicNews) SetUrl(v string) *HottopicNews {
	s.Url = &v
	return s
}

func (s *HottopicNews) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HottopicNewsComments struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s HottopicNewsComments) String() string {
	return dara.Prettify(s)
}

func (s HottopicNewsComments) GoString() string {
	return s.String()
}

func (s *HottopicNewsComments) GetText() *string {
	return s.Text
}

func (s *HottopicNewsComments) SetText(v string) *HottopicNewsComments {
	s.Text = &v
	return s
}

func (s *HottopicNewsComments) Validate() error {
	return dara.Validate(s)
}
