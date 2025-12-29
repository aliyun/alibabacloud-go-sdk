// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceQAResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceQAResponseBody) *UpdateServiceQAResponse
	GetBody() *UpdateServiceQAResponseBody
}

type UpdateServiceQAResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceQAResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceQAResponse) GetBody() *UpdateServiceQAResponseBody {
	return s.Body
}

func (s *UpdateServiceQAResponse) SetHeaders(v map[string]*string) *UpdateServiceQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceQAResponse) SetStatusCode(v int32) *UpdateServiceQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceQAResponse) SetBody(v *UpdateServiceQAResponseBody) *UpdateServiceQAResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
