// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustrationTaskCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundType(v int32) *IllustrationTaskCreateCmd
	GetBackgroundType() *int32
	SetDstHeight(v int32) *IllustrationTaskCreateCmd
	GetDstHeight() *int32
	SetDstWidth(v int32) *IllustrationTaskCreateCmd
	GetDstWidth() *int32
	SetIdempotentId(v string) *IllustrationTaskCreateCmd
	GetIdempotentId() *string
	SetImageUrls(v []*string) *IllustrationTaskCreateCmd
	GetImageUrls() []*string
	SetNums(v int32) *IllustrationTaskCreateCmd
	GetNums() *int32
	SetOssPaths(v []*string) *IllustrationTaskCreateCmd
	GetOssPaths() []*string
	SetStickerText(v string) *IllustrationTaskCreateCmd
	GetStickerText() *string
}

type IllustrationTaskCreateCmd struct {
	// example:
	//
	// 0-不换背景，1-换背景
	BackgroundType *int32 `json:"backgroundType,omitempty" xml:"backgroundType,omitempty"`
	// example:
	//
	// 1024
	DstHeight *int32 `json:"dstHeight,omitempty" xml:"dstHeight,omitempty"`
	// example:
	//
	// 1024
	DstWidth *int32 `json:"dstWidth,omitempty" xml:"dstWidth,omitempty"`
	// example:
	//
	// 28274623764834
	IdempotentId *string   `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ImageUrls    []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	Nums        *int32    `json:"nums,omitempty" xml:"nums,omitempty"`
	OssPaths    []*string `json:"ossPaths,omitempty" xml:"ossPaths,omitempty" type:"Repeated"`
	StickerText *string   `json:"stickerText,omitempty" xml:"stickerText,omitempty"`
}

func (s IllustrationTaskCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s IllustrationTaskCreateCmd) GoString() string {
	return s.String()
}

func (s *IllustrationTaskCreateCmd) GetBackgroundType() *int32 {
	return s.BackgroundType
}

func (s *IllustrationTaskCreateCmd) GetDstHeight() *int32 {
	return s.DstHeight
}

func (s *IllustrationTaskCreateCmd) GetDstWidth() *int32 {
	return s.DstWidth
}

func (s *IllustrationTaskCreateCmd) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *IllustrationTaskCreateCmd) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *IllustrationTaskCreateCmd) GetNums() *int32 {
	return s.Nums
}

func (s *IllustrationTaskCreateCmd) GetOssPaths() []*string {
	return s.OssPaths
}

func (s *IllustrationTaskCreateCmd) GetStickerText() *string {
	return s.StickerText
}

func (s *IllustrationTaskCreateCmd) SetBackgroundType(v int32) *IllustrationTaskCreateCmd {
	s.BackgroundType = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetDstHeight(v int32) *IllustrationTaskCreateCmd {
	s.DstHeight = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetDstWidth(v int32) *IllustrationTaskCreateCmd {
	s.DstWidth = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetIdempotentId(v string) *IllustrationTaskCreateCmd {
	s.IdempotentId = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetImageUrls(v []*string) *IllustrationTaskCreateCmd {
	s.ImageUrls = v
	return s
}

func (s *IllustrationTaskCreateCmd) SetNums(v int32) *IllustrationTaskCreateCmd {
	s.Nums = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetOssPaths(v []*string) *IllustrationTaskCreateCmd {
	s.OssPaths = v
	return s
}

func (s *IllustrationTaskCreateCmd) SetStickerText(v string) *IllustrationTaskCreateCmd {
	s.StickerText = &v
	return s
}

func (s *IllustrationTaskCreateCmd) Validate() error {
	return dara.Validate(s)
}
