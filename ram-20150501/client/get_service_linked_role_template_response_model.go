// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceLinkedRoleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceLinkedRoleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceLinkedRoleTemplateResponseBody) *GetServiceLinkedRoleTemplateResponse
	GetBody() *GetServiceLinkedRoleTemplateResponseBody
}

type GetServiceLinkedRoleTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLinkedRoleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLinkedRoleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceLinkedRoleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceLinkedRoleTemplateResponse) GetBody() *GetServiceLinkedRoleTemplateResponseBody {
	return s.Body
}

func (s *GetServiceLinkedRoleTemplateResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponse) SetStatusCode(v int32) *GetServiceLinkedRoleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponse) SetBody(v *GetServiceLinkedRoleTemplateResponseBody) *GetServiceLinkedRoleTemplateResponse {
	s.Body = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
