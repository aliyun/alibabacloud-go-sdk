// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedQueryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSimpleQuery(v bool) *UpdateAdvancedQueryTemplateRequest
	GetSimpleQuery() *bool
	SetTemplateId(v string) *UpdateAdvancedQueryTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateAdvancedQueryTemplateRequest
	GetTemplateName() *string
	SetTemplateSql(v string) *UpdateAdvancedQueryTemplateRequest
	GetTemplateSql() *string
}

type UpdateAdvancedQueryTemplateRequest struct {
	// Specifies whether to enable the simple query mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// utpl-QNL3dpYkQcyjZxrIQC****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template. The maximum length is 64 characters.
	//
	// example:
	//
	// example-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The query statement of the template.
	//
	// example:
	//
	// event.eventName: ConsoleSignin AND event.userIdentity.type: root-account
	TemplateSql *string `json:"TemplateSql,omitempty" xml:"TemplateSql,omitempty"`
}

func (s UpdateAdvancedQueryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedQueryTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedQueryTemplateRequest) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *UpdateAdvancedQueryTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateAdvancedQueryTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateAdvancedQueryTemplateRequest) GetTemplateSql() *string {
	return s.TemplateSql
}

func (s *UpdateAdvancedQueryTemplateRequest) SetSimpleQuery(v bool) *UpdateAdvancedQueryTemplateRequest {
	s.SimpleQuery = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateRequest) SetTemplateId(v string) *UpdateAdvancedQueryTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateRequest) SetTemplateName(v string) *UpdateAdvancedQueryTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateRequest) SetTemplateSql(v string) *UpdateAdvancedQueryTemplateRequest {
	s.TemplateSql = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
