// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVideoTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteFaceVideoTemplateRequest
	GetTemplateId() *string
}

type DeleteFaceVideoTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteFaceVideoTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteFaceVideoTemplateRequest) SetTemplateId(v string) *DeleteFaceVideoTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteFaceVideoTemplateRequest) Validate() error {
	return dara.Validate(s)
}
