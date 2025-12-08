// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceImageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteFaceImageTemplateRequest
	GetTemplateId() *string
}

type DeleteFaceImageTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteFaceImageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteFaceImageTemplateRequest) SetTemplateId(v string) *DeleteFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteFaceImageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
