// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetAppTemplateRequest
	GetTemplateId() *string
}

type GetAppTemplateRequest struct {
	// example:
	//
	// JKSVNY04LH7DRI6F
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAppTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAppTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAppTemplateRequest) SetTemplateId(v string) *GetAppTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAppTemplateRequest) Validate() error {
	return dara.Validate(s)
}
