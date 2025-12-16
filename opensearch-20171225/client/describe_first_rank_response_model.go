// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirstRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirstRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirstRankResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirstRankResponseBody) *DescribeFirstRankResponse
	GetBody() *DescribeFirstRankResponseBody
}

type DescribeFirstRankResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirstRankResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirstRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirstRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirstRankResponse) GetBody() *DescribeFirstRankResponseBody {
	return s.Body
}

func (s *DescribeFirstRankResponse) SetHeaders(v map[string]*string) *DescribeFirstRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirstRankResponse) SetStatusCode(v int32) *DescribeFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirstRankResponse) SetBody(v *DescribeFirstRankResponseBody) *DescribeFirstRankResponse {
	s.Body = v
	return s
}

func (s *DescribeFirstRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
