// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizUserTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBizUserTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBizUserTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBizUserTypeResponseBody) *DescribeBizUserTypeResponse
	GetBody() *DescribeBizUserTypeResponseBody
}

type DescribeBizUserTypeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizUserTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizUserTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizUserTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizUserTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBizUserTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBizUserTypeResponse) GetBody() *DescribeBizUserTypeResponseBody {
	return s.Body
}

func (s *DescribeBizUserTypeResponse) SetHeaders(v map[string]*string) *DescribeBizUserTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizUserTypeResponse) SetStatusCode(v int32) *DescribeBizUserTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizUserTypeResponse) SetBody(v *DescribeBizUserTypeResponseBody) *DescribeBizUserTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeBizUserTypeResponse) Validate() error {
	return dara.Validate(s)
}
