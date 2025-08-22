// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnWafDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetDcdnWafDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetDcdnWafDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetDcdnWafDomainConfigsResponseBody) *BatchSetDcdnWafDomainConfigsResponse
	GetBody() *BatchSetDcdnWafDomainConfigsResponseBody
}

type BatchSetDcdnWafDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetDcdnWafDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetDcdnWafDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnWafDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnWafDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetDcdnWafDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetDcdnWafDomainConfigsResponse) GetBody() *BatchSetDcdnWafDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetDcdnWafDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetDcdnWafDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsResponse) SetStatusCode(v int32) *BatchSetDcdnWafDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsResponse) SetBody(v *BatchSetDcdnWafDomainConfigsResponseBody) *BatchSetDcdnWafDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
