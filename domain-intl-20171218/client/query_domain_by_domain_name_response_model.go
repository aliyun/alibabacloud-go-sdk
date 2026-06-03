// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByDomainNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainByDomainNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainByDomainNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainByDomainNameResponseBody) *QueryDomainByDomainNameResponse
	GetBody() *QueryDomainByDomainNameResponseBody
}

type QueryDomainByDomainNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainByDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainByDomainNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainByDomainNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainByDomainNameResponse) GetBody() *QueryDomainByDomainNameResponseBody {
	return s.Body
}

func (s *QueryDomainByDomainNameResponse) SetHeaders(v map[string]*string) *QueryDomainByDomainNameResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByDomainNameResponse) SetStatusCode(v int32) *QueryDomainByDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByDomainNameResponse) SetBody(v *QueryDomainByDomainNameResponseBody) *QueryDomainByDomainNameResponse {
	s.Body = v
	return s
}

func (s *QueryDomainByDomainNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
