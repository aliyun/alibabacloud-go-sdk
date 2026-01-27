// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBTablesRecoveryStateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBTablesRecoveryStateResponseBody) *DescribeDBTablesRecoveryStateResponse
	GetBody() *DescribeDBTablesRecoveryStateResponseBody
}

type DescribeDBTablesRecoveryStateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryStateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBTablesRecoveryStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBTablesRecoveryStateResponse) GetBody() *DescribeDBTablesRecoveryStateResponseBody {
	return s.Body
}

func (s *DescribeDBTablesRecoveryStateResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponse) SetBody(v *DescribeDBTablesRecoveryStateResponseBody) *DescribeDBTablesRecoveryStateResponse {
	s.Body = v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
