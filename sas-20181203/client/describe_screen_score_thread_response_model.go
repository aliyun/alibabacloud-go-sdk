// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScreenScoreThreadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScreenScoreThreadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScreenScoreThreadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScreenScoreThreadResponseBody) *DescribeScreenScoreThreadResponse
	GetBody() *DescribeScreenScoreThreadResponseBody
}

type DescribeScreenScoreThreadResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenScoreThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenScoreThreadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScreenScoreThreadResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScreenScoreThreadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScreenScoreThreadResponse) GetBody() *DescribeScreenScoreThreadResponseBody {
	return s.Body
}

func (s *DescribeScreenScoreThreadResponse) SetHeaders(v map[string]*string) *DescribeScreenScoreThreadResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenScoreThreadResponse) SetStatusCode(v int32) *DescribeScreenScoreThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenScoreThreadResponse) SetBody(v *DescribeScreenScoreThreadResponseBody) *DescribeScreenScoreThreadResponse {
	s.Body = v
	return s
}

func (s *DescribeScreenScoreThreadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
