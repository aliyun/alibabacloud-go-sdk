// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyIntelligentCaptchaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptchaVerifyParam(v string) *VerifyIntelligentCaptchaRequest
	GetCaptchaVerifyParam() *string
	SetSceneId(v string) *VerifyIntelligentCaptchaRequest
	GetSceneId() *string
}

type VerifyIntelligentCaptchaRequest struct {
	// example:
	//
	// dsjidsjidsjkds*djsjdiskds
	CaptchaVerifyParam *string `json:"CaptchaVerifyParam,omitempty" xml:"CaptchaVerifyParam,omitempty"`
	SceneId            *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s VerifyIntelligentCaptchaRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyIntelligentCaptchaRequest) GoString() string {
	return s.String()
}

func (s *VerifyIntelligentCaptchaRequest) GetCaptchaVerifyParam() *string {
	return s.CaptchaVerifyParam
}

func (s *VerifyIntelligentCaptchaRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *VerifyIntelligentCaptchaRequest) SetCaptchaVerifyParam(v string) *VerifyIntelligentCaptchaRequest {
	s.CaptchaVerifyParam = &v
	return s
}

func (s *VerifyIntelligentCaptchaRequest) SetSceneId(v string) *VerifyIntelligentCaptchaRequest {
	s.SceneId = &v
	return s
}

func (s *VerifyIntelligentCaptchaRequest) Validate() error {
	return dara.Validate(s)
}
