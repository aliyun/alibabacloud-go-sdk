// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdbMySqlTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdbMySqlTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdbMySqlTableMetaResponseBody) *DescribeAdbMySqlTableMetaResponse
	GetBody() *DescribeAdbMySqlTableMetaResponseBody
}

type DescribeAdbMySqlTableMetaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTableMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdbMySqlTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdbMySqlTableMetaResponse) GetBody() *DescribeAdbMySqlTableMetaResponseBody {
	return s.Body
}

func (s *DescribeAdbMySqlTableMetaResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlTableMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponse) SetStatusCode(v int32) *DescribeAdbMySqlTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponse) SetBody(v *DescribeAdbMySqlTableMetaResponseBody) *DescribeAdbMySqlTableMetaResponse {
	s.Body = v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
