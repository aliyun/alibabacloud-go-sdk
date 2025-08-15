// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceAclResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceAclResponseBody) *CreateInstanceAclResponse
	GetBody() *CreateInstanceAclResponseBody
}

type CreateInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceAclResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceAclResponse) GetBody() *CreateInstanceAclResponseBody {
	return s.Body
}

func (s *CreateInstanceAclResponse) SetHeaders(v map[string]*string) *CreateInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAclResponse) SetStatusCode(v int32) *CreateInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceAclResponse) SetBody(v *CreateInstanceAclResponseBody) *CreateInstanceAclResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceAclResponse) Validate() error {
	return dara.Validate(s)
}
