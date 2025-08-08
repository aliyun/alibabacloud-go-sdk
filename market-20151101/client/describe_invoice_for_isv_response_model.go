// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvoiceForIsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvoiceForIsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvoiceForIsvResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvoiceForIsvResponseBody) *DescribeInvoiceForIsvResponse
	GetBody() *DescribeInvoiceForIsvResponseBody
}

type DescribeInvoiceForIsvResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvoiceForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvoiceForIsvResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvoiceForIsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvoiceForIsvResponse) GetBody() *DescribeInvoiceForIsvResponseBody {
	return s.Body
}

func (s *DescribeInvoiceForIsvResponse) SetHeaders(v map[string]*string) *DescribeInvoiceForIsvResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvoiceForIsvResponse) SetStatusCode(v int32) *DescribeInvoiceForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvoiceForIsvResponse) SetBody(v *DescribeInvoiceForIsvResponseBody) *DescribeInvoiceForIsvResponse {
	s.Body = v
	return s
}

func (s *DescribeInvoiceForIsvResponse) Validate() error {
	return dara.Validate(s)
}
