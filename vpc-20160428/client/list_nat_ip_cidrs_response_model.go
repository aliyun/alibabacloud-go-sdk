// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpCidrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNatIpCidrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNatIpCidrsResponse
	GetStatusCode() *int32
	SetBody(v *ListNatIpCidrsResponseBody) *ListNatIpCidrsResponse
	GetBody() *ListNatIpCidrsResponseBody
}

type ListNatIpCidrsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNatIpCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNatIpCidrsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpCidrsResponse) GoString() string {
	return s.String()
}

func (s *ListNatIpCidrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNatIpCidrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNatIpCidrsResponse) GetBody() *ListNatIpCidrsResponseBody {
	return s.Body
}

func (s *ListNatIpCidrsResponse) SetHeaders(v map[string]*string) *ListNatIpCidrsResponse {
	s.Headers = v
	return s
}

func (s *ListNatIpCidrsResponse) SetStatusCode(v int32) *ListNatIpCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNatIpCidrsResponse) SetBody(v *ListNatIpCidrsResponseBody) *ListNatIpCidrsResponse {
	s.Body = v
	return s
}

func (s *ListNatIpCidrsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
