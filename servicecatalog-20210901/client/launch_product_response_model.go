// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LaunchProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LaunchProductResponse
	GetStatusCode() *int32
	SetBody(v *LaunchProductResponseBody) *LaunchProductResponse
	GetBody() *LaunchProductResponseBody
}

type LaunchProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchProductResponse) String() string {
	return dara.Prettify(s)
}

func (s LaunchProductResponse) GoString() string {
	return s.String()
}

func (s *LaunchProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LaunchProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LaunchProductResponse) GetBody() *LaunchProductResponseBody {
	return s.Body
}

func (s *LaunchProductResponse) SetHeaders(v map[string]*string) *LaunchProductResponse {
	s.Headers = v
	return s
}

func (s *LaunchProductResponse) SetStatusCode(v int32) *LaunchProductResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchProductResponse) SetBody(v *LaunchProductResponseBody) *LaunchProductResponse {
	s.Body = v
	return s
}

func (s *LaunchProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
