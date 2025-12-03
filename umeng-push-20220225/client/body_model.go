// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivity(v string) *Body
	GetActivity() *string
	SetAddBadge(v int32) *Body
	GetAddBadge() *int32
	SetAfterOpen(v string) *Body
	GetAfterOpen() *string
	SetBuilderId(v int64) *Body
	GetBuilderId() *int64
	SetCustom(v string) *Body
	GetCustom() *string
	SetExpandImage(v string) *Body
	GetExpandImage() *string
	SetIcon(v string) *Body
	GetIcon() *string
	SetImg(v string) *Body
	GetImg() *string
	SetPlayLights(v bool) *Body
	GetPlayLights() *bool
	SetPlaySound(v bool) *Body
	GetPlaySound() *bool
	SetPlayVibrate(v bool) *Body
	GetPlayVibrate() *bool
	SetRePop(v int32) *Body
	GetRePop() *int32
	SetSetBadge(v int32) *Body
	GetSetBadge() *int32
	SetSound(v string) *Body
	GetSound() *string
	SetText(v string) *Body
	GetText() *string
	SetTitle(v string) *Body
	GetTitle() *string
	SetUrl(v string) *Body
	GetUrl() *string
}

type Body struct {
	Activity    *string `json:"activity,omitempty" xml:"activity,omitempty"`
	AddBadge    *int32  `json:"addBadge,omitempty" xml:"addBadge,omitempty"`
	AfterOpen   *string `json:"afterOpen,omitempty" xml:"afterOpen,omitempty"`
	BuilderId   *int64  `json:"builderId,omitempty" xml:"builderId,omitempty"`
	Custom      *string `json:"custom,omitempty" xml:"custom,omitempty"`
	ExpandImage *string `json:"expandImage,omitempty" xml:"expandImage,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Img         *string `json:"img,omitempty" xml:"img,omitempty"`
	PlayLights  *bool   `json:"playLights,omitempty" xml:"playLights,omitempty"`
	PlaySound   *bool   `json:"playSound,omitempty" xml:"playSound,omitempty"`
	PlayVibrate *bool   `json:"playVibrate,omitempty" xml:"playVibrate,omitempty"`
	RePop       *int32  `json:"rePop,omitempty" xml:"rePop,omitempty"`
	SetBadge    *int32  `json:"setBadge,omitempty" xml:"setBadge,omitempty"`
	Sound       *string `json:"sound,omitempty" xml:"sound,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s Body) String() string {
	return dara.Prettify(s)
}

func (s Body) GoString() string {
	return s.String()
}

func (s *Body) GetActivity() *string {
	return s.Activity
}

func (s *Body) GetAddBadge() *int32 {
	return s.AddBadge
}

func (s *Body) GetAfterOpen() *string {
	return s.AfterOpen
}

func (s *Body) GetBuilderId() *int64 {
	return s.BuilderId
}

func (s *Body) GetCustom() *string {
	return s.Custom
}

func (s *Body) GetExpandImage() *string {
	return s.ExpandImage
}

func (s *Body) GetIcon() *string {
	return s.Icon
}

func (s *Body) GetImg() *string {
	return s.Img
}

func (s *Body) GetPlayLights() *bool {
	return s.PlayLights
}

func (s *Body) GetPlaySound() *bool {
	return s.PlaySound
}

func (s *Body) GetPlayVibrate() *bool {
	return s.PlayVibrate
}

func (s *Body) GetRePop() *int32 {
	return s.RePop
}

func (s *Body) GetSetBadge() *int32 {
	return s.SetBadge
}

func (s *Body) GetSound() *string {
	return s.Sound
}

func (s *Body) GetText() *string {
	return s.Text
}

func (s *Body) GetTitle() *string {
	return s.Title
}

func (s *Body) GetUrl() *string {
	return s.Url
}

func (s *Body) SetActivity(v string) *Body {
	s.Activity = &v
	return s
}

func (s *Body) SetAddBadge(v int32) *Body {
	s.AddBadge = &v
	return s
}

func (s *Body) SetAfterOpen(v string) *Body {
	s.AfterOpen = &v
	return s
}

func (s *Body) SetBuilderId(v int64) *Body {
	s.BuilderId = &v
	return s
}

func (s *Body) SetCustom(v string) *Body {
	s.Custom = &v
	return s
}

func (s *Body) SetExpandImage(v string) *Body {
	s.ExpandImage = &v
	return s
}

func (s *Body) SetIcon(v string) *Body {
	s.Icon = &v
	return s
}

func (s *Body) SetImg(v string) *Body {
	s.Img = &v
	return s
}

func (s *Body) SetPlayLights(v bool) *Body {
	s.PlayLights = &v
	return s
}

func (s *Body) SetPlaySound(v bool) *Body {
	s.PlaySound = &v
	return s
}

func (s *Body) SetPlayVibrate(v bool) *Body {
	s.PlayVibrate = &v
	return s
}

func (s *Body) SetRePop(v int32) *Body {
	s.RePop = &v
	return s
}

func (s *Body) SetSetBadge(v int32) *Body {
	s.SetBadge = &v
	return s
}

func (s *Body) SetSound(v string) *Body {
	s.Sound = &v
	return s
}

func (s *Body) SetText(v string) *Body {
	s.Text = &v
	return s
}

func (s *Body) SetTitle(v string) *Body {
	s.Title = &v
	return s
}

func (s *Body) SetUrl(v string) *Body {
	s.Url = &v
	return s
}

func (s *Body) Validate() error {
	return dara.Validate(s)
}
