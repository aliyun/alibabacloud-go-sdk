// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecuritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppSecuritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppSecuritiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppSecuritiesResponseBody) *DescribeAppSecuritiesResponse
	GetBody() *DescribeAppSecuritiesResponseBody
}

type DescribeAppSecuritiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppSecuritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppSecuritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecuritiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppSecuritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppSecuritiesResponse) GetBody() *DescribeAppSecuritiesResponseBody {
	return s.Body
}

func (s *DescribeAppSecuritiesResponse) SetHeaders(v map[string]*string) *DescribeAppSecuritiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSecuritiesResponse) SetStatusCode(v int32) *DescribeAppSecuritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppSecuritiesResponse) SetBody(v *DescribeAppSecuritiesResponseBody) *DescribeAppSecuritiesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppSecuritiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
