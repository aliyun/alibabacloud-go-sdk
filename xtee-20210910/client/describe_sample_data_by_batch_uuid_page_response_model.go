// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataByBatchUUidPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDataByBatchUUidPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDataByBatchUUidPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDataByBatchUUidPageResponseBody) *DescribeSampleDataByBatchUUidPageResponse
	GetBody() *DescribeSampleDataByBatchUUidPageResponseBody
}

type DescribeSampleDataByBatchUUidPageResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDataByBatchUUidPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDataByBatchUUidPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataByBatchUUidPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataByBatchUUidPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDataByBatchUUidPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDataByBatchUUidPageResponse) GetBody() *DescribeSampleDataByBatchUUidPageResponseBody {
	return s.Body
}

func (s *DescribeSampleDataByBatchUUidPageResponse) SetHeaders(v map[string]*string) *DescribeSampleDataByBatchUUidPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponse) SetStatusCode(v int32) *DescribeSampleDataByBatchUUidPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponse) SetBody(v *DescribeSampleDataByBatchUUidPageResponseBody) *DescribeSampleDataByBatchUUidPageResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponse) Validate() error {
	return dara.Validate(s)
}
