// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastStickerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateBroadcastStickerRequest
	GetFileName() *string
	SetOssKey(v string) *CreateBroadcastStickerRequest
	GetOssKey() *string
}

type CreateBroadcastStickerRequest struct {
	// example:
	//
	// sticker.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// material/INPUT_BROADCAST_STICKER/Mt.CPRLVQRR27YU2
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateBroadcastStickerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastStickerRequest) GoString() string {
	return s.String()
}

func (s *CreateBroadcastStickerRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateBroadcastStickerRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateBroadcastStickerRequest) SetFileName(v string) *CreateBroadcastStickerRequest {
	s.FileName = &v
	return s
}

func (s *CreateBroadcastStickerRequest) SetOssKey(v string) *CreateBroadcastStickerRequest {
	s.OssKey = &v
	return s
}

func (s *CreateBroadcastStickerRequest) Validate() error {
	return dara.Validate(s)
}
