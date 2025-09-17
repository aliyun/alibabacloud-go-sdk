// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePushMeteringDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePushMeteringDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePushMeteringDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribePushMeteringDataResponseBody) *DescribePushMeteringDataResponse
	GetBody() *DescribePushMeteringDataResponseBody
}

type DescribePushMeteringDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePushMeteringDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *DescribePushMeteringDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePushMeteringDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePushMeteringDataResponse) GetBody() *DescribePushMeteringDataResponseBody {
	return s.Body
}

func (s *DescribePushMeteringDataResponse) SetHeaders(v map[string]*string) *DescribePushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *DescribePushMeteringDataResponse) SetStatusCode(v int32) *DescribePushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePushMeteringDataResponse) SetBody(v *DescribePushMeteringDataResponseBody) *DescribePushMeteringDataResponse {
	s.Body = v
	return s
}

func (s *DescribePushMeteringDataResponse) Validate() error {
	return dara.Validate(s)
}
