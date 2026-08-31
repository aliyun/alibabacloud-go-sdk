// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByAccessKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserByAccessKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserByAccessKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetUserByAccessKeyResponseBody) *GetUserByAccessKeyResponse
	GetBody() *GetUserByAccessKeyResponseBody
}

type GetUserByAccessKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserByAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserByAccessKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserByAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *GetUserByAccessKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserByAccessKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserByAccessKeyResponse) GetBody() *GetUserByAccessKeyResponseBody {
	return s.Body
}

func (s *GetUserByAccessKeyResponse) SetHeaders(v map[string]*string) *GetUserByAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *GetUserByAccessKeyResponse) SetStatusCode(v int32) *GetUserByAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserByAccessKeyResponse) SetBody(v *GetUserByAccessKeyResponseBody) *GetUserByAccessKeyResponse {
	s.Body = v
	return s
}

func (s *GetUserByAccessKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
