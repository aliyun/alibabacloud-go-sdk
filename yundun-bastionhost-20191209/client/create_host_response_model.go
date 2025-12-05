// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHostResponse
	GetStatusCode() *int32
	SetBody(v *CreateHostResponseBody) *CreateHostResponse
	GetBody() *CreateHostResponseBody
}

type CreateHostResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHostResponse) GoString() string {
	return s.String()
}

func (s *CreateHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHostResponse) GetBody() *CreateHostResponseBody {
	return s.Body
}

func (s *CreateHostResponse) SetHeaders(v map[string]*string) *CreateHostResponse {
	s.Headers = v
	return s
}

func (s *CreateHostResponse) SetStatusCode(v int32) *CreateHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostResponse) SetBody(v *CreateHostResponseBody) *CreateHostResponse {
	s.Body = v
	return s
}

func (s *CreateHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
