// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReadDBInstanceDelayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReadDBInstanceDelayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReadDBInstanceDelayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReadDBInstanceDelayResponseBody) *DescribeReadDBInstanceDelayResponse
	GetBody() *DescribeReadDBInstanceDelayResponseBody
}

type DescribeReadDBInstanceDelayResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReadDBInstanceDelayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReadDBInstanceDelayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayResponse) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReadDBInstanceDelayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReadDBInstanceDelayResponse) GetBody() *DescribeReadDBInstanceDelayResponseBody {
	return s.Body
}

func (s *DescribeReadDBInstanceDelayResponse) SetHeaders(v map[string]*string) *DescribeReadDBInstanceDelayResponse {
	s.Headers = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponse) SetStatusCode(v int32) *DescribeReadDBInstanceDelayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReadDBInstanceDelayResponse) SetBody(v *DescribeReadDBInstanceDelayResponseBody) *DescribeReadDBInstanceDelayResponse {
	s.Body = v
	return s
}

func (s *DescribeReadDBInstanceDelayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
