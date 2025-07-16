// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinaryLogListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBinaryLogListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBinaryLogListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBinaryLogListResponseBody) *DescribeBinaryLogListResponse
	GetBody() *DescribeBinaryLogListResponseBody
}

type DescribeBinaryLogListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBinaryLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBinaryLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinaryLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBinaryLogListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBinaryLogListResponse) GetBody() *DescribeBinaryLogListResponseBody {
	return s.Body
}

func (s *DescribeBinaryLogListResponse) SetHeaders(v map[string]*string) *DescribeBinaryLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBinaryLogListResponse) SetStatusCode(v int32) *DescribeBinaryLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBinaryLogListResponse) SetBody(v *DescribeBinaryLogListResponseBody) *DescribeBinaryLogListResponse {
	s.Body = v
	return s
}

func (s *DescribeBinaryLogListResponse) Validate() error {
	return dara.Validate(s)
}
