// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountProfileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccountProfileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccountProfileInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccountProfileInfoResponseBody) *CreateAccountProfileInfoResponse
	GetBody() *CreateAccountProfileInfoResponseBody
}

type CreateAccountProfileInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountProfileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountProfileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountProfileInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountProfileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccountProfileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccountProfileInfoResponse) GetBody() *CreateAccountProfileInfoResponseBody {
	return s.Body
}

func (s *CreateAccountProfileInfoResponse) SetHeaders(v map[string]*string) *CreateAccountProfileInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountProfileInfoResponse) SetStatusCode(v int32) *CreateAccountProfileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountProfileInfoResponse) SetBody(v *CreateAccountProfileInfoResponseBody) *CreateAccountProfileInfoResponse {
	s.Body = v
	return s
}

func (s *CreateAccountProfileInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
