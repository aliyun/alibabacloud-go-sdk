// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDisposableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDisposableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDisposableResponse
	GetStatusCode() *int32
	SetBody(v *CheckDisposableResponseBody) *CheckDisposableResponse
	GetBody() *CheckDisposableResponseBody
}

type CheckDisposableResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDisposableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDisposableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDisposableResponse) GoString() string {
	return s.String()
}

func (s *CheckDisposableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDisposableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDisposableResponse) GetBody() *CheckDisposableResponseBody {
	return s.Body
}

func (s *CheckDisposableResponse) SetHeaders(v map[string]*string) *CheckDisposableResponse {
	s.Headers = v
	return s
}

func (s *CheckDisposableResponse) SetStatusCode(v int32) *CheckDisposableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDisposableResponse) SetBody(v *CheckDisposableResponseBody) *CheckDisposableResponse {
	s.Body = v
	return s
}

func (s *CheckDisposableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
