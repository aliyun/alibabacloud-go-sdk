// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetDcdnDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetDcdnDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetDcdnDomainConfigsResponseBody) *BatchSetDcdnDomainConfigsResponse
	GetBody() *BatchSetDcdnDomainConfigsResponseBody
}

type BatchSetDcdnDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetDcdnDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetDcdnDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetDcdnDomainConfigsResponse) GetBody() *BatchSetDcdnDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponse) SetStatusCode(v int32) *BatchSetDcdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponse) SetBody(v *BatchSetDcdnDomainConfigsResponseBody) *BatchSetDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
