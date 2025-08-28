// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDegradeVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DegradeVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DegradeVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *DegradeVideoFileResponseBody) *DegradeVideoFileResponse
	GetBody() *DegradeVideoFileResponseBody
}

type DegradeVideoFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DegradeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DegradeVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DegradeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DegradeVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DegradeVideoFileResponse) GetBody() *DegradeVideoFileResponseBody {
	return s.Body
}

func (s *DegradeVideoFileResponse) SetHeaders(v map[string]*string) *DegradeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *DegradeVideoFileResponse) SetStatusCode(v int32) *DegradeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DegradeVideoFileResponse) SetBody(v *DegradeVideoFileResponseBody) *DegradeVideoFileResponse {
	s.Body = v
	return s
}

func (s *DegradeVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
