// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEngineVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEngineVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEngineVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEngineVersionResponseBody) *DescribeEngineVersionResponse
	GetBody() *DescribeEngineVersionResponseBody
}

type DescribeEngineVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEngineVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEngineVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEngineVersionResponse) GetBody() *DescribeEngineVersionResponseBody {
	return s.Body
}

func (s *DescribeEngineVersionResponse) SetHeaders(v map[string]*string) *DescribeEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeEngineVersionResponse) SetStatusCode(v int32) *DescribeEngineVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEngineVersionResponse) SetBody(v *DescribeEngineVersionResponseBody) *DescribeEngineVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeEngineVersionResponse) Validate() error {
	return dara.Validate(s)
}
