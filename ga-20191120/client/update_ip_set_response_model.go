// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpSetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpSetResponseBody) *UpdateIpSetResponse
	GetBody() *UpdateIpSetResponseBody
}

type UpdateIpSetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpSetResponse) GetBody() *UpdateIpSetResponseBody {
	return s.Body
}

func (s *UpdateIpSetResponse) SetHeaders(v map[string]*string) *UpdateIpSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpSetResponse) SetStatusCode(v int32) *UpdateIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpSetResponse) SetBody(v *UpdateIpSetResponseBody) *UpdateIpSetResponse {
	s.Body = v
	return s
}

func (s *UpdateIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
