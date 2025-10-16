// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateCdsFileResponseBody) *CreateCdsFileResponse
	GetBody() *CreateCdsFileResponseBody
}

type CreateCdsFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileResponse) GoString() string {
	return s.String()
}

func (s *CreateCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCdsFileResponse) GetBody() *CreateCdsFileResponseBody {
	return s.Body
}

func (s *CreateCdsFileResponse) SetHeaders(v map[string]*string) *CreateCdsFileResponse {
	s.Headers = v
	return s
}

func (s *CreateCdsFileResponse) SetStatusCode(v int32) *CreateCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdsFileResponse) SetBody(v *CreateCdsFileResponseBody) *CreateCdsFileResponse {
	s.Body = v
	return s
}

func (s *CreateCdsFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
