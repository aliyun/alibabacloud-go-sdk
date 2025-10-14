// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSDGResponse
	GetStatusCode() *int32
	SetBody(v *CreateSDGResponseBody) *CreateSDGResponse
	GetBody() *CreateSDGResponseBody
}

type CreateSDGResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSDGResponse) GoString() string {
	return s.String()
}

func (s *CreateSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSDGResponse) GetBody() *CreateSDGResponseBody {
	return s.Body
}

func (s *CreateSDGResponse) SetHeaders(v map[string]*string) *CreateSDGResponse {
	s.Headers = v
	return s
}

func (s *CreateSDGResponse) SetStatusCode(v int32) *CreateSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSDGResponse) SetBody(v *CreateSDGResponseBody) *CreateSDGResponse {
	s.Body = v
	return s
}

func (s *CreateSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
