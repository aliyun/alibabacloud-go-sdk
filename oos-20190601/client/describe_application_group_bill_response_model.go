// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationGroupBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationGroupBillResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationGroupBillResponseBody) *DescribeApplicationGroupBillResponse
	GetBody() *DescribeApplicationGroupBillResponseBody
}

type DescribeApplicationGroupBillResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationGroupBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationGroupBillResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationGroupBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationGroupBillResponse) GetBody() *DescribeApplicationGroupBillResponseBody {
	return s.Body
}

func (s *DescribeApplicationGroupBillResponse) SetHeaders(v map[string]*string) *DescribeApplicationGroupBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationGroupBillResponse) SetStatusCode(v int32) *DescribeApplicationGroupBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationGroupBillResponse) SetBody(v *DescribeApplicationGroupBillResponseBody) *DescribeApplicationGroupBillResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationGroupBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
