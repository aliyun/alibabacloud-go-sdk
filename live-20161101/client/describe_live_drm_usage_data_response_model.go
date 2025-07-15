// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDrmUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDrmUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDrmUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDrmUsageDataResponseBody) *DescribeLiveDrmUsageDataResponse
	GetBody() *DescribeLiveDrmUsageDataResponseBody
}

type DescribeLiveDrmUsageDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDrmUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDrmUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDrmUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDrmUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDrmUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDrmUsageDataResponse) GetBody() *DescribeLiveDrmUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDrmUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDrmUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDrmUsageDataResponse) SetStatusCode(v int32) *DescribeLiveDrmUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponse) SetBody(v *DescribeLiveDrmUsageDataResponseBody) *DescribeLiveDrmUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDrmUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
