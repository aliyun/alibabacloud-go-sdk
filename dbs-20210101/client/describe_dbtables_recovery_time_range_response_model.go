// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryTimeRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryTimeRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBTablesRecoveryTimeRangeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBTablesRecoveryTimeRangeResponseBody) *DescribeDBTablesRecoveryTimeRangeResponse
	GetBody() *DescribeDBTablesRecoveryTimeRangeResponseBody
}

type DescribeDBTablesRecoveryTimeRangeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) GetBody() *DescribeDBTablesRecoveryTimeRangeResponseBody {
	return s.Body
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetBody(v *DescribeDBTablesRecoveryTimeRangeResponseBody) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
