// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEndUserResponse
	GetStatusCode() *int32
	SetBody(v *AttachEndUserResponseBody) *AttachEndUserResponse
	GetBody() *AttachEndUserResponseBody
}

type AttachEndUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUserResponse) GoString() string {
	return s.String()
}

func (s *AttachEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEndUserResponse) GetBody() *AttachEndUserResponseBody {
	return s.Body
}

func (s *AttachEndUserResponse) SetHeaders(v map[string]*string) *AttachEndUserResponse {
	s.Headers = v
	return s
}

func (s *AttachEndUserResponse) SetStatusCode(v int32) *AttachEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEndUserResponse) SetBody(v *AttachEndUserResponseBody) *AttachEndUserResponse {
	s.Body = v
	return s
}

func (s *AttachEndUserResponse) Validate() error {
	return dara.Validate(s)
}
