// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAclResponse
	GetStatusCode() *int32
	SetBody(v *CreateAclResponseBody) *CreateAclResponse
	GetBody() *CreateAclResponseBody
}

type CreateAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAclResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAclResponse) GoString() string {
	return s.String()
}

func (s *CreateAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAclResponse) GetBody() *CreateAclResponseBody {
	return s.Body
}

func (s *CreateAclResponse) SetHeaders(v map[string]*string) *CreateAclResponse {
	s.Headers = v
	return s
}

func (s *CreateAclResponse) SetStatusCode(v int32) *CreateAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

func (s *CreateAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
