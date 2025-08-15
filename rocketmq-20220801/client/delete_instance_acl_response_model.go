// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceAclResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceAclResponseBody) *DeleteInstanceAclResponse
	GetBody() *DeleteInstanceAclResponseBody
}

type DeleteInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceAclResponse) GetBody() *DeleteInstanceAclResponseBody {
	return s.Body
}

func (s *DeleteInstanceAclResponse) SetHeaders(v map[string]*string) *DeleteInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceAclResponse) SetStatusCode(v int32) *DeleteInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceAclResponse) SetBody(v *DeleteInstanceAclResponseBody) *DeleteInstanceAclResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceAclResponse) Validate() error {
	return dara.Validate(s)
}
