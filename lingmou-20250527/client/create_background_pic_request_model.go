// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackgroundPicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilename(v string) *CreateBackgroundPicRequest
	GetFilename() *string
	SetOssKey(v string) *CreateBackgroundPicRequest
	GetOssKey() *string
}

type CreateBackgroundPicRequest struct {
	// example:
	//
	// 1.jpg
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// example:
	//
	// material/INPUT_CHAT_BACKGROUND_PIC/Mt.CN2VNOPRC5QU2
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateBackgroundPicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackgroundPicRequest) GoString() string {
	return s.String()
}

func (s *CreateBackgroundPicRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateBackgroundPicRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateBackgroundPicRequest) SetFilename(v string) *CreateBackgroundPicRequest {
	s.Filename = &v
	return s
}

func (s *CreateBackgroundPicRequest) SetOssKey(v string) *CreateBackgroundPicRequest {
	s.OssKey = &v
	return s
}

func (s *CreateBackgroundPicRequest) Validate() error {
	return dara.Validate(s)
}
