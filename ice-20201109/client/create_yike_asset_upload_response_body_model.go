// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeAssetUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateYikeAssetUploadResponseBody
	GetFileURL() *string
	SetRequestId(v string) *CreateYikeAssetUploadResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateYikeAssetUploadResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateYikeAssetUploadResponseBody
	GetUploadAuth() *string
}

type CreateYikeAssetUploadResponseBody struct {
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// eyJTZWN1cml0a2VuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// example:
	//
	// eyJFbmRwb2ludCI6Imm****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateYikeAssetUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeAssetUploadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYikeAssetUploadResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateYikeAssetUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYikeAssetUploadResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateYikeAssetUploadResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateYikeAssetUploadResponseBody) SetFileURL(v string) *CreateYikeAssetUploadResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateYikeAssetUploadResponseBody) SetRequestId(v string) *CreateYikeAssetUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYikeAssetUploadResponseBody) SetUploadAddress(v string) *CreateYikeAssetUploadResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateYikeAssetUploadResponseBody) SetUploadAuth(v string) *CreateYikeAssetUploadResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateYikeAssetUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
