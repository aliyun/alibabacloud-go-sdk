// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDcdnDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDcdnDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDcdnDomainConfigsResponseBody) *BatchDeleteDcdnDomainConfigsResponse
	GetBody() *BatchDeleteDcdnDomainConfigsResponseBody
}

type BatchDeleteDcdnDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDcdnDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDcdnDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDcdnDomainConfigsResponse) GetBody() *BatchDeleteDcdnDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchDeleteDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsResponse) SetStatusCode(v int32) *BatchDeleteDcdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsResponse) SetBody(v *BatchDeleteDcdnDomainConfigsResponseBody) *BatchDeleteDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
