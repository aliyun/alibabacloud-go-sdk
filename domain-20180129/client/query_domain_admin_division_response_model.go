// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainAdminDivisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainAdminDivisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainAdminDivisionResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainAdminDivisionResponseBody) *QueryDomainAdminDivisionResponse
	GetBody() *QueryDomainAdminDivisionResponseBody
}

type QueryDomainAdminDivisionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainAdminDivisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainAdminDivisionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainAdminDivisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainAdminDivisionResponse) GetBody() *QueryDomainAdminDivisionResponseBody {
	return s.Body
}

func (s *QueryDomainAdminDivisionResponse) SetHeaders(v map[string]*string) *QueryDomainAdminDivisionResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainAdminDivisionResponse) SetStatusCode(v int32) *QueryDomainAdminDivisionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainAdminDivisionResponse) SetBody(v *QueryDomainAdminDivisionResponseBody) *QueryDomainAdminDivisionResponse {
	s.Body = v
	return s
}

func (s *QueryDomainAdminDivisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
