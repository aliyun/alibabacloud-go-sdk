// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelatedMediaidFlag(v string) *GetTemplateRequest
	GetRelatedMediaidFlag() *string
	SetTemplateId(v string) *GetTemplateRequest
	GetTemplateId() *string
}

type GetTemplateRequest struct {
	// Specifies whether to return the information about the associated materials. Default value: 0. Valid values: 0 and 1. A value of 1 specifies that the information about the associated materials is returned. This parameter is valid only for regular templates.
	//
	// example:
	//
	// 0
	RelatedMediaidFlag *string `json:"RelatedMediaidFlag,omitempty" xml:"RelatedMediaidFlag,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetRelatedMediaidFlag() *string {
	return s.RelatedMediaidFlag
}

func (s *GetTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateRequest) SetRelatedMediaidFlag(v string) *GetTemplateRequest {
	s.RelatedMediaidFlag = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
