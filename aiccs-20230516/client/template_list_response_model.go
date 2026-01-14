// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TemplateListResponse
	GetStatusCode() *int32
	SetBody(v *TemplateListResponseBody) *TemplateListResponse
	GetBody() *TemplateListResponseBody
}

type TemplateListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s TemplateListResponse) GoString() string {
	return s.String()
}

func (s *TemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TemplateListResponse) GetBody() *TemplateListResponseBody {
	return s.Body
}

func (s *TemplateListResponse) SetHeaders(v map[string]*string) *TemplateListResponse {
	s.Headers = v
	return s
}

func (s *TemplateListResponse) SetStatusCode(v int32) *TemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *TemplateListResponse) SetBody(v *TemplateListResponseBody) *TemplateListResponse {
	s.Body = v
	return s
}

func (s *TemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
