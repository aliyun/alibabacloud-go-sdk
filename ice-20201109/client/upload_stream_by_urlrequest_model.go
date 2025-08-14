// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStreamByURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UploadStreamByURLRequest
	GetDefinition() *string
	SetFileExtension(v string) *UploadStreamByURLRequest
	GetFileExtension() *string
	SetHDRType(v string) *UploadStreamByURLRequest
	GetHDRType() *string
	SetMediaId(v string) *UploadStreamByURLRequest
	GetMediaId() *string
	SetStreamURL(v string) *UploadStreamByURLRequest
	GetStreamURL() *string
	SetUserData(v string) *UploadStreamByURLRequest
	GetUserData() *string
}

type UploadStreamByURLRequest struct {
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
	// mp4
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
	// 411bed50018971edb60b0764a0ec6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the transcoded stream file.
	//
	// If the URL of the transcoded stream requires authentication, you must specify the authentication parameters in the stream URL and make sure that the URL can be accessed over the Internet.
	//
	// example:
	//
	// https://example.com/sample-stream.mp4
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://test.test.com"}, "Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UploadStreamByURLRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadStreamByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLRequest) GetDefinition() *string {
	return s.Definition
}

func (s *UploadStreamByURLRequest) GetFileExtension() *string {
	return s.FileExtension
}

func (s *UploadStreamByURLRequest) GetHDRType() *string {
	return s.HDRType
}

func (s *UploadStreamByURLRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UploadStreamByURLRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *UploadStreamByURLRequest) GetUserData() *string {
	return s.UserData
}

func (s *UploadStreamByURLRequest) SetDefinition(v string) *UploadStreamByURLRequest {
	s.Definition = &v
	return s
}

func (s *UploadStreamByURLRequest) SetFileExtension(v string) *UploadStreamByURLRequest {
	s.FileExtension = &v
	return s
}

func (s *UploadStreamByURLRequest) SetHDRType(v string) *UploadStreamByURLRequest {
	s.HDRType = &v
	return s
}

func (s *UploadStreamByURLRequest) SetMediaId(v string) *UploadStreamByURLRequest {
	s.MediaId = &v
	return s
}

func (s *UploadStreamByURLRequest) SetStreamURL(v string) *UploadStreamByURLRequest {
	s.StreamURL = &v
	return s
}

func (s *UploadStreamByURLRequest) SetUserData(v string) *UploadStreamByURLRequest {
	s.UserData = &v
	return s
}

func (s *UploadStreamByURLRequest) Validate() error {
	return dara.Validate(s)
}
