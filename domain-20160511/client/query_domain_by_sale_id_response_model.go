// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainBySaleIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainBySaleIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainBySaleIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainBySaleIdResponseBody) *QueryDomainBySaleIdResponse
	GetBody() *QueryDomainBySaleIdResponseBody
}

type QueryDomainBySaleIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainBySaleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainBySaleIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainBySaleIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainBySaleIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainBySaleIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainBySaleIdResponse) GetBody() *QueryDomainBySaleIdResponseBody {
	return s.Body
}

func (s *QueryDomainBySaleIdResponse) SetHeaders(v map[string]*string) *QueryDomainBySaleIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainBySaleIdResponse) SetStatusCode(v int32) *QueryDomainBySaleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainBySaleIdResponse) SetBody(v *QueryDomainBySaleIdResponseBody) *QueryDomainBySaleIdResponse {
	s.Body = v
	return s
}

func (s *QueryDomainBySaleIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
