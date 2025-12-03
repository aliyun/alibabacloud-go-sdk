// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkitemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkitemResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkitemResponseBody) *CreateWorkitemResponse
	GetBody() *CreateWorkitemResponseBody
}

type CreateWorkitemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkitemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkitemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkitemResponse) GetBody() *CreateWorkitemResponseBody {
	return s.Body
}

func (s *CreateWorkitemResponse) SetHeaders(v map[string]*string) *CreateWorkitemResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemResponse) SetStatusCode(v int32) *CreateWorkitemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemResponse) SetBody(v *CreateWorkitemResponseBody) *CreateWorkitemResponse {
	s.Body = v
	return s
}

func (s *CreateWorkitemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
