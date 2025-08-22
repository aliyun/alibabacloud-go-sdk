// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTaskByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnRefreshTaskByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnRefreshTaskByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnRefreshTaskByIdResponseBody) *DescribeDcdnRefreshTaskByIdResponse
	GetBody() *DescribeDcdnRefreshTaskByIdResponseBody
}

type DescribeDcdnRefreshTaskByIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnRefreshTaskByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnRefreshTaskByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnRefreshTaskByIdResponse) GetBody() *DescribeDcdnRefreshTaskByIdResponseBody {
	return s.Body
}

func (s *DescribeDcdnRefreshTaskByIdResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshTaskByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponse) SetStatusCode(v int32) *DescribeDcdnRefreshTaskByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponse) SetBody(v *DescribeDcdnRefreshTaskByIdResponseBody) *DescribeDcdnRefreshTaskByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponse) Validate() error {
	return dara.Validate(s)
}
