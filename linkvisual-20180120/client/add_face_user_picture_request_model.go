// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserPictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFacePicUrl(v string) *AddFaceUserPictureRequest
	GetFacePicUrl() *string
	SetIsolationId(v string) *AddFaceUserPictureRequest
	GetIsolationId() *string
	SetUserId(v string) *AddFaceUserPictureRequest
	GetUserId() *string
}

type AddFaceUserPictureRequest struct {
	// The download URL of the user\\"s face image. Only JPEG and JPG images are supported, and the image URL cannot contain Chinese characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test.jpeg
	FacePicUrl *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. Keep this value consistent with the **IsolationId*	- set when calling the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) API operation to add a user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The user ID. Pass in the **UserId*	- value that was obtained and saved locally after calling the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) API operation to add a user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3mtjaqebeq4d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceUserPictureRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureRequest) GetFacePicUrl() *string {
	return s.FacePicUrl
}

func (s *AddFaceUserPictureRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceUserPictureRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddFaceUserPictureRequest) SetFacePicUrl(v string) *AddFaceUserPictureRequest {
	s.FacePicUrl = &v
	return s
}

func (s *AddFaceUserPictureRequest) SetIsolationId(v string) *AddFaceUserPictureRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserPictureRequest) SetUserId(v string) *AddFaceUserPictureRequest {
	s.UserId = &v
	return s
}

func (s *AddFaceUserPictureRequest) Validate() error {
	return dara.Validate(s)
}
