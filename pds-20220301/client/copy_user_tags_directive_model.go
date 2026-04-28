// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyUserTagsDirective interface {
	dara.Model
	String() string
	GoString() string
	SetDirective(v string) *CopyUserTagsDirective
	GetDirective() *string
	SetKeys(v []*string) *CopyUserTagsDirective
	GetKeys() []*string
}

type CopyUserTagsDirective struct {
	// example:
	//
	// all, include, none, exclude
	Directive *string   `json:"directive,omitempty" xml:"directive,omitempty"`
	Keys      []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
}

func (s CopyUserTagsDirective) String() string {
	return dara.Prettify(s)
}

func (s CopyUserTagsDirective) GoString() string {
	return s.String()
}

func (s *CopyUserTagsDirective) GetDirective() *string {
	return s.Directive
}

func (s *CopyUserTagsDirective) GetKeys() []*string {
	return s.Keys
}

func (s *CopyUserTagsDirective) SetDirective(v string) *CopyUserTagsDirective {
	s.Directive = &v
	return s
}

func (s *CopyUserTagsDirective) SetKeys(v []*string) *CopyUserTagsDirective {
	s.Keys = v
	return s
}

func (s *CopyUserTagsDirective) Validate() error {
	return dara.Validate(s)
}
