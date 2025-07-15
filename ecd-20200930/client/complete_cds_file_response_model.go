// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *CompleteCdsFileResponseBody) *CompleteCdsFileResponse
	GetBody() *CompleteCdsFileResponseBody
}

type CompleteCdsFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteCdsFileResponse) GoString() string {
	return s.String()
}

func (s *CompleteCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteCdsFileResponse) GetBody() *CompleteCdsFileResponseBody {
	return s.Body
}

func (s *CompleteCdsFileResponse) SetHeaders(v map[string]*string) *CompleteCdsFileResponse {
	s.Headers = v
	return s
}

func (s *CompleteCdsFileResponse) SetStatusCode(v int32) *CompleteCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteCdsFileResponse) SetBody(v *CompleteCdsFileResponseBody) *CompleteCdsFileResponse {
	s.Body = v
	return s
}

func (s *CompleteCdsFileResponse) Validate() error {
	return dara.Validate(s)
}
