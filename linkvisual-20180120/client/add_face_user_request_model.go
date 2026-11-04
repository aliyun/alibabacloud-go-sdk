// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomUserId(v string) *AddFaceUserRequest
	GetCustomUserId() *string
	SetFacePicUrl(v string) *AddFaceUserRequest
	GetFacePicUrl() *string
	SetIsolationId(v string) *AddFaceUserRequest
	GetIsolationId() *string
	SetName(v string) *AddFaceUserRequest
	GetName() *string
	SetParams(v string) *AddFaceUserRequest
	GetParams() *string
}

type AddFaceUserRequest struct {
	// The user identification ID. It supports any uppercase letter and digits. Recommended Settings: the user\\"s ID card number, student ID, or similar identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The URL for downloading the user\\"s face image. Only JPEG and JPG image formats are supported, and the image URL cannot contain Chinese characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test.jpeg
	FacePicUrl *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	// The business isolation ID. It is used for data isolation between different business applications under the same Alibaba Cloud account. It supports uppercase and lowercase English letters and digits.
	//
	// For example, if you have face recognition businesses in Region A and Region B, you can set different business isolation IDs to distinguish the face recognition businesses in Region A and Region B.
	//
	// 	Notice:
	//
	// Save the business isolation ID properly for use in subsequent calls to other API operations in this business isolation zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The username. It supports Chinese characters, uppercase and lowercase English letters, digits, underscores (_), and hyphens (-). The length cannot exceed 64 characters.
	//
	// example:
	//
	// Tony
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom user parameters.
	//
	// example:
	//
	// {"key1":"value1", "key2":"value2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s AddFaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserRequest) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *AddFaceUserRequest) GetFacePicUrl() *string {
	return s.FacePicUrl
}

func (s *AddFaceUserRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceUserRequest) GetName() *string {
	return s.Name
}

func (s *AddFaceUserRequest) GetParams() *string {
	return s.Params
}

func (s *AddFaceUserRequest) SetCustomUserId(v string) *AddFaceUserRequest {
	s.CustomUserId = &v
	return s
}

func (s *AddFaceUserRequest) SetFacePicUrl(v string) *AddFaceUserRequest {
	s.FacePicUrl = &v
	return s
}

func (s *AddFaceUserRequest) SetIsolationId(v string) *AddFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserRequest) SetName(v string) *AddFaceUserRequest {
	s.Name = &v
	return s
}

func (s *AddFaceUserRequest) SetParams(v string) *AddFaceUserRequest {
	s.Params = &v
	return s
}

func (s *AddFaceUserRequest) Validate() error {
	return dara.Validate(s)
}
