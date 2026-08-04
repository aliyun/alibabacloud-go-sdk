// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountProfileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountProfileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountProfileInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountProfileInfoResponseBody) *UpdateAccountProfileInfoResponse
	GetBody() *UpdateAccountProfileInfoResponseBody
}

type UpdateAccountProfileInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountProfileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountProfileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountProfileInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountProfileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountProfileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountProfileInfoResponse) GetBody() *UpdateAccountProfileInfoResponseBody {
	return s.Body
}

func (s *UpdateAccountProfileInfoResponse) SetHeaders(v map[string]*string) *UpdateAccountProfileInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountProfileInfoResponse) SetStatusCode(v int32) *UpdateAccountProfileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountProfileInfoResponse) SetBody(v *UpdateAccountProfileInfoResponseBody) *UpdateAccountProfileInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountProfileInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
