// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseSecurityIPListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseSecurityIPListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseSecurityIPListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseSecurityIPListResponseBody) *DescribeLangfuseSecurityIPListResponse
	GetBody() *DescribeLangfuseSecurityIPListResponseBody
}

type DescribeLangfuseSecurityIPListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseSecurityIPListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseSecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseSecurityIPListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseSecurityIPListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseSecurityIPListResponse) GetBody() *DescribeLangfuseSecurityIPListResponseBody {
	return s.Body
}

func (s *DescribeLangfuseSecurityIPListResponse) SetHeaders(v map[string]*string) *DescribeLangfuseSecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponse) SetStatusCode(v int32) *DescribeLangfuseSecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponse) SetBody(v *DescribeLangfuseSecurityIPListResponseBody) *DescribeLangfuseSecurityIPListResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
