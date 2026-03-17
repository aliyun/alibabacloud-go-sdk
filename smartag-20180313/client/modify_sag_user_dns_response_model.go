// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagUserDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagUserDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagUserDnsResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagUserDnsResponseBody) *ModifySagUserDnsResponse
	GetBody() *ModifySagUserDnsResponseBody
}

type ModifySagUserDnsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagUserDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagUserDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagUserDnsResponse) GoString() string {
	return s.String()
}

func (s *ModifySagUserDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagUserDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagUserDnsResponse) GetBody() *ModifySagUserDnsResponseBody {
	return s.Body
}

func (s *ModifySagUserDnsResponse) SetHeaders(v map[string]*string) *ModifySagUserDnsResponse {
	s.Headers = v
	return s
}

func (s *ModifySagUserDnsResponse) SetStatusCode(v int32) *ModifySagUserDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagUserDnsResponse) SetBody(v *ModifySagUserDnsResponseBody) *ModifySagUserDnsResponse {
	s.Body = v
	return s
}

func (s *ModifySagUserDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
