// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadStreamResponseBody
	GetFileURL() *string
	SetMediaId(v string) *CreateUploadStreamResponseBody
	GetMediaId() *string
	SetRequestId(v string) *CreateUploadStreamResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadStreamResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadStreamResponseBody
	GetUploadAuth() *string
}

type CreateUploadStreamResponseBody struct {
	// The Object Storage Service (OSS) URL of the file. The URL does not contain the information used for authentication.
	//
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// >  The returned upload URL is a Base64-encoded URL. You must decode the Base64-encoded upload URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAddress only if you use OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJFbmRwb2ludCI6Imh0dHBzOi8vb3NzLWNuLXNoYW5naGFpLmFsaXl1bmNzLmNvbSIsIkJ1Y2tldCI6InN6aGQtdmlkZW8iLCJGaWxlTmFtZSI6InZvZC0yOTYzMWEvc3YvNTBmYTJlODQtMTgxMjdhZGRiMTcvNTBmYTJlODQtMTgxMjdhZGRiM***
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// >  The returned upload credential is a Base64-encoded value. You must decode the Base64-encoded upload URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAuth only if you use OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJBY2Nlc3NLZXlJZCI6IkxUQUk0Rm53bTk1dHdxQjMxR3IzSE5hRCIsIkFjY2Vzc0tleVNlY3JldCI6Ik9lWllKR0dTMTlkNkZaM1E3UVpJQmdmSVdnM3BPaiIsIkV4cGlyYXRpb24iOiI***
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadStreamResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadStreamResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadStreamResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadStreamResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadStreamResponseBody) SetFileURL(v string) *CreateUploadStreamResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadStreamResponseBody) SetMediaId(v string) *CreateUploadStreamResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadStreamResponseBody) SetRequestId(v string) *CreateUploadStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadStreamResponseBody) SetUploadAddress(v string) *CreateUploadStreamResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadStreamResponseBody) SetUploadAuth(v string) *CreateUploadStreamResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
