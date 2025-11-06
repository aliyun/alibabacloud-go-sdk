// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainByInstanceIdResponseBody) *QueryDomainByInstanceIdResponse
	GetBody() *QueryDomainByInstanceIdResponseBody
}

type QueryDomainByInstanceIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainByInstanceIdResponse) GetBody() *QueryDomainByInstanceIdResponseBody {
	return s.Body
}

func (s *QueryDomainByInstanceIdResponse) SetHeaders(v map[string]*string) *QueryDomainByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByInstanceIdResponse) SetStatusCode(v int32) *QueryDomainByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByInstanceIdResponse) SetBody(v *QueryDomainByInstanceIdResponseBody) *QueryDomainByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *QueryDomainByInstanceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
