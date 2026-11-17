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
	SetBuiltInApplicationActionType(v string) *GetApplicationTemplateRequest
	GetBuiltInApplicationActionType() *string
}

type GetApplicationTemplateRequest struct {
	// The ID of the application template to query. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The return scope for built-in application actions. Valid values:
	//
	// - **Custom**: returns only applications whose actions have been customized by the user. This is the default value.
	//
	// - **All**: returns all built-in applications.
	//
	// The priority of action effectiveness is: user customization > template default > application library default.
	//
	// example:
	//
	// Custom
	BuiltInApplicationActionType *string `json:"BuiltInApplicationActionType,omitempty" xml:"BuiltInApplicationActionType,omitempty"`
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

func (s *GetApplicationTemplateRequest) GetBuiltInApplicationActionType() *string {
	return s.BuiltInApplicationActionType
}

func (s *GetApplicationTemplateRequest) SetApplicationTemplateId(v string) *GetApplicationTemplateRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationTemplateRequest) SetBuiltInApplicationActionType(v string) *GetApplicationTemplateRequest {
	s.BuiltInApplicationActionType = &v
	return s
}

func (s *GetApplicationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
