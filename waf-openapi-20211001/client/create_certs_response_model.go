// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCertsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCertsResponseBody) *CreateCertsResponse
	GetBody() *CreateCertsResponseBody
}

type CreateCertsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCertsResponse) GoString() string {
	return s.String()
}

func (s *CreateCertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCertsResponse) GetBody() *CreateCertsResponseBody {
	return s.Body
}

func (s *CreateCertsResponse) SetHeaders(v map[string]*string) *CreateCertsResponse {
	s.Headers = v
	return s
}

func (s *CreateCertsResponse) SetStatusCode(v int32) *CreateCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertsResponse) SetBody(v *CreateCertsResponseBody) *CreateCertsResponse {
	s.Body = v
	return s
}

func (s *CreateCertsResponse) Validate() error {
	return dara.Validate(s)
}
