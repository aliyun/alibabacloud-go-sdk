// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamAuthCheckingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamAuthCheckingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamAuthCheckingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamAuthCheckingResponseBody) *DescribeLiveStreamAuthCheckingResponse
	GetBody() *DescribeLiveStreamAuthCheckingResponseBody
}

type DescribeLiveStreamAuthCheckingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamAuthCheckingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamAuthCheckingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamAuthCheckingResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamAuthCheckingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamAuthCheckingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamAuthCheckingResponse) GetBody() *DescribeLiveStreamAuthCheckingResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamAuthCheckingResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamAuthCheckingResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponse) SetStatusCode(v int32) *DescribeLiveStreamAuthCheckingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponse) SetBody(v *DescribeLiveStreamAuthCheckingResponseBody) *DescribeLiveStreamAuthCheckingResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponse) Validate() error {
	return dara.Validate(s)
}
