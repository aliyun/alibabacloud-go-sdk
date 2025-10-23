// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainByParamResponseBody) *QueryDomainByParamResponse
	GetBody() *QueryDomainByParamResponseBody
}

type QueryDomainByParamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainByParamResponse) GetBody() *QueryDomainByParamResponseBody {
	return s.Body
}

func (s *QueryDomainByParamResponse) SetHeaders(v map[string]*string) *QueryDomainByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByParamResponse) SetStatusCode(v int32) *QueryDomainByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByParamResponse) SetBody(v *QueryDomainByParamResponseBody) *QueryDomainByParamResponse {
	s.Body = v
	return s
}

func (s *QueryDomainByParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
