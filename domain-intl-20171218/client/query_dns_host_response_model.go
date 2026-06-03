// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDnsHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDnsHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDnsHostResponse
	GetStatusCode() *int32
	SetBody(v *QueryDnsHostResponseBody) *QueryDnsHostResponse
	GetBody() *QueryDnsHostResponseBody
}

type QueryDnsHostResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDnsHostResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDnsHostResponse) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDnsHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDnsHostResponse) GetBody() *QueryDnsHostResponseBody {
	return s.Body
}

func (s *QueryDnsHostResponse) SetHeaders(v map[string]*string) *QueryDnsHostResponse {
	s.Headers = v
	return s
}

func (s *QueryDnsHostResponse) SetStatusCode(v int32) *QueryDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDnsHostResponse) SetBody(v *QueryDnsHostResponseBody) *QueryDnsHostResponse {
	s.Body = v
	return s
}

func (s *QueryDnsHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
