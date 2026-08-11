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
	// The OSS URL of the file (without authentication).
	//
	// example:
	//
	// http://outin-***.oss-cn-north-2-gov-1.aliyuncs.com/sv/40360f05-181f63c3110-0004-cd8e-27f-de3c9.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media asset URL.
	//
	// > If a CDN domain name is configured, a CDN URL is returned. Otherwise, an OSS URL is returned. If the returned MediaURL is inaccessible (403) in a browser, URL signing is enabled for the VOD domain name. Disable URL signing or generate a signing signature.
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
	// The upload address.
	//
	// > The upload address returned by the operation is a Base64-encoded value. When you call an SDK or API to upload media assets, decode the value by using Base64 before use. Only uploads through the native OSS SDK or OSS API require you to parse UploadAddress.
	//
	// example:
	//
	// eyJTZWN1cml0a2VuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The upload credential returned by the operation is a Base64-encoded value. When you call an SDK or API to upload media assets, decode the value by using Base64 before use. Only uploads through the native OSS SDK or OSS API require you to parse UploadAuth.
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
