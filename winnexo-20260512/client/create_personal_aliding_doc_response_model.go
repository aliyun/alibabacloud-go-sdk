// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalAlidingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalAlidingDocResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalAlidingDocResponseBody) *CreatePersonalAlidingDocResponse
	GetBody() *CreatePersonalAlidingDocResponseBody
}

type CreatePersonalAlidingDocResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalAlidingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalAlidingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingDocResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalAlidingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalAlidingDocResponse) GetBody() *CreatePersonalAlidingDocResponseBody {
	return s.Body
}

func (s *CreatePersonalAlidingDocResponse) SetHeaders(v map[string]*string) *CreatePersonalAlidingDocResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalAlidingDocResponse) SetStatusCode(v int32) *CreatePersonalAlidingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalAlidingDocResponse) SetBody(v *CreatePersonalAlidingDocResponseBody) *CreatePersonalAlidingDocResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalAlidingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
