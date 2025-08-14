// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableConfigListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustVariableConfigListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustVariableConfigListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustVariableConfigListResponseBody) *DescribeCustVariableConfigListResponse
	GetBody() *DescribeCustVariableConfigListResponseBody
}

type DescribeCustVariableConfigListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustVariableConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustVariableConfigListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableConfigListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustVariableConfigListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustVariableConfigListResponse) GetBody() *DescribeCustVariableConfigListResponseBody {
	return s.Body
}

func (s *DescribeCustVariableConfigListResponse) SetHeaders(v map[string]*string) *DescribeCustVariableConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustVariableConfigListResponse) SetStatusCode(v int32) *DescribeCustVariableConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustVariableConfigListResponse) SetBody(v *DescribeCustVariableConfigListResponseBody) *DescribeCustVariableConfigListResponse {
	s.Body = v
	return s
}

func (s *DescribeCustVariableConfigListResponse) Validate() error {
	return dara.Validate(s)
}
