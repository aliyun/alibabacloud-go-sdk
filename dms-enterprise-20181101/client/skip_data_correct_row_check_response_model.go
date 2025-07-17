// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipDataCorrectRowCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipDataCorrectRowCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipDataCorrectRowCheckResponse
	GetStatusCode() *int32
	SetBody(v *SkipDataCorrectRowCheckResponseBody) *SkipDataCorrectRowCheckResponse
	GetBody() *SkipDataCorrectRowCheckResponseBody
}

type SkipDataCorrectRowCheckResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipDataCorrectRowCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipDataCorrectRowCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipDataCorrectRowCheckResponse) GoString() string {
	return s.String()
}

func (s *SkipDataCorrectRowCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipDataCorrectRowCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipDataCorrectRowCheckResponse) GetBody() *SkipDataCorrectRowCheckResponseBody {
	return s.Body
}

func (s *SkipDataCorrectRowCheckResponse) SetHeaders(v map[string]*string) *SkipDataCorrectRowCheckResponse {
	s.Headers = v
	return s
}

func (s *SkipDataCorrectRowCheckResponse) SetStatusCode(v int32) *SkipDataCorrectRowCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipDataCorrectRowCheckResponse) SetBody(v *SkipDataCorrectRowCheckResponseBody) *SkipDataCorrectRowCheckResponse {
	s.Body = v
	return s
}

func (s *SkipDataCorrectRowCheckResponse) Validate() error {
	return dara.Validate(s)
}
