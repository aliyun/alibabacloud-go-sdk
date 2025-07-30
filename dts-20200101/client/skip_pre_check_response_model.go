// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *SkipPreCheckResponseBody) *SkipPreCheckResponse
	GetBody() *SkipPreCheckResponseBody
}

type SkipPreCheckResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipPreCheckResponse) GoString() string {
	return s.String()
}

func (s *SkipPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipPreCheckResponse) GetBody() *SkipPreCheckResponseBody {
	return s.Body
}

func (s *SkipPreCheckResponse) SetHeaders(v map[string]*string) *SkipPreCheckResponse {
	s.Headers = v
	return s
}

func (s *SkipPreCheckResponse) SetStatusCode(v int32) *SkipPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipPreCheckResponse) SetBody(v *SkipPreCheckResponseBody) *SkipPreCheckResponse {
	s.Body = v
	return s
}

func (s *SkipPreCheckResponse) Validate() error {
	return dara.Validate(s)
}
