// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHarmonyBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *HarmonyBody
	GetAction() *string
	SetAddBadge(v int32) *HarmonyBody
	GetAddBadge() *int32
	SetAfterOpen(v string) *HarmonyBody
	GetAfterOpen() *string
	SetBigBody(v string) *HarmonyBody
	GetBigBody() *string
	SetCustom(v string) *HarmonyBody
	GetCustom() *string
	SetImg(v string) *HarmonyBody
	GetImg() *string
	SetLargeIcon(v string) *HarmonyBody
	GetLargeIcon() *string
	SetText(v string) *HarmonyBody
	GetText() *string
	SetTitle(v string) *HarmonyBody
	GetTitle() *string
	SetUri(v string) *HarmonyBody
	GetUri() *string
}

type HarmonyBody struct {
	Action    *string `json:"action,omitempty" xml:"action,omitempty"`
	AddBadge  *int32  `json:"addBadge,omitempty" xml:"addBadge,omitempty"`
	AfterOpen *string `json:"afterOpen,omitempty" xml:"afterOpen,omitempty"`
	BigBody   *string `json:"bigBody,omitempty" xml:"bigBody,omitempty"`
	Custom    *string `json:"custom,omitempty" xml:"custom,omitempty"`
	Img       *string `json:"img,omitempty" xml:"img,omitempty"`
	LargeIcon *string `json:"largeIcon,omitempty" xml:"largeIcon,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Uri       *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s HarmonyBody) String() string {
	return dara.Prettify(s)
}

func (s HarmonyBody) GoString() string {
	return s.String()
}

func (s *HarmonyBody) GetAction() *string {
	return s.Action
}

func (s *HarmonyBody) GetAddBadge() *int32 {
	return s.AddBadge
}

func (s *HarmonyBody) GetAfterOpen() *string {
	return s.AfterOpen
}

func (s *HarmonyBody) GetBigBody() *string {
	return s.BigBody
}

func (s *HarmonyBody) GetCustom() *string {
	return s.Custom
}

func (s *HarmonyBody) GetImg() *string {
	return s.Img
}

func (s *HarmonyBody) GetLargeIcon() *string {
	return s.LargeIcon
}

func (s *HarmonyBody) GetText() *string {
	return s.Text
}

func (s *HarmonyBody) GetTitle() *string {
	return s.Title
}

func (s *HarmonyBody) GetUri() *string {
	return s.Uri
}

func (s *HarmonyBody) SetAction(v string) *HarmonyBody {
	s.Action = &v
	return s
}

func (s *HarmonyBody) SetAddBadge(v int32) *HarmonyBody {
	s.AddBadge = &v
	return s
}

func (s *HarmonyBody) SetAfterOpen(v string) *HarmonyBody {
	s.AfterOpen = &v
	return s
}

func (s *HarmonyBody) SetBigBody(v string) *HarmonyBody {
	s.BigBody = &v
	return s
}

func (s *HarmonyBody) SetCustom(v string) *HarmonyBody {
	s.Custom = &v
	return s
}

func (s *HarmonyBody) SetImg(v string) *HarmonyBody {
	s.Img = &v
	return s
}

func (s *HarmonyBody) SetLargeIcon(v string) *HarmonyBody {
	s.LargeIcon = &v
	return s
}

func (s *HarmonyBody) SetText(v string) *HarmonyBody {
	s.Text = &v
	return s
}

func (s *HarmonyBody) SetTitle(v string) *HarmonyBody {
	s.Title = &v
	return s
}

func (s *HarmonyBody) SetUri(v string) *HarmonyBody {
	s.Uri = &v
	return s
}

func (s *HarmonyBody) Validate() error {
	return dara.Validate(s)
}
