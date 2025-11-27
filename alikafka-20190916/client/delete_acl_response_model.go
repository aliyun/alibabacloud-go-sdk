// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAclResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAclResponseBody) *DeleteAclResponse
	GetBody() *DeleteAclResponseBody
}

type DeleteAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAclResponse) GetBody() *DeleteAclResponseBody {
	return s.Body
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclResponse) SetStatusCode(v int32) *DeleteAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

func (s *DeleteAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
