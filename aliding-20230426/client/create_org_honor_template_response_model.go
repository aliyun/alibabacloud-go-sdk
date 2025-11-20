// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrgHonorTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrgHonorTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrgHonorTemplateResponseBody) *CreateOrgHonorTemplateResponse
	GetBody() *CreateOrgHonorTemplateResponseBody
}

type CreateOrgHonorTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrgHonorTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrgHonorTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrgHonorTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrgHonorTemplateResponse) GetBody() *CreateOrgHonorTemplateResponseBody {
	return s.Body
}

func (s *CreateOrgHonorTemplateResponse) SetHeaders(v map[string]*string) *CreateOrgHonorTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateOrgHonorTemplateResponse) SetStatusCode(v int32) *CreateOrgHonorTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrgHonorTemplateResponse) SetBody(v *CreateOrgHonorTemplateResponseBody) *CreateOrgHonorTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateOrgHonorTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
