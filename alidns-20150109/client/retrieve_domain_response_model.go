// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveDomainResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveDomainResponseBody) *RetrieveDomainResponse
	GetBody() *RetrieveDomainResponseBody
}

type RetrieveDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveDomainResponse) GoString() string {
	return s.String()
}

func (s *RetrieveDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveDomainResponse) GetBody() *RetrieveDomainResponseBody {
	return s.Body
}

func (s *RetrieveDomainResponse) SetHeaders(v map[string]*string) *RetrieveDomainResponse {
	s.Headers = v
	return s
}

func (s *RetrieveDomainResponse) SetStatusCode(v int32) *RetrieveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveDomainResponse) SetBody(v *RetrieveDomainResponseBody) *RetrieveDomainResponse {
	s.Body = v
	return s
}

func (s *RetrieveDomainResponse) Validate() error {
	return dara.Validate(s)
}
