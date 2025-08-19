// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *DescribeFaceVerifyRequest
	GetCertifyId() *string
	SetPictureReturnType(v string) *DescribeFaceVerifyRequest
	GetPictureReturnType() *string
	SetSceneId(v int64) *DescribeFaceVerifyRequest
	GetSceneId() *int64
}

type DescribeFaceVerifyRequest struct {
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// JPG
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeFaceVerifyRequest) GetPictureReturnType() *string {
	return s.PictureReturnType
}

func (s *DescribeFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeFaceVerifyRequest) SetCertifyId(v string) *DescribeFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetPictureReturnType(v string) *DescribeFaceVerifyRequest {
	s.PictureReturnType = &v
	return s
}

func (s *DescribeFaceVerifyRequest) SetSceneId(v int64) *DescribeFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
