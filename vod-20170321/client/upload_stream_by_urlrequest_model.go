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
	// The definition of the video stream.
	//
	// For valid values of this parameter, see [Media asset parameter description - Definition](https://help.aliyun.com/document_detail/124671.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The file name extension of the transcoded stream file.
	//
	// For supported audio and video file formats, see [Overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// If this parameter is not empty, it overwrites the file name extension in the StreamURL.
	//
	// 	Notice: This parameter is required if the StreamURL does not contain a file name extension.
	//
	// example:
	//
	// mp4
	FileExtension *string `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	// The HDR type of the transcoded stream. Valid values:
	//
	// - HDR
	//
	// - HDR10
	//
	// - HLG
	//
	// - DolbyVision
	//
	// - HDRVivid
	//
	// - SDR+
	//
	// > - Case-insensitive.
	//
	// > - Leave this parameter empty for non-HDR videos.
	//
	// example:
	//
	// HDR10
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The ID of the ApsaraVideo VOD media asset that corresponds to the transcoded stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca3a8f6e49*****57b65806709586
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the transcoded stream file.
	//
	// If the URL of the transcoded stream requires authentication, include the authentication parameters in StreamURL and make sure the URL is accessible through public network access.
	//
	// >You can obtain the audio or video URL from the console or by invoking the GetPlayInfo operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/lesson-01.mp4
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The metadata of the media file to upload. The value is a JSON string.
	//
	// - For more information, see the **UploadMetadata*	- table below.
	//
	// example:
	//
	// {"AddressMapping":"1","CustomPath":"test/xxx","CustomFileName":"xxx.mp4","isOverwritePath":"0"}
	UploadMetadata *string `json:"UploadMetadata,omitempty" xml:"UploadMetadata,omitempty"`
	// The custom parameter. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// > To use the message callback in this parameter, configure the HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. For information about how to configure HTTP callbacks in the console, see [Callback settings](https://help.aliyun.com/document_detail/86071.html).
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
