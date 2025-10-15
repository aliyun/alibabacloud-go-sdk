// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateId(v string) *GetApplicationTemplateRequest
	GetApplicationTemplateId() *string
}

type GetApplicationTemplateRequest struct {
	// 应用模板id
	//
	// This parameter is required.
	//
	// example:
	//
	// apt_ramuser_xxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
}

func (s GetApplicationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *GetApplicationTemplateRequest) SetApplicationTemplateId(v string) *GetApplicationTemplateRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
