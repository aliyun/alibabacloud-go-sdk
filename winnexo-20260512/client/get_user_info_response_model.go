// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse
	GetBody() *GetUserInfoResponseBody
}

type GetUserInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserInfoResponse) GetBody() *GetUserInfoResponseBody {
	return s.Body
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetStatusCode(v int32) *GetUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
