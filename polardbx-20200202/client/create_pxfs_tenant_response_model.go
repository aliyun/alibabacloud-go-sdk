// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePxfsTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePxfsTenantResponse
	GetStatusCode() *int32
	SetBody(v *CreatePxfsTenantResponseBody) *CreatePxfsTenantResponse
	GetBody() *CreatePxfsTenantResponseBody
}

type CreatePxfsTenantResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePxfsTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePxfsTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsTenantResponse) GoString() string {
	return s.String()
}

func (s *CreatePxfsTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePxfsTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePxfsTenantResponse) GetBody() *CreatePxfsTenantResponseBody {
	return s.Body
}

func (s *CreatePxfsTenantResponse) SetHeaders(v map[string]*string) *CreatePxfsTenantResponse {
	s.Headers = v
	return s
}

func (s *CreatePxfsTenantResponse) SetStatusCode(v int32) *CreatePxfsTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePxfsTenantResponse) SetBody(v *CreatePxfsTenantResponseBody) *CreatePxfsTenantResponse {
	s.Body = v
	return s
}

func (s *CreatePxfsTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
