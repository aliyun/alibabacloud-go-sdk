// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConnDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConnDataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConnDataResponseBody) *UpdateConnDataResponse
	GetBody() *UpdateConnDataResponseBody
}

type UpdateConnDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConnDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConnDataResponse) GetBody() *UpdateConnDataResponseBody {
	return s.Body
}

func (s *UpdateConnDataResponse) SetHeaders(v map[string]*string) *UpdateConnDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnDataResponse) SetStatusCode(v int32) *UpdateConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnDataResponse) SetBody(v *UpdateConnDataResponseBody) *UpdateConnDataResponse {
	s.Body = v
	return s
}

func (s *UpdateConnDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
