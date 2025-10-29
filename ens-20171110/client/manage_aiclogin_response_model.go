// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAICLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageAICLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageAICLoginResponse
	GetStatusCode() *int32
	SetBody(v *ManageAICLoginResponseBody) *ManageAICLoginResponse
	GetBody() *ManageAICLoginResponseBody
}

type ManageAICLoginResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageAICLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAICLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageAICLoginResponse) GoString() string {
	return s.String()
}

func (s *ManageAICLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageAICLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageAICLoginResponse) GetBody() *ManageAICLoginResponseBody {
	return s.Body
}

func (s *ManageAICLoginResponse) SetHeaders(v map[string]*string) *ManageAICLoginResponse {
	s.Headers = v
	return s
}

func (s *ManageAICLoginResponse) SetStatusCode(v int32) *ManageAICLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageAICLoginResponse) SetBody(v *ManageAICLoginResponseBody) *ManageAICLoginResponse {
	s.Body = v
	return s
}

func (s *ManageAICLoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
