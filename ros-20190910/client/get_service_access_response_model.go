// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceAccessResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceAccessResponseBody) *GetServiceAccessResponse
	GetBody() *GetServiceAccessResponseBody
}

type GetServiceAccessResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessResponse) GoString() string {
	return s.String()
}

func (s *GetServiceAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceAccessResponse) GetBody() *GetServiceAccessResponseBody {
	return s.Body
}

func (s *GetServiceAccessResponse) SetHeaders(v map[string]*string) *GetServiceAccessResponse {
	s.Headers = v
	return s
}

func (s *GetServiceAccessResponse) SetStatusCode(v int32) *GetServiceAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceAccessResponse) SetBody(v *GetServiceAccessResponseBody) *GetServiceAccessResponse {
	s.Body = v
	return s
}

func (s *GetServiceAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
