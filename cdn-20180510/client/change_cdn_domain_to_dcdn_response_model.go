// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCdnDomainToDcdnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCdnDomainToDcdnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCdnDomainToDcdnResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCdnDomainToDcdnResponseBody) *ChangeCdnDomainToDcdnResponse
	GetBody() *ChangeCdnDomainToDcdnResponseBody
}

type ChangeCdnDomainToDcdnResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCdnDomainToDcdnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCdnDomainToDcdnResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCdnDomainToDcdnResponse) GoString() string {
	return s.String()
}

func (s *ChangeCdnDomainToDcdnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCdnDomainToDcdnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCdnDomainToDcdnResponse) GetBody() *ChangeCdnDomainToDcdnResponseBody {
	return s.Body
}

func (s *ChangeCdnDomainToDcdnResponse) SetHeaders(v map[string]*string) *ChangeCdnDomainToDcdnResponse {
	s.Headers = v
	return s
}

func (s *ChangeCdnDomainToDcdnResponse) SetStatusCode(v int32) *ChangeCdnDomainToDcdnResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCdnDomainToDcdnResponse) SetBody(v *ChangeCdnDomainToDcdnResponseBody) *ChangeCdnDomainToDcdnResponse {
	s.Body = v
	return s
}

func (s *ChangeCdnDomainToDcdnResponse) Validate() error {
	return dara.Validate(s)
}
