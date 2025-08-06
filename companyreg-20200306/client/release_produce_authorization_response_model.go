// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseProduceAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseProduceAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseProduceAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseProduceAuthorizationResponseBody) *ReleaseProduceAuthorizationResponse
	GetBody() *ReleaseProduceAuthorizationResponseBody
}

type ReleaseProduceAuthorizationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseProduceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseProduceAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseProduceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseProduceAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseProduceAuthorizationResponse) GetBody() *ReleaseProduceAuthorizationResponseBody {
	return s.Body
}

func (s *ReleaseProduceAuthorizationResponse) SetHeaders(v map[string]*string) *ReleaseProduceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ReleaseProduceAuthorizationResponse) SetStatusCode(v int32) *ReleaseProduceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponse) SetBody(v *ReleaseProduceAuthorizationResponseBody) *ReleaseProduceAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ReleaseProduceAuthorizationResponse) Validate() error {
	return dara.Validate(s)
}
