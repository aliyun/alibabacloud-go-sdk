// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSasInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseSasInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseSasInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseSasInstanceResponseBody) *ReleaseSasInstanceResponse
	GetBody() *ReleaseSasInstanceResponseBody
}

type ReleaseSasInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSasInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSasInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSasInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSasInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseSasInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseSasInstanceResponse) GetBody() *ReleaseSasInstanceResponseBody {
	return s.Body
}

func (s *ReleaseSasInstanceResponse) SetHeaders(v map[string]*string) *ReleaseSasInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSasInstanceResponse) SetStatusCode(v int32) *ReleaseSasInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSasInstanceResponse) SetBody(v *ReleaseSasInstanceResponseBody) *ReleaseSasInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseSasInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
