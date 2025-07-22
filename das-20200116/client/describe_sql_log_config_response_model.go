// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlLogConfigResponseBody) *DescribeSqlLogConfigResponse
	GetBody() *DescribeSqlLogConfigResponseBody
}

type DescribeSqlLogConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlLogConfigResponse) GetBody() *DescribeSqlLogConfigResponseBody {
	return s.Body
}

func (s *DescribeSqlLogConfigResponse) SetHeaders(v map[string]*string) *DescribeSqlLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlLogConfigResponse) SetStatusCode(v int32) *DescribeSqlLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlLogConfigResponse) SetBody(v *DescribeSqlLogConfigResponseBody) *DescribeSqlLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
