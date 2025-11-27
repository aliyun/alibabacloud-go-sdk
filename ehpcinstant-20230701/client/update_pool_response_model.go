// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePoolResponseBody) *UpdatePoolResponse
	GetBody() *UpdatePoolResponseBody
}

type UpdatePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePoolResponse) GoString() string {
	return s.String()
}

func (s *UpdatePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePoolResponse) GetBody() *UpdatePoolResponseBody {
	return s.Body
}

func (s *UpdatePoolResponse) SetHeaders(v map[string]*string) *UpdatePoolResponse {
	s.Headers = v
	return s
}

func (s *UpdatePoolResponse) SetStatusCode(v int32) *UpdatePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePoolResponse) SetBody(v *UpdatePoolResponseBody) *UpdatePoolResponse {
	s.Body = v
	return s
}

func (s *UpdatePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
