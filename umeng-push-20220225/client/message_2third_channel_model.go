// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMessage2ThirdChannel interface {
	dara.Model
	String() string
	GoString() string
	SetSetBadge(v int64) *Message2ThirdChannel
	GetSetBadge() *int64
	SetAddBadge(v int64) *Message2ThirdChannel
	GetAddBadge() *int64
	SetBigBody(v string) *Message2ThirdChannel
	GetBigBody() *string
	SetBigTitle(v string) *Message2ThirdChannel
	GetBigTitle() *string
	SetExpandImage(v string) *Message2ThirdChannel
	GetExpandImage() *string
	SetImg(v string) *Message2ThirdChannel
	GetImg() *string
	SetSound(v string) *Message2ThirdChannel
	GetSound() *string
	SetText(v string) *Message2ThirdChannel
	GetText() *string
	SetTitle(v string) *Message2ThirdChannel
	GetTitle() *string
}

type Message2ThirdChannel struct {
	SetBadge    *int64  `json:"SetBadge,omitempty" xml:"SetBadge,omitempty"`
	AddBadge    *int64  `json:"addBadge,omitempty" xml:"addBadge,omitempty"`
	BigBody     *string `json:"bigBody,omitempty" xml:"bigBody,omitempty"`
	BigTitle    *string `json:"bigTitle,omitempty" xml:"bigTitle,omitempty"`
	ExpandImage *string `json:"expandImage,omitempty" xml:"expandImage,omitempty"`
	Img         *string `json:"img,omitempty" xml:"img,omitempty"`
	Sound       *string `json:"sound,omitempty" xml:"sound,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Message2ThirdChannel) String() string {
	return dara.Prettify(s)
}

func (s Message2ThirdChannel) GoString() string {
	return s.String()
}

func (s *Message2ThirdChannel) GetSetBadge() *int64 {
	return s.SetBadge
}

func (s *Message2ThirdChannel) GetAddBadge() *int64 {
	return s.AddBadge
}

func (s *Message2ThirdChannel) GetBigBody() *string {
	return s.BigBody
}

func (s *Message2ThirdChannel) GetBigTitle() *string {
	return s.BigTitle
}

func (s *Message2ThirdChannel) GetExpandImage() *string {
	return s.ExpandImage
}

func (s *Message2ThirdChannel) GetImg() *string {
	return s.Img
}

func (s *Message2ThirdChannel) GetSound() *string {
	return s.Sound
}

func (s *Message2ThirdChannel) GetText() *string {
	return s.Text
}

func (s *Message2ThirdChannel) GetTitle() *string {
	return s.Title
}

func (s *Message2ThirdChannel) SetSetBadge(v int64) *Message2ThirdChannel {
	s.SetBadge = &v
	return s
}

func (s *Message2ThirdChannel) SetAddBadge(v int64) *Message2ThirdChannel {
	s.AddBadge = &v
	return s
}

func (s *Message2ThirdChannel) SetBigBody(v string) *Message2ThirdChannel {
	s.BigBody = &v
	return s
}

func (s *Message2ThirdChannel) SetBigTitle(v string) *Message2ThirdChannel {
	s.BigTitle = &v
	return s
}

func (s *Message2ThirdChannel) SetExpandImage(v string) *Message2ThirdChannel {
	s.ExpandImage = &v
	return s
}

func (s *Message2ThirdChannel) SetImg(v string) *Message2ThirdChannel {
	s.Img = &v
	return s
}

func (s *Message2ThirdChannel) SetSound(v string) *Message2ThirdChannel {
	s.Sound = &v
	return s
}

func (s *Message2ThirdChannel) SetText(v string) *Message2ThirdChannel {
	s.Text = &v
	return s
}

func (s *Message2ThirdChannel) SetTitle(v string) *Message2ThirdChannel {
	s.Title = &v
	return s
}

func (s *Message2ThirdChannel) Validate() error {
	return dara.Validate(s)
}
