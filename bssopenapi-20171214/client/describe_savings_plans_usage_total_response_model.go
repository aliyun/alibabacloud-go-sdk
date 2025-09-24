// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlansUsageTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlansUsageTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlansUsageTotalResponseBody) *DescribeSavingsPlansUsageTotalResponse
	GetBody() *DescribeSavingsPlansUsageTotalResponseBody
}

type DescribeSavingsPlansUsageTotalResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlansUsageTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlansUsageTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlansUsageTotalResponse) GetBody() *DescribeSavingsPlansUsageTotalResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlansUsageTotalResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlansUsageTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponse) SetStatusCode(v int32) *DescribeSavingsPlansUsageTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponse) SetBody(v *DescribeSavingsPlansUsageTotalResponseBody) *DescribeSavingsPlansUsageTotalResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponse) Validate() error {
	return dara.Validate(s)
}
