// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustVariableDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustVariableDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustVariableDetailResponseBody) *DescribeCustVariableDetailResponse
	GetBody() *DescribeCustVariableDetailResponseBody
}

type DescribeCustVariableDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustVariableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustVariableDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustVariableDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustVariableDetailResponse) GetBody() *DescribeCustVariableDetailResponseBody {
	return s.Body
}

func (s *DescribeCustVariableDetailResponse) SetHeaders(v map[string]*string) *DescribeCustVariableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustVariableDetailResponse) SetStatusCode(v int32) *DescribeCustVariableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustVariableDetailResponse) SetBody(v *DescribeCustVariableDetailResponseBody) *DescribeCustVariableDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCustVariableDetailResponse) Validate() error {
	return dara.Validate(s)
}
