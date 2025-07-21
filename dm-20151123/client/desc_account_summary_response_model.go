// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescAccountSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescAccountSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescAccountSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescAccountSummaryResponseBody) *DescAccountSummaryResponse
	GetBody() *DescAccountSummaryResponseBody
}

type DescAccountSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescAccountSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescAccountSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescAccountSummaryResponse) GetBody() *DescAccountSummaryResponseBody {
	return s.Body
}

func (s *DescAccountSummaryResponse) SetHeaders(v map[string]*string) *DescAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescAccountSummaryResponse) SetStatusCode(v int32) *DescAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescAccountSummaryResponse) SetBody(v *DescAccountSummaryResponseBody) *DescAccountSummaryResponse {
	s.Body = v
	return s
}

func (s *DescAccountSummaryResponse) Validate() error {
	return dara.Validate(s)
}
