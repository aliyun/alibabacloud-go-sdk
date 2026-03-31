// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTemplateAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTemplateAttributesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTemplateAttributesResponseBody) *UpdateTemplateAttributesResponse
	GetBody() *UpdateTemplateAttributesResponseBody
}

type UpdateTemplateAttributesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTemplateAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTemplateAttributesResponse) GetBody() *UpdateTemplateAttributesResponseBody {
	return s.Body
}

func (s *UpdateTemplateAttributesResponse) SetHeaders(v map[string]*string) *UpdateTemplateAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateAttributesResponse) SetStatusCode(v int32) *UpdateTemplateAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateAttributesResponse) SetBody(v *UpdateTemplateAttributesResponseBody) *UpdateTemplateAttributesResponse {
	s.Body = v
	return s
}

func (s *UpdateTemplateAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
