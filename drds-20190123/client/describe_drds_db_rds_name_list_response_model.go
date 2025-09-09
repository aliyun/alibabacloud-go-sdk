// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbRdsNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDbRdsNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDbRdsNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDbRdsNameListResponseBody) *DescribeDrdsDbRdsNameListResponse
	GetBody() *DescribeDrdsDbRdsNameListResponseBody
}

type DescribeDrdsDbRdsNameListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbRdsNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbRdsNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDbRdsNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDbRdsNameListResponse) GetBody() *DescribeDrdsDbRdsNameListResponseBody {
	return s.Body
}

func (s *DescribeDrdsDbRdsNameListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbRdsNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetStatusCode(v int32) *DescribeDrdsDbRdsNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) SetBody(v *DescribeDrdsDbRdsNameListResponseBody) *DescribeDrdsDbRdsNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponse) Validate() error {
	return dara.Validate(s)
}
