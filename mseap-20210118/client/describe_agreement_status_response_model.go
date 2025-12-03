// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgreementStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgreementStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgreementStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgreementStatusResponseBody) *DescribeAgreementStatusResponse
	GetBody() *DescribeAgreementStatusResponseBody
}

type DescribeAgreementStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgreementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgreementStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgreementStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgreementStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgreementStatusResponse) GetBody() *DescribeAgreementStatusResponseBody {
	return s.Body
}

func (s *DescribeAgreementStatusResponse) SetHeaders(v map[string]*string) *DescribeAgreementStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgreementStatusResponse) SetStatusCode(v int32) *DescribeAgreementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgreementStatusResponse) SetBody(v *DescribeAgreementStatusResponseBody) *DescribeAgreementStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAgreementStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
