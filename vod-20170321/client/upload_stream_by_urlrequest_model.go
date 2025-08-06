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
	SetUploadMetadata(v string) *UploadStreamByURLRequest
	GetUploadMetadata() *string
	SetUserData(v string) *UploadStreamByURLRequest
	GetUserData() *string
}

type UploadStreamByURLRequest struct {
	// The quality of the video stream.
	//
	// For more information about valid values of this parameter, see [Parameters for media assets](https://help.aliyun.com/document_detail/124671.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The file name extension of the transcoded stream.
	//
	// For more information, see the Supported media file formats section in [Overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// If you set a value for this parameter, the file name extension specified in StreamURL is overwritten.
	//
	// >  This parameter is required if you do not specify a file name extension in StreamURL.
	//
	// example:
	//
	// mp4
	FileExtension *string `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	// The HDR type of the transcoded stream. Valid values:
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
	// 	- The HDR type of the transcoded stream is not case-sensitive.
	//
	// 	- You can leave this parameter empty for non-HDR streams.
	//
	// example:
	//
	// HDR10
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The media ID in ApsaraVideo VOD.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca3a8f6e49*****57b65806709586
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the transcoded stream.
	//
	// If URL authentication is required, you must pass authentication information in this parameter and make sure that the URL can be accessed over the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/lesson-01.mp4
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// Metadata information for uploading media files, in JSON string format.
	//
	// For more information, please refer to the table below for UploadMetadata.
	//
	// example:
	//
	// {"AddressMapping":"1","CustomPath":"test/xxx","CustomFileName":"xxx.mp4","isOverwritePath":"0"}
	UploadMetadata *string `json:"UploadMetadata,omitempty" xml:"UploadMetadata,omitempty"`
	// The user-defined parameter. For more information, see the "UserData: specifies the custom configurations for media upload" section of the [Request parameters](https://help.aliyun.com/document_detail/86952.html) topic.
	//
	// >  The callback configurations you specify for this parameter take effect only after you specify the HTTP callback URL and select specific callback events in the ApsaraVideo VOD console. For more information about how to configure HTTP callback settings in the ApsaraVideo VOD console, see [Configure callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://aliyundoc.com"}, "Extend":{"localId":"xxx","test":"www"}}
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

func (s *UploadStreamByURLRequest) GetUploadMetadata() *string {
	return s.UploadMetadata
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

func (s *UploadStreamByURLRequest) SetUploadMetadata(v string) *UploadStreamByURLRequest {
	s.UploadMetadata = &v
	return s
}

func (s *UploadStreamByURLRequest) SetUserData(v string) *UploadStreamByURLRequest {
	s.UserData = &v
	return s
}

func (s *UploadStreamByURLRequest) Validate() error {
	return dara.Validate(s)
}
