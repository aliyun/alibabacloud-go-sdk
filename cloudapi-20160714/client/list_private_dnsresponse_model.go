// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateDNSResponseBody) *ListPrivateDNSResponse
	GetBody() *ListPrivateDNSResponseBody
}

type ListPrivateDNSResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateDNSResponse) GetBody() *ListPrivateDNSResponseBody {
	return s.Body
}

func (s *ListPrivateDNSResponse) SetHeaders(v map[string]*string) *ListPrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateDNSResponse) SetStatusCode(v int32) *ListPrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateDNSResponse) SetBody(v *ListPrivateDNSResponseBody) *ListPrivateDNSResponse {
	s.Body = v
	return s
}

func (s *ListPrivateDNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
