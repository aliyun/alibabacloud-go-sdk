// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedQueryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetAdvancedQueryTemplateRequest
	GetTemplateId() *string
}

type GetAdvancedQueryTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// utpl-N9fpjnFBSWauSXhVNP3erw
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAdvancedQueryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedQueryTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAdvancedQueryTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAdvancedQueryTemplateRequest) SetTemplateId(v string) *GetAdvancedQueryTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAdvancedQueryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
