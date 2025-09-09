// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmoothExpandPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmoothExpandPreCheckResponseBody) *SubmitSmoothExpandPreCheckResponse
	GetBody() *SubmitSmoothExpandPreCheckResponseBody
}

type SubmitSmoothExpandPreCheckResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmoothExpandPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmoothExpandPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmoothExpandPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmoothExpandPreCheckResponse) GetBody() *SubmitSmoothExpandPreCheckResponseBody {
	return s.Body
}

func (s *SubmitSmoothExpandPreCheckResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) SetBody(v *SubmitSmoothExpandPreCheckResponseBody) *SubmitSmoothExpandPreCheckResponse {
	s.Body = v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponse) Validate() error {
	return dara.Validate(s)
}
