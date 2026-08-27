// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalDingtalkMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalDingtalkMinutesResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalDingtalkMinutesResponseBody) *CreatePersonalDingtalkMinutesResponse
	GetBody() *CreatePersonalDingtalkMinutesResponseBody
}

type CreatePersonalDingtalkMinutesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalDingtalkMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalDingtalkMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMinutesResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalDingtalkMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalDingtalkMinutesResponse) GetBody() *CreatePersonalDingtalkMinutesResponseBody {
	return s.Body
}

func (s *CreatePersonalDingtalkMinutesResponse) SetHeaders(v map[string]*string) *CreatePersonalDingtalkMinutesResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponse) SetStatusCode(v int32) *CreatePersonalDingtalkMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponse) SetBody(v *CreatePersonalDingtalkMinutesResponseBody) *CreatePersonalDingtalkMinutesResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
