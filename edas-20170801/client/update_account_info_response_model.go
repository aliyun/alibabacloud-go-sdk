// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountInfoResponseBody) *UpdateAccountInfoResponse
	GetBody() *UpdateAccountInfoResponseBody
}

type UpdateAccountInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountInfoResponse) GetBody() *UpdateAccountInfoResponseBody {
	return s.Body
}

func (s *UpdateAccountInfoResponse) SetHeaders(v map[string]*string) *UpdateAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountInfoResponse) SetStatusCode(v int32) *UpdateAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountInfoResponse) SetBody(v *UpdateAccountInfoResponseBody) *UpdateAccountInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
