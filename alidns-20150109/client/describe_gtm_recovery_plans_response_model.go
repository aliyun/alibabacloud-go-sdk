// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmRecoveryPlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmRecoveryPlansResponseBody) *DescribeGtmRecoveryPlansResponse
	GetBody() *DescribeGtmRecoveryPlansResponseBody
}

type DescribeGtmRecoveryPlansResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmRecoveryPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmRecoveryPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmRecoveryPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmRecoveryPlansResponse) GetBody() *DescribeGtmRecoveryPlansResponseBody {
	return s.Body
}

func (s *DescribeGtmRecoveryPlansResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponse) SetStatusCode(v int32) *DescribeGtmRecoveryPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponse) SetBody(v *DescribeGtmRecoveryPlansResponseBody) *DescribeGtmRecoveryPlansResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponse) Validate() error {
	return dara.Validate(s)
}
