// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceSourceResponseBody) *UpdateServiceSourceResponse
	GetBody() *UpdateServiceSourceResponseBody
}

type UpdateServiceSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceSourceResponse) GetBody() *UpdateServiceSourceResponseBody {
	return s.Body
}

func (s *UpdateServiceSourceResponse) SetHeaders(v map[string]*string) *UpdateServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceSourceResponse) SetStatusCode(v int32) *UpdateServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceSourceResponse) SetBody(v *UpdateServiceSourceResponseBody) *UpdateServiceSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
