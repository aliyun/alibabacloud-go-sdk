// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppTemplateLikeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiked(v bool) *OperateAppTemplateLikeRequest
	GetLiked() *bool
	SetTemplateId(v string) *OperateAppTemplateLikeRequest
	GetTemplateId() *string
}

type OperateAppTemplateLikeRequest struct {
	Liked *bool `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// example:
	//
	// K191WHV12URYQN06
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s OperateAppTemplateLikeRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAppTemplateLikeRequest) GoString() string {
	return s.String()
}

func (s *OperateAppTemplateLikeRequest) GetLiked() *bool {
	return s.Liked
}

func (s *OperateAppTemplateLikeRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *OperateAppTemplateLikeRequest) SetLiked(v bool) *OperateAppTemplateLikeRequest {
	s.Liked = &v
	return s
}

func (s *OperateAppTemplateLikeRequest) SetTemplateId(v string) *OperateAppTemplateLikeRequest {
	s.TemplateId = &v
	return s
}

func (s *OperateAppTemplateLikeRequest) Validate() error {
	return dara.Validate(s)
}
