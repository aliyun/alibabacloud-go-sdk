// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserPictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFacePicMd5(v string) *DeleteFaceUserPictureRequest
	GetFacePicMd5() *string
	SetIsolationId(v string) *DeleteFaceUserPictureRequest
	GetIsolationId() *string
	SetUserId(v string) *DeleteFaceUserPictureRequest
	GetUserId() *string
}

type DeleteFaceUserPictureRequest struct {
	// The unique identifier of the user\\"s face image. Set this parameter to the **Data*	- value that you obtained and saved locally after you called the [AddFaceUserPicutre](https://help.aliyun.com/document_detail/312040.html) operation to add the face image for the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1bb580ef7****
	FacePicMd5 *string `json:"FacePicMd5,omitempty" xml:"FacePicMd5,omitempty"`
	// The business isolation ID. It is used to isolate data between different business applications under the same Alibaba Cloud account. Set this parameter to the same value as the **IsolationId*	- that you specified when you called the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The user ID. Set this parameter to the **UserId*	- value that you obtained and saved locally after you called the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3mtjaqebeq4d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteFaceUserPictureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserPictureRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureRequest) GetFacePicMd5() *string {
	return s.FacePicMd5
}

func (s *DeleteFaceUserPictureRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *DeleteFaceUserPictureRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteFaceUserPictureRequest) SetFacePicMd5(v string) *DeleteFaceUserPictureRequest {
	s.FacePicMd5 = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) SetIsolationId(v string) *DeleteFaceUserPictureRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) SetUserId(v string) *DeleteFaceUserPictureRequest {
	s.UserId = &v
	return s
}

func (s *DeleteFaceUserPictureRequest) Validate() error {
	return dara.Validate(s)
}
