// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *GetUserPasswordResponseBody) *GetUserPasswordResponse
	GetBody() *GetUserPasswordResponseBody
}

type GetUserPasswordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *GetUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserPasswordResponse) GetBody() *GetUserPasswordResponseBody {
	return s.Body
}

func (s *GetUserPasswordResponse) SetHeaders(v map[string]*string) *GetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *GetUserPasswordResponse) SetStatusCode(v int32) *GetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPasswordResponse) SetBody(v *GetUserPasswordResponseBody) *GetUserPasswordResponse {
	s.Body = v
	return s
}

func (s *GetUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
