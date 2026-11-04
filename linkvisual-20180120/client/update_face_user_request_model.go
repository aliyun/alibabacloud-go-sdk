// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomUserId(v string) *UpdateFaceUserRequest
	GetCustomUserId() *string
	SetFacePicUrl(v string) *UpdateFaceUserRequest
	GetFacePicUrl() *string
	SetIsolationId(v string) *UpdateFaceUserRequest
	GetIsolationId() *string
	SetName(v string) *UpdateFaceUserRequest
	GetName() *string
	SetParams(v string) *UpdateFaceUserRequest
	GetParams() *string
	SetUserId(v string) *UpdateFaceUserRequest
	GetUserId() *string
}

type UpdateFaceUserRequest struct {
	// The unique ID of the user in the caller\\"s business system, such as an ID card number or student ID.
	//
	// example:
	//
	// 52242819650301xxxx
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The image download URL. Currently, JPG and JPEG formats are supported.
	//
	// example:
	//
	// https://test.com/test.jpg
	FacePicUrl *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	// The isolation ID.<br>It is used for data isolation between different applications under a tenant. For example, if a user has business operations in both Zhejiang and Shaanxi, you can create two isolation IDs for the operations in Zhejiang and Xi\\"an respectively.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolati****
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The username.
	//
	// example:
	//
	// ZhangSan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user information.
	//
	// example:
	//
	// {"age":28, "sex":"male"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tibce3fsgqel****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateFaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserRequest) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *UpdateFaceUserRequest) GetFacePicUrl() *string {
	return s.FacePicUrl
}

func (s *UpdateFaceUserRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *UpdateFaceUserRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFaceUserRequest) GetParams() *string {
	return s.Params
}

func (s *UpdateFaceUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateFaceUserRequest) SetCustomUserId(v string) *UpdateFaceUserRequest {
	s.CustomUserId = &v
	return s
}

func (s *UpdateFaceUserRequest) SetFacePicUrl(v string) *UpdateFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *UpdateFaceUserRequest) SetIsolationId(v string) *UpdateFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserRequest) SetName(v string) *UpdateFaceUserRequest {
	s.Name = &v
	return s
}

func (s *UpdateFaceUserRequest) SetParams(v string) *UpdateFaceUserRequest {
	s.Params = &v
	return s
}

func (s *UpdateFaceUserRequest) SetUserId(v string) *UpdateFaceUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateFaceUserRequest) Validate() error {
	return dara.Validate(s)
}
