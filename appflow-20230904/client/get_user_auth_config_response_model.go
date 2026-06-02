// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAuthConfigResponseBody) *GetUserAuthConfigResponse
	GetBody() *GetUserAuthConfigResponseBody
}

type GetUserAuthConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAuthConfigResponse) GetBody() *GetUserAuthConfigResponseBody {
	return s.Body
}

func (s *GetUserAuthConfigResponse) SetHeaders(v map[string]*string) *GetUserAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserAuthConfigResponse) SetStatusCode(v int32) *GetUserAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAuthConfigResponse) SetBody(v *GetUserAuthConfigResponseBody) *GetUserAuthConfigResponse {
	s.Body = v
	return s
}

func (s *GetUserAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
