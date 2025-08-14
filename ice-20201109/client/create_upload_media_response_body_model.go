// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadMediaResponseBody
	GetFileURL() *string
	SetMediaId(v string) *CreateUploadMediaResponseBody
	GetMediaId() *string
	SetMediaURL(v string) *CreateUploadMediaResponseBody
	GetMediaURL() *string
	SetRequestId(v string) *CreateUploadMediaResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadMediaResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadMediaResponseBody
	GetUploadAuth() *string
}

type CreateUploadMediaResponseBody struct {
	// The OSS URL of the file. The URL does not contain the information used for authentication.
	//
	// example:
	//
	// http://outin-***.oss-cn-north-2-gov-1.aliyuncs.com/sv/40360f05-181f63c3110-0004-cd8e-27f-de3c9.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the media asset.
	//
	// >  If a domain name for Alibaba Cloud CDN (CDN) is specified, a CDN URL is returned. Otherwise, an OSS URL is returned. If the HTTP status code 403 is returned when you access the URL from your browser, the URL authentication feature of ApsaraVideo VOD is enabled. To resolve this issue, disable URL authentication or generate an authentication signature.
	//
	// example:
	//
	// https://xxq-live-playback.oss-cn-shanghai.aliyuncs.com/capture/5d96d2b4-111b-4e5d-a0e5-20f44405bb55.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
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

func (s CreateUploadMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadMediaResponseBody) GetMediaURL() *string {
	return s.MediaURL
}

func (s *CreateUploadMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadMediaResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadMediaResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadMediaResponseBody) SetFileURL(v string) *CreateUploadMediaResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetMediaId(v string) *CreateUploadMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetMediaURL(v string) *CreateUploadMediaResponseBody {
	s.MediaURL = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetRequestId(v string) *CreateUploadMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetUploadAddress(v string) *CreateUploadMediaResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetUploadAuth(v string) *CreateUploadMediaResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
