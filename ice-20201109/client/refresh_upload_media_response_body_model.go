// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *RefreshUploadMediaResponseBody
	GetMediaId() *string
	SetRequestId(v string) *RefreshUploadMediaResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *RefreshUploadMediaResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *RefreshUploadMediaResponseBody
	GetUploadAuth() *string
}

type RefreshUploadMediaResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// c2e77390f75271ec802f0674a2ce6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// >  The returned upload URL is a Base64-encoded URL. You must decode the Base64-encoded upload URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAddress only if you use Object Storage Service (OSS) SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJFbmRwb2ludCI6Imh0dHBzOi8vb3NzLWNuLXNoYW5naGFpLmFsaXl1bmNzLmNvbSIsIkJ1Y2tldCI6InN6aGQtdmlkZW8iLCJGaWxlTmFtZSI6InZvZC0yOTYzMWEvc3YvNTBmYTJlODQtMTgxMjdhZGRiMTcvNTBmYTJlODQtMTgxMjdhZGRiM***
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// >  The returned upload credential is a Base64-encoded value. You must decode the Base64-encoded upload credential before you use an SDK or call an API operation to upload media files. You need to parse UploadAuth only if you use OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJBY2Nlc3NLZXlJZCI6IkxUQUk0Rm53bTk1dHdxQjMxR3IzSE5hRCIsIkFjY2Vzc0tleVNlY3JldCI6Ik9lWllKR0dTMTlkNkZaM1E3UVpJQmdmSVdnM3BPaiIsIkV4cGlyYXRpb24iOiI***
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s RefreshUploadMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadMediaResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshUploadMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *RefreshUploadMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshUploadMediaResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *RefreshUploadMediaResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *RefreshUploadMediaResponseBody) SetMediaId(v string) *RefreshUploadMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *RefreshUploadMediaResponseBody) SetRequestId(v string) *RefreshUploadMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshUploadMediaResponseBody) SetUploadAddress(v string) *RefreshUploadMediaResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *RefreshUploadMediaResponseBody) SetUploadAuth(v string) *RefreshUploadMediaResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *RefreshUploadMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
