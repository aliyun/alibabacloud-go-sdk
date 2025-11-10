// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAclResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAclResponseBody) *GetInstanceAclResponse
	GetBody() *GetInstanceAclResponseBody
}

type GetInstanceAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAclResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAclResponse) GetBody() *GetInstanceAclResponseBody {
	return s.Body
}

func (s *GetInstanceAclResponse) SetHeaders(v map[string]*string) *GetInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAclResponse) SetStatusCode(v int32) *GetInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAclResponse) SetBody(v *GetInstanceAclResponseBody) *GetInstanceAclResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
