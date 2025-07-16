// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddCdnDomainResponseBody) *BatchAddCdnDomainResponse
	GetBody() *BatchAddCdnDomainResponseBody
}

type BatchAddCdnDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddCdnDomainResponse) GetBody() *BatchAddCdnDomainResponseBody {
	return s.Body
}

func (s *BatchAddCdnDomainResponse) SetHeaders(v map[string]*string) *BatchAddCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchAddCdnDomainResponse) SetStatusCode(v int32) *BatchAddCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddCdnDomainResponse) SetBody(v *BatchAddCdnDomainResponseBody) *BatchAddCdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchAddCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
