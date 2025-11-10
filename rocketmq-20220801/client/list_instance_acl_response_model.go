// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceAclResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceAclResponseBody) *ListInstanceAclResponse
	GetBody() *ListInstanceAclResponseBody
}

type ListInstanceAclResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAclResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceAclResponse) GetBody() *ListInstanceAclResponseBody {
	return s.Body
}

func (s *ListInstanceAclResponse) SetHeaders(v map[string]*string) *ListInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAclResponse) SetStatusCode(v int32) *ListInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAclResponse) SetBody(v *ListInstanceAclResponseBody) *ListInstanceAclResponse {
	s.Body = v
	return s
}

func (s *ListInstanceAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
