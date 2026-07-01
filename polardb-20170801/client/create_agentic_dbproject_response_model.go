// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticDBProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticDBProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticDBProjectResponseBody) *CreateAgenticDBProjectResponse
	GetBody() *CreateAgenticDBProjectResponseBody
}

type CreateAgenticDBProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticDBProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticDBProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticDBProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticDBProjectResponse) GetBody() *CreateAgenticDBProjectResponseBody {
	return s.Body
}

func (s *CreateAgenticDBProjectResponse) SetHeaders(v map[string]*string) *CreateAgenticDBProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticDBProjectResponse) SetStatusCode(v int32) *CreateAgenticDBProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticDBProjectResponse) SetBody(v *CreateAgenticDBProjectResponseBody) *CreateAgenticDBProjectResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticDBProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
