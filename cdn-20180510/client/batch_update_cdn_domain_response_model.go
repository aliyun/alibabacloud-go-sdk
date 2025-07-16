// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateCdnDomainResponseBody) *BatchUpdateCdnDomainResponse
	GetBody() *BatchUpdateCdnDomainResponseBody
}

type BatchUpdateCdnDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateCdnDomainResponse) GetBody() *BatchUpdateCdnDomainResponseBody {
	return s.Body
}

func (s *BatchUpdateCdnDomainResponse) SetHeaders(v map[string]*string) *BatchUpdateCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateCdnDomainResponse) SetStatusCode(v int32) *BatchUpdateCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateCdnDomainResponse) SetBody(v *BatchUpdateCdnDomainResponseBody) *BatchUpdateCdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
