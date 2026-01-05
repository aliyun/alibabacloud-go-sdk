// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntriesToAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEntriesToAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEntriesToAclResponse
	GetStatusCode() *int32
	SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse
	GetBody() *AddEntriesToAclResponseBody
}

type AddEntriesToAclResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEntriesToAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEntriesToAclResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclResponse) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEntriesToAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEntriesToAclResponse) GetBody() *AddEntriesToAclResponseBody {
	return s.Body
}

func (s *AddEntriesToAclResponse) SetHeaders(v map[string]*string) *AddEntriesToAclResponse {
	s.Headers = v
	return s
}

func (s *AddEntriesToAclResponse) SetStatusCode(v int32) *AddEntriesToAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEntriesToAclResponse) SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse {
	s.Body = v
	return s
}

func (s *AddEntriesToAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
