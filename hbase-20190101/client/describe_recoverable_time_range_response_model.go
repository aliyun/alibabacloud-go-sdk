// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableTimeRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecoverableTimeRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecoverableTimeRangeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecoverableTimeRangeResponseBody) *DescribeRecoverableTimeRangeResponse
	GetBody() *DescribeRecoverableTimeRangeResponseBody
}

type DescribeRecoverableTimeRangeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecoverableTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecoverableTimeRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecoverableTimeRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecoverableTimeRangeResponse) GetBody() *DescribeRecoverableTimeRangeResponseBody {
	return s.Body
}

func (s *DescribeRecoverableTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeRecoverableTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) SetStatusCode(v int32) *DescribeRecoverableTimeRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) SetBody(v *DescribeRecoverableTimeRangeResponseBody) *DescribeRecoverableTimeRangeResponse {
	s.Body = v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
