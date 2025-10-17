// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFaqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFaqResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFaqResponseBody) *UpdateFaqResponse
	GetBody() *UpdateFaqResponseBody
}

type UpdateFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaqResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaqResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFaqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFaqResponse) GetBody() *UpdateFaqResponseBody {
	return s.Body
}

func (s *UpdateFaqResponse) SetHeaders(v map[string]*string) *UpdateFaqResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaqResponse) SetStatusCode(v int32) *UpdateFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaqResponse) SetBody(v *UpdateFaqResponseBody) *UpdateFaqResponse {
	s.Body = v
	return s
}

func (s *UpdateFaqResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
