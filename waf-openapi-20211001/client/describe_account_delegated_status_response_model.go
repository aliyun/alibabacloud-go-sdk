// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountDelegatedStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountDelegatedStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountDelegatedStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountDelegatedStatusResponseBody) *DescribeAccountDelegatedStatusResponse
	GetBody() *DescribeAccountDelegatedStatusResponseBody
}

type DescribeAccountDelegatedStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountDelegatedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountDelegatedStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountDelegatedStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountDelegatedStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountDelegatedStatusResponse) GetBody() *DescribeAccountDelegatedStatusResponseBody {
	return s.Body
}

func (s *DescribeAccountDelegatedStatusResponse) SetHeaders(v map[string]*string) *DescribeAccountDelegatedStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountDelegatedStatusResponse) SetStatusCode(v int32) *DescribeAccountDelegatedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponse) SetBody(v *DescribeAccountDelegatedStatusResponseBody) *DescribeAccountDelegatedStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountDelegatedStatusResponse) Validate() error {
	return dara.Validate(s)
}
