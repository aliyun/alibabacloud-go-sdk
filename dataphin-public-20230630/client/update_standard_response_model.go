// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardResponseBody) *UpdateStandardResponse
	GetBody() *UpdateStandardResponseBody
}

type UpdateStandardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardResponse) GetBody() *UpdateStandardResponseBody {
	return s.Body
}

func (s *UpdateStandardResponse) SetHeaders(v map[string]*string) *UpdateStandardResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardResponse) SetStatusCode(v int32) *UpdateStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardResponse) SetBody(v *UpdateStandardResponseBody) *UpdateStandardResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
