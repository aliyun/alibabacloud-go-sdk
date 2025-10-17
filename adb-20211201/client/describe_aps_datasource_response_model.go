// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsDatasourceResponseBody) *DescribeApsDatasourceResponse
	GetBody() *DescribeApsDatasourceResponseBody
}

type DescribeApsDatasourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsDatasourceResponse) GetBody() *DescribeApsDatasourceResponseBody {
	return s.Body
}

func (s *DescribeApsDatasourceResponse) SetHeaders(v map[string]*string) *DescribeApsDatasourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsDatasourceResponse) SetStatusCode(v int32) *DescribeApsDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsDatasourceResponse) SetBody(v *DescribeApsDatasourceResponseBody) *DescribeApsDatasourceResponse {
	s.Body = v
	return s
}

func (s *DescribeApsDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
