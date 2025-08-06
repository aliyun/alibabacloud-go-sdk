// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadAttachedMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadAttachedMediaResponseBody
	GetFileURL() *string
	SetMediaId(v string) *CreateUploadAttachedMediaResponseBody
	GetMediaId() *string
	SetMediaURL(v string) *CreateUploadAttachedMediaResponseBody
	GetMediaURL() *string
	SetRequestId(v string) *CreateUploadAttachedMediaResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadAttachedMediaResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadAttachedMediaResponseBody
	GetUploadAuth() *string
}

type CreateUploadAttachedMediaResponseBody struct {
	// The URL of the auxiliary media asset file. The URL is an Object Storage Service (OSS) URL and does not contain the information used for URL signing.
	//
	// You can use specify this value for the `FileUrl` parameter when you call the [AddWatermark](~~AddWatermark~~) operation to create a watermark template.
	//
	// example:
	//
	// https://****.oss-cn-shanghai.aliyuncs.com/watermark/****.mov
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the auxiliary media asset.
	//
	// example:
	//
	// 97dc17a5abc3668489b84ce9****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the auxiliary media asset.
	//
	// If a domain name for Alibaba Cloud CDN is specified, a CDN URL is returned. Otherwise, an OSS URL is returned.
	//
	// >  If you enable the URL signing feature of ApsaraVideo VOD, you may be unable to access the returned URL of the auxiliary media asset by using a browser and the HTTP status code 403 may be returned. To resolve this issue, you can disable the [URL signing](https://help.aliyun.com/document_detail/86090.html) feature or [generate a signed URL](https://help.aliyun.com/document_detail/57007.html).
	//
	// example:
	//
	// http://example.aliyundoc.com/watermark/****.mov?auth_key=****
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73254DE5-F260-4720-D06856B63C01****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// >  The upload URL returned by this operation is Base64-encoded. Before you can use an SDK or an API operation to upload a media asset based on the upload URL, you must decode the upload URL by using the Base64 algorithm. You must parse the upload URL only if you use native OSS SDKs or OSS API for uploads.
	//
	// example:
	//
	// LWNuLXNoYW5naGFpLmFsaXl1b****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// >  The upload credential returned by this operation is Base64-encoded. Before you can use an SDK or an API operation to upload a media asset based on the upload credential, you must decode the upload credential by using the Base64 algorithm. You must parse the upload credential only if you use native OSS SDKs or OSS API for uploads.
	//
	// example:
	//
	// UzFnUjFxNkZ0NUIZTaklyNWJoQ00zdHF****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadAttachedMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadAttachedMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadAttachedMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadAttachedMediaResponseBody) GetMediaURL() *string {
	return s.MediaURL
}

func (s *CreateUploadAttachedMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadAttachedMediaResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadAttachedMediaResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadAttachedMediaResponseBody) SetFileURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaId(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetRequestId(v string) *CreateUploadAttachedMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAddress(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAuth(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
