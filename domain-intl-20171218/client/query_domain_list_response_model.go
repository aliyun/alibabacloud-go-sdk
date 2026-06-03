// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainListResponseBody) *QueryDomainListResponse
	GetBody() *QueryDomainListResponseBody
}

type QueryDomainListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainListResponse) GetBody() *QueryDomainListResponseBody {
	return s.Body
}

func (s *QueryDomainListResponse) SetHeaders(v map[string]*string) *QueryDomainListResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainListResponse) SetStatusCode(v int32) *QueryDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainListResponse) SetBody(v *QueryDomainListResponseBody) *QueryDomainListResponse {
	s.Body = v
	return s
}

func (s *QueryDomainListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
