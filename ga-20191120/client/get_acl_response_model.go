// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAclResponse
	GetStatusCode() *int32
	SetBody(v *GetAclResponseBody) *GetAclResponse
	GetBody() *GetAclResponseBody
}

type GetAclResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAclResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAclResponse) GoString() string {
	return s.String()
}

func (s *GetAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAclResponse) GetBody() *GetAclResponseBody {
	return s.Body
}

func (s *GetAclResponse) SetHeaders(v map[string]*string) *GetAclResponse {
	s.Headers = v
	return s
}

func (s *GetAclResponse) SetStatusCode(v int32) *GetAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAclResponse) SetBody(v *GetAclResponseBody) *GetAclResponse {
	s.Body = v
	return s
}

func (s *GetAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
