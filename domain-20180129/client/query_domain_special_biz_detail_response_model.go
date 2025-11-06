// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainSpecialBizDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainSpecialBizDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainSpecialBizDetailResponseBody) *QueryDomainSpecialBizDetailResponse
	GetBody() *QueryDomainSpecialBizDetailResponseBody
}

type QueryDomainSpecialBizDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainSpecialBizDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainSpecialBizDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainSpecialBizDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainSpecialBizDetailResponse) GetBody() *QueryDomainSpecialBizDetailResponseBody {
	return s.Body
}

func (s *QueryDomainSpecialBizDetailResponse) SetHeaders(v map[string]*string) *QueryDomainSpecialBizDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponse) SetStatusCode(v int32) *QueryDomainSpecialBizDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponse) SetBody(v *QueryDomainSpecialBizDetailResponseBody) *QueryDomainSpecialBizDetailResponse {
	s.Body = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
