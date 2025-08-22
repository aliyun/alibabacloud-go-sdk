// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosAllEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosAllEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosAllEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosAllEventListResponseBody) *DescribeDdosAllEventListResponse
	GetBody() *DescribeDdosAllEventListResponseBody
}

type DescribeDdosAllEventListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosAllEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosAllEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosAllEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosAllEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosAllEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosAllEventListResponse) GetBody() *DescribeDdosAllEventListResponseBody {
	return s.Body
}

func (s *DescribeDdosAllEventListResponse) SetHeaders(v map[string]*string) *DescribeDdosAllEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosAllEventListResponse) SetStatusCode(v int32) *DescribeDdosAllEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosAllEventListResponse) SetBody(v *DescribeDdosAllEventListResponseBody) *DescribeDdosAllEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosAllEventListResponse) Validate() error {
	return dara.Validate(s)
}
