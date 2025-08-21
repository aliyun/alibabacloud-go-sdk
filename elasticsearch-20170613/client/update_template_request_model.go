// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTemplateRequest
	GetClientToken() *string
	SetBody(v string) *UpdateTemplateRequest
	GetBody() *string
}

type UpdateTemplateRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {     "content": "{\\n\\t\\"persistent\\":{\\n\\t\\t\\"search\\":{\\n\\t\\t\\t\\"max_buckets\\":\\"10000\\"\\n\\t\\t}\\n\\t}\\n}" }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTemplateRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateTemplateRequest) SetClientToken(v string) *UpdateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateRequest) SetBody(v string) *UpdateTemplateRequest {
	s.Body = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
