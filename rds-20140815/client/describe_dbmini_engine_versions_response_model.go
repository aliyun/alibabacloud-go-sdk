// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBMiniEngineVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBMiniEngineVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBMiniEngineVersionsResponseBody) *DescribeDBMiniEngineVersionsResponse
	GetBody() *DescribeDBMiniEngineVersionsResponseBody
}

type DescribeDBMiniEngineVersionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBMiniEngineVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBMiniEngineVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBMiniEngineVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBMiniEngineVersionsResponse) GetBody() *DescribeDBMiniEngineVersionsResponseBody {
	return s.Body
}

func (s *DescribeDBMiniEngineVersionsResponse) SetHeaders(v map[string]*string) *DescribeDBMiniEngineVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponse) SetStatusCode(v int32) *DescribeDBMiniEngineVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponse) SetBody(v *DescribeDBMiniEngineVersionsResponseBody) *DescribeDBMiniEngineVersionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
