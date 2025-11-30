// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaObject interface {
	dara.Model
	String() string
	GoString() string
	SetMedia(v string) *MediaObject
	GetMedia() *string
	SetType(v string) *MediaObject
	GetType() *string
}

type MediaObject struct {
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaObject) String() string {
	return dara.Prettify(s)
}

func (s MediaObject) GoString() string {
	return s.String()
}

func (s *MediaObject) GetMedia() *string {
	return s.Media
}

func (s *MediaObject) GetType() *string {
	return s.Type
}

func (s *MediaObject) SetMedia(v string) *MediaObject {
	s.Media = &v
	return s
}

func (s *MediaObject) SetType(v string) *MediaObject {
	s.Type = &v
	return s
}

func (s *MediaObject) Validate() error {
	return dara.Validate(s)
}
