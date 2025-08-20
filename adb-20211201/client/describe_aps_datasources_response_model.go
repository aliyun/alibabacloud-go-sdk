// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsDatasourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsDatasourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsDatasourcesResponseBody) *DescribeApsDatasourcesResponse
	GetBody() *DescribeApsDatasourcesResponseBody
}

type DescribeApsDatasourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsDatasourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsDatasourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsDatasourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsDatasourcesResponse) GetBody() *DescribeApsDatasourcesResponseBody {
	return s.Body
}

func (s *DescribeApsDatasourcesResponse) SetHeaders(v map[string]*string) *DescribeApsDatasourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsDatasourcesResponse) SetStatusCode(v int32) *DescribeApsDatasourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsDatasourcesResponse) SetBody(v *DescribeApsDatasourcesResponseBody) *DescribeApsDatasourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeApsDatasourcesResponse) Validate() error {
	return dara.Validate(s)
}
