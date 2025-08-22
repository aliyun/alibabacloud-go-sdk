// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnIpaDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetDcdnIpaDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetDcdnIpaDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetDcdnIpaDomainConfigsResponseBody) *BatchSetDcdnIpaDomainConfigsResponse
	GetBody() *BatchSetDcdnIpaDomainConfigsResponseBody
}

type BatchSetDcdnIpaDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetDcdnIpaDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetDcdnIpaDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) GetBody() *BatchSetDcdnIpaDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetDcdnIpaDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) SetStatusCode(v int32) *BatchSetDcdnIpaDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) SetBody(v *BatchSetDcdnIpaDomainConfigsResponseBody) *BatchSetDcdnIpaDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
