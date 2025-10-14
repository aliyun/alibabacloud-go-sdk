// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddDomainGroupResponseBody) *AddDomainGroupResponse
	GetBody() *AddDomainGroupResponseBody
}

type AddDomainGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDomainGroupResponse) GetBody() *AddDomainGroupResponseBody {
	return s.Body
}

func (s *AddDomainGroupResponse) SetHeaders(v map[string]*string) *AddDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDomainGroupResponse) SetStatusCode(v int32) *AddDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainGroupResponse) SetBody(v *AddDomainGroupResponseBody) *AddDomainGroupResponse {
	s.Body = v
	return s
}

func (s *AddDomainGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
