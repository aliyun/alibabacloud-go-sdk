// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlert interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *Alert
	GetBody() *string
	SetSubtitle(v string) *Alert
	GetSubtitle() *string
	SetTitle(v string) *Alert
	GetTitle() *string
}

type Alert struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Subtitle *string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Alert) String() string {
	return dara.Prettify(s)
}

func (s Alert) GoString() string {
	return s.String()
}

func (s *Alert) GetBody() *string {
	return s.Body
}

func (s *Alert) GetSubtitle() *string {
	return s.Subtitle
}

func (s *Alert) GetTitle() *string {
	return s.Title
}

func (s *Alert) SetBody(v string) *Alert {
	s.Body = &v
	return s
}

func (s *Alert) SetSubtitle(v string) *Alert {
	s.Subtitle = &v
	return s
}

func (s *Alert) SetTitle(v string) *Alert {
	s.Title = &v
	return s
}

func (s *Alert) Validate() error {
	return dara.Validate(s)
}
