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
	// The OSS URL of the file, without authentication parameters.
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
	// > This will be a CDN URL if a CDN domain is configured, or an OSS URL otherwise. If you receive a 403 error when accessing this URL in a browser, it is likely because URL authentication is enabled for the VOD domain. To resolve this, either disable URL authentication or generate a signed URL for access.
	//
	// example:
	//
	// https://xxq-live-playback.oss-cn-shanghai.aliyuncs.com/capture/5d96d2b4-111b-4e5d-a0e5-20f44405bb55.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload address.
	//
	// > The returned upload address is Base64-encoded and must be decoded before use. You only need to manually decode this address if you are using a native OSS SDK or an OSS API to perform the upload.
	//
	// example:
	//
	// eyJTZWN1cml0a2VuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The returned upload credential is Base64-encoded and must be decoded before use. You only need to manually decode this credential if you are using a native OSS SDK or an OSS API to perform the upload.
	//
	// example:
	//
	// eyJFbmRwb2ludCI6Imm****
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
