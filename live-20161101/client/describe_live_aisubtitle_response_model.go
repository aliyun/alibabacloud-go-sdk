// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAISubtitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveAISubtitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveAISubtitleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveAISubtitleResponseBody) *DescribeLiveAISubtitleResponse
	GetBody() *DescribeLiveAISubtitleResponseBody
}

type DescribeLiveAISubtitleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveAISubtitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveAISubtitleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAISubtitleResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAISubtitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveAISubtitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveAISubtitleResponse) GetBody() *DescribeLiveAISubtitleResponseBody {
	return s.Body
}

func (s *DescribeLiveAISubtitleResponse) SetHeaders(v map[string]*string) *DescribeLiveAISubtitleResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveAISubtitleResponse) SetStatusCode(v int32) *DescribeLiveAISubtitleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveAISubtitleResponse) SetBody(v *DescribeLiveAISubtitleResponseBody) *DescribeLiveAISubtitleResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveAISubtitleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
