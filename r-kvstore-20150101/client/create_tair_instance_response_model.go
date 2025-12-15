// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTairInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTairInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateTairInstanceResponseBody) *CreateTairInstanceResponse
	GetBody() *CreateTairInstanceResponseBody
}

type CreateTairInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTairInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTairInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTairInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTairInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTairInstanceResponse) GetBody() *CreateTairInstanceResponseBody {
	return s.Body
}

func (s *CreateTairInstanceResponse) SetHeaders(v map[string]*string) *CreateTairInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateTairInstanceResponse) SetStatusCode(v int32) *CreateTairInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTairInstanceResponse) SetBody(v *CreateTairInstanceResponseBody) *CreateTairInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateTairInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
