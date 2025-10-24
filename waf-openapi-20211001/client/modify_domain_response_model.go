// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDomainResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDomainResponseBody) *ModifyDomainResponse
	GetBody() *ModifyDomainResponseBody
}

type ModifyDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDomainResponse) GetBody() *ModifyDomainResponseBody {
	return s.Body
}

func (s *ModifyDomainResponse) SetHeaders(v map[string]*string) *ModifyDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainResponse) SetStatusCode(v int32) *ModifyDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainResponse) SetBody(v *ModifyDomainResponseBody) *ModifyDomainResponse {
	s.Body = v
	return s
}

func (s *ModifyDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
