// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableEngineVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableEngineVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableEngineVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableEngineVersionResponseBody) *DescribeAvailableEngineVersionResponse
	GetBody() *DescribeAvailableEngineVersionResponseBody
}

type DescribeAvailableEngineVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableEngineVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEngineVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableEngineVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableEngineVersionResponse) GetBody() *DescribeAvailableEngineVersionResponseBody {
	return s.Body
}

func (s *DescribeAvailableEngineVersionResponse) SetHeaders(v map[string]*string) *DescribeAvailableEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableEngineVersionResponse) SetStatusCode(v int32) *DescribeAvailableEngineVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableEngineVersionResponse) SetBody(v *DescribeAvailableEngineVersionResponseBody) *DescribeAvailableEngineVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableEngineVersionResponse) Validate() error {
	return dara.Validate(s)
}
