// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReferenceTag interface {
	dara.Model
	String() string
	GoString() string
	SetReferenceContent(v string) *ReferenceTag
	GetReferenceContent() *string
	SetReferenceTitle(v string) *ReferenceTag
	GetReferenceTitle() *string
}

type ReferenceTag struct {
	ReferenceContent *string `json:"referenceContent,omitempty" xml:"referenceContent,omitempty"`
	ReferenceTitle   *string `json:"referenceTitle,omitempty" xml:"referenceTitle,omitempty"`
}

func (s ReferenceTag) String() string {
	return dara.Prettify(s)
}

func (s ReferenceTag) GoString() string {
	return s.String()
}

func (s *ReferenceTag) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *ReferenceTag) GetReferenceTitle() *string {
	return s.ReferenceTitle
}

func (s *ReferenceTag) SetReferenceContent(v string) *ReferenceTag {
	s.ReferenceContent = &v
	return s
}

func (s *ReferenceTag) SetReferenceTitle(v string) *ReferenceTag {
	s.ReferenceTitle = &v
	return s
}

func (s *ReferenceTag) Validate() error {
	return dara.Validate(s)
}
