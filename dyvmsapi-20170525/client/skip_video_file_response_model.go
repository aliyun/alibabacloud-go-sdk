// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *SkipVideoFileResponseBody) *SkipVideoFileResponse
	GetBody() *SkipVideoFileResponseBody
}

type SkipVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipVideoFileResponse) GoString() string {
	return s.String()
}

func (s *SkipVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipVideoFileResponse) GetBody() *SkipVideoFileResponseBody {
	return s.Body
}

func (s *SkipVideoFileResponse) SetHeaders(v map[string]*string) *SkipVideoFileResponse {
	s.Headers = v
	return s
}

func (s *SkipVideoFileResponse) SetStatusCode(v int32) *SkipVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipVideoFileResponse) SetBody(v *SkipVideoFileResponseBody) *SkipVideoFileResponse {
	s.Body = v
	return s
}

func (s *SkipVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
