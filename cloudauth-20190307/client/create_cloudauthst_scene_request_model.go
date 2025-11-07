// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudauthstSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *CreateCloudauthstSceneRequest
	GetProductCode() *string
	SetSceneName(v string) *CreateCloudauthstSceneRequest
	GetSceneName() *string
	SetStoreImage(v string) *CreateCloudauthstSceneRequest
	GetStoreImage() *string
}

type CreateCloudauthstSceneRequest struct {
	// Product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMART_COMPARE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Scene name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试场景
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Whether to deliver the files generated from the authentication to the customer\\"s OSS:
	//
	// - **Y**: Enable
	//
	// - **N**: Disable
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
}

func (s CreateCloudauthstSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudauthstSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudauthstSceneRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateCloudauthstSceneRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *CreateCloudauthstSceneRequest) GetStoreImage() *string {
	return s.StoreImage
}

func (s *CreateCloudauthstSceneRequest) SetProductCode(v string) *CreateCloudauthstSceneRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCloudauthstSceneRequest) SetSceneName(v string) *CreateCloudauthstSceneRequest {
	s.SceneName = &v
	return s
}

func (s *CreateCloudauthstSceneRequest) SetStoreImage(v string) *CreateCloudauthstSceneRequest {
	s.StoreImage = &v
	return s
}

func (s *CreateCloudauthstSceneRequest) Validate() error {
	return dara.Validate(s)
}
