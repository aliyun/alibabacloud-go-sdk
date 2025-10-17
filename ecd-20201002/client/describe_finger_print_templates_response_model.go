// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFingerPrintTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFingerPrintTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFingerPrintTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFingerPrintTemplatesResponseBody) *DescribeFingerPrintTemplatesResponse
	GetBody() *DescribeFingerPrintTemplatesResponseBody
}

type DescribeFingerPrintTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFingerPrintTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFingerPrintTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFingerPrintTemplatesResponse) GetBody() *DescribeFingerPrintTemplatesResponseBody {
	return s.Body
}

func (s *DescribeFingerPrintTemplatesResponse) SetHeaders(v map[string]*string) *DescribeFingerPrintTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetStatusCode(v int32) *DescribeFingerPrintTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetBody(v *DescribeFingerPrintTemplatesResponseBody) *DescribeFingerPrintTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
