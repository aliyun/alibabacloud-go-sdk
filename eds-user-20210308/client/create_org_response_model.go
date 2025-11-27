// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrgResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrgResponseBody) *CreateOrgResponse
	GetBody() *CreateOrgResponseBody
}

type CreateOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrgResponse) GetBody() *CreateOrgResponseBody {
	return s.Body
}

func (s *CreateOrgResponse) SetHeaders(v map[string]*string) *CreateOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateOrgResponse) SetStatusCode(v int32) *CreateOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrgResponse) SetBody(v *CreateOrgResponseBody) *CreateOrgResponse {
	s.Body = v
	return s
}

func (s *CreateOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
