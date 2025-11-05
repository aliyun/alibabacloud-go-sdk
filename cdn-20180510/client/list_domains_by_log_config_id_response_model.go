// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsByLogConfigIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainsByLogConfigIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainsByLogConfigIdResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainsByLogConfigIdResponseBody) *ListDomainsByLogConfigIdResponse
	GetBody() *ListDomainsByLogConfigIdResponseBody
}

type ListDomainsByLogConfigIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsByLogConfigIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsByLogConfigIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainsByLogConfigIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainsByLogConfigIdResponse) GetBody() *ListDomainsByLogConfigIdResponseBody {
	return s.Body
}

func (s *ListDomainsByLogConfigIdResponse) SetHeaders(v map[string]*string) *ListDomainsByLogConfigIdResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsByLogConfigIdResponse) SetStatusCode(v int32) *ListDomainsByLogConfigIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsByLogConfigIdResponse) SetBody(v *ListDomainsByLogConfigIdResponseBody) *ListDomainsByLogConfigIdResponse {
	s.Body = v
	return s
}

func (s *ListDomainsByLogConfigIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
