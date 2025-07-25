// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBatchResultDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBatchResultDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBatchResultDetailResponseBody) *DescribeBatchResultDetailResponse
	GetBody() *DescribeBatchResultDetailResponseBody
}

type DescribeBatchResultDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBatchResultDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBatchResultDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBatchResultDetailResponse) GetBody() *DescribeBatchResultDetailResponseBody {
	return s.Body
}

func (s *DescribeBatchResultDetailResponse) SetHeaders(v map[string]*string) *DescribeBatchResultDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchResultDetailResponse) SetStatusCode(v int32) *DescribeBatchResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchResultDetailResponse) SetBody(v *DescribeBatchResultDetailResponseBody) *DescribeBatchResultDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeBatchResultDetailResponse) Validate() error {
	return dara.Validate(s)
}
