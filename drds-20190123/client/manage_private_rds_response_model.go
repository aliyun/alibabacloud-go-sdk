// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagePrivateRdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManagePrivateRdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManagePrivateRdsResponse
	GetStatusCode() *int32
	SetBody(v *ManagePrivateRdsResponseBody) *ManagePrivateRdsResponse
	GetBody() *ManagePrivateRdsResponseBody
}

type ManagePrivateRdsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManagePrivateRdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManagePrivateRdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ManagePrivateRdsResponse) GoString() string {
	return s.String()
}

func (s *ManagePrivateRdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManagePrivateRdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManagePrivateRdsResponse) GetBody() *ManagePrivateRdsResponseBody {
	return s.Body
}

func (s *ManagePrivateRdsResponse) SetHeaders(v map[string]*string) *ManagePrivateRdsResponse {
	s.Headers = v
	return s
}

func (s *ManagePrivateRdsResponse) SetStatusCode(v int32) *ManagePrivateRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ManagePrivateRdsResponse) SetBody(v *ManagePrivateRdsResponseBody) *ManagePrivateRdsResponse {
	s.Body = v
	return s
}

func (s *ManagePrivateRdsResponse) Validate() error {
	return dara.Validate(s)
}
