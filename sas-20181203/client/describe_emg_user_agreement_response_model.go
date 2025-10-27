// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmgUserAgreementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEmgUserAgreementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEmgUserAgreementResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEmgUserAgreementResponseBody) *DescribeEmgUserAgreementResponse
	GetBody() *DescribeEmgUserAgreementResponseBody
}

type DescribeEmgUserAgreementResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmgUserAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEmgUserAgreementResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgUserAgreementResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmgUserAgreementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEmgUserAgreementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEmgUserAgreementResponse) GetBody() *DescribeEmgUserAgreementResponseBody {
	return s.Body
}

func (s *DescribeEmgUserAgreementResponse) SetHeaders(v map[string]*string) *DescribeEmgUserAgreementResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmgUserAgreementResponse) SetStatusCode(v int32) *DescribeEmgUserAgreementResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmgUserAgreementResponse) SetBody(v *DescribeEmgUserAgreementResponseBody) *DescribeEmgUserAgreementResponse {
	s.Body = v
	return s
}

func (s *DescribeEmgUserAgreementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
