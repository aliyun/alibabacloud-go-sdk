// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateUploadStreamRequest
	GetDefinition() *string
	SetFileExtension(v string) *CreateUploadStreamRequest
	GetFileExtension() *string
	SetHDRType(v string) *CreateUploadStreamRequest
	GetHDRType() *string
	SetMediaId(v string) *CreateUploadStreamRequest
	GetMediaId() *string
	SetUserData(v string) *CreateUploadStreamRequest
	GetUserData() *string
}

type CreateUploadStreamRequest struct {
	// The quality of the media stream. Valid values:
	//
	// 	- FD: low definition.
	//
	// 	- LD: standard definition.
	//
	// 	- SD: high definition.
	//
	// 	- HD: ultra-high definition.
	//
	// 	- OD: original quality.
	//
	// 	- 2K: 2K resolution.
	//
	// 	- 4K: 4K resolution.
	//
	// 	- SQ: standard sound quality.
	//
	// 	- HQ: high sound quality.
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The file name extension of the media stream.
	//
	// example:
	//
	// MP4
	FileExtension *string `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	// The high dynamic range (HDR) format of the transcoded stream. Valid values:
	//
	// 	- HDR
	//
	// 	- HDR10
	//
	// 	- HLG
	//
	// 	- DolbyVision
	//
	// 	- HDRVivid
	//
	// 	- SDR+
	//
	// >
	//
	// 	- The value is not case-sensitive,
	//
	// 	- You can leave this parameter empty for non-HDR streams.
	//
	// example:
	//
	// HDR10
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://aliyundoc.com"}, "Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadStreamRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadStreamRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateUploadStreamRequest) GetFileExtension() *string {
	return s.FileExtension
}

func (s *CreateUploadStreamRequest) GetHDRType() *string {
	return s.HDRType
}

func (s *CreateUploadStreamRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadStreamRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadStreamRequest) SetDefinition(v string) *CreateUploadStreamRequest {
	s.Definition = &v
	return s
}

func (s *CreateUploadStreamRequest) SetFileExtension(v string) *CreateUploadStreamRequest {
	s.FileExtension = &v
	return s
}

func (s *CreateUploadStreamRequest) SetHDRType(v string) *CreateUploadStreamRequest {
	s.HDRType = &v
	return s
}

func (s *CreateUploadStreamRequest) SetMediaId(v string) *CreateUploadStreamRequest {
	s.MediaId = &v
	return s
}

func (s *CreateUploadStreamRequest) SetUserData(v string) *CreateUploadStreamRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadStreamRequest) Validate() error {
	return dara.Validate(s)
}
