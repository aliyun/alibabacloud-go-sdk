// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizInfoByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainSpecialBizInfoByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainSpecialBizInfoByDomainResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainSpecialBizInfoByDomainResponseBody) *QueryDomainSpecialBizInfoByDomainResponse
	GetBody() *QueryDomainSpecialBizInfoByDomainResponseBody
}

type QueryDomainSpecialBizInfoByDomainResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainSpecialBizInfoByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) GetBody() *QueryDomainSpecialBizInfoByDomainResponseBody {
	return s.Body
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) SetHeaders(v map[string]*string) *QueryDomainSpecialBizInfoByDomainResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) SetStatusCode(v int32) *QueryDomainSpecialBizInfoByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) SetBody(v *QueryDomainSpecialBizInfoByDomainResponseBody) *QueryDomainSpecialBizInfoByDomainResponse {
	s.Body = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
