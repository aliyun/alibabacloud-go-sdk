// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppRecordTemplateRequest
	GetAppId() *string
	SetClientToken(v string) *DeleteAppRecordTemplateRequest
	GetClientToken() *string
	SetTemplate(v *DeleteAppRecordTemplateRequestTemplate) *DeleteAppRecordTemplateRequest
	GetTemplate() *DeleteAppRecordTemplateRequestTemplate
}

type DeleteAppRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Template *DeleteAppRecordTemplateRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DeleteAppRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppRecordTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppRecordTemplateRequest) GetTemplate() *DeleteAppRecordTemplateRequestTemplate {
	return s.Template
}

func (s *DeleteAppRecordTemplateRequest) SetAppId(v string) *DeleteAppRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRecordTemplateRequest) SetClientToken(v string) *DeleteAppRecordTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppRecordTemplateRequest) SetTemplate(v *DeleteAppRecordTemplateRequestTemplate) *DeleteAppRecordTemplateRequest {
	s.Template = v
	return s
}

func (s *DeleteAppRecordTemplateRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppRecordTemplateRequestTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// 2xh6****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAppRecordTemplateRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRecordTemplateRequestTemplate) GoString() string {
	return s.String()
}

func (s *DeleteAppRecordTemplateRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAppRecordTemplateRequestTemplate) SetTemplateId(v string) *DeleteAppRecordTemplateRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *DeleteAppRecordTemplateRequestTemplate) Validate() error {
	return dara.Validate(s)
}
