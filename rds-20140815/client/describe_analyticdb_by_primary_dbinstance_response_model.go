// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalyticdbByPrimaryDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnalyticdbByPrimaryDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnalyticdbByPrimaryDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) *DescribeAnalyticdbByPrimaryDBInstanceResponse
	GetBody() *DescribeAnalyticdbByPrimaryDBInstanceResponseBody
}

type DescribeAnalyticdbByPrimaryDBInstanceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnalyticdbByPrimaryDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnalyticdbByPrimaryDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalyticdbByPrimaryDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) GetBody() *DescribeAnalyticdbByPrimaryDBInstanceResponseBody {
	return s.Body
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) SetHeaders(v map[string]*string) *DescribeAnalyticdbByPrimaryDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) SetStatusCode(v int32) *DescribeAnalyticdbByPrimaryDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) SetBody(v *DescribeAnalyticdbByPrimaryDBInstanceResponseBody) *DescribeAnalyticdbByPrimaryDBInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
