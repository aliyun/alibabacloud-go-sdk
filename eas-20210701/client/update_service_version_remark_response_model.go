// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceVersionRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceVersionRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceVersionRemarkResponseBody) *UpdateServiceVersionRemarkResponse
	GetBody() *UpdateServiceVersionRemarkResponseBody
}

type UpdateServiceVersionRemarkResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceVersionRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceVersionRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceVersionRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceVersionRemarkResponse) GetBody() *UpdateServiceVersionRemarkResponseBody {
	return s.Body
}

func (s *UpdateServiceVersionRemarkResponse) SetHeaders(v map[string]*string) *UpdateServiceVersionRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceVersionRemarkResponse) SetStatusCode(v int32) *UpdateServiceVersionRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceVersionRemarkResponse) SetBody(v *UpdateServiceVersionRemarkResponseBody) *UpdateServiceVersionRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceVersionRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
