// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSuffixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainSuffixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainSuffixResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainSuffixResponseBody) *QueryDomainSuffixResponse
	GetBody() *QueryDomainSuffixResponseBody
}

type QueryDomainSuffixResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainSuffixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainSuffixResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSuffixResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainSuffixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainSuffixResponse) GetBody() *QueryDomainSuffixResponseBody {
	return s.Body
}

func (s *QueryDomainSuffixResponse) SetHeaders(v map[string]*string) *QueryDomainSuffixResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainSuffixResponse) SetStatusCode(v int32) *QueryDomainSuffixResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainSuffixResponse) SetBody(v *QueryDomainSuffixResponseBody) *QueryDomainSuffixResponse {
	s.Body = v
	return s
}

func (s *QueryDomainSuffixResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
