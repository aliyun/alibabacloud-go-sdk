// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSprintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSprintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSprintResponse
	GetStatusCode() *int32
	SetBody(v *CreateSprintResponseBody) *CreateSprintResponse
	GetBody() *CreateSprintResponseBody
}

type CreateSprintResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSprintResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSprintResponse) GoString() string {
	return s.String()
}

func (s *CreateSprintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSprintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSprintResponse) GetBody() *CreateSprintResponseBody {
	return s.Body
}

func (s *CreateSprintResponse) SetHeaders(v map[string]*string) *CreateSprintResponse {
	s.Headers = v
	return s
}

func (s *CreateSprintResponse) SetStatusCode(v int32) *CreateSprintResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSprintResponse) SetBody(v *CreateSprintResponseBody) *CreateSprintResponse {
	s.Body = v
	return s
}

func (s *CreateSprintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
