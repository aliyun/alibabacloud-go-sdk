// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosEventListResponseBody) *DescribeDdosEventListResponse
	GetBody() *DescribeDdosEventListResponseBody
}

type DescribeDdosEventListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosEventListResponse) GetBody() *DescribeDdosEventListResponseBody {
	return s.Body
}

func (s *DescribeDdosEventListResponse) SetHeaders(v map[string]*string) *DescribeDdosEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventListResponse) SetStatusCode(v int32) *DescribeDdosEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventListResponse) SetBody(v *DescribeDdosEventListResponseBody) *DescribeDdosEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosEventListResponse) Validate() error {
	return dara.Validate(s)
}
