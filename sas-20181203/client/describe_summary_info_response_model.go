// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSummaryInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSummaryInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSummaryInfoResponseBody) *DescribeSummaryInfoResponse
	GetBody() *DescribeSummaryInfoResponseBody
}

type DescribeSummaryInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSummaryInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSummaryInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSummaryInfoResponse) GetBody() *DescribeSummaryInfoResponseBody {
	return s.Body
}

func (s *DescribeSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSummaryInfoResponse) SetStatusCode(v int32) *DescribeSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSummaryInfoResponse) SetBody(v *DescribeSummaryInfoResponseBody) *DescribeSummaryInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSummaryInfoResponse) Validate() error {
	return dara.Validate(s)
}
