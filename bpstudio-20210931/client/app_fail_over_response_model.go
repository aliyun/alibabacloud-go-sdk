// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailOverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppFailOverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppFailOverResponse
	GetStatusCode() *int32
	SetBody(v *AppFailOverResponseBody) *AppFailOverResponse
	GetBody() *AppFailOverResponseBody
}

type AppFailOverResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppFailOverResponse) String() string {
	return dara.Prettify(s)
}

func (s AppFailOverResponse) GoString() string {
	return s.String()
}

func (s *AppFailOverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppFailOverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppFailOverResponse) GetBody() *AppFailOverResponseBody {
	return s.Body
}

func (s *AppFailOverResponse) SetHeaders(v map[string]*string) *AppFailOverResponse {
	s.Headers = v
	return s
}

func (s *AppFailOverResponse) SetStatusCode(v int32) *AppFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *AppFailOverResponse) SetBody(v *AppFailOverResponseBody) *AppFailOverResponse {
	s.Body = v
	return s
}

func (s *AppFailOverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
