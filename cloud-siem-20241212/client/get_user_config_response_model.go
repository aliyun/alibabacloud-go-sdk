// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetUserConfigResponseBody) *GetUserConfigResponse
	GetBody() *GetUserConfigResponseBody
}

type GetUserConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserConfigResponse) GetBody() *GetUserConfigResponseBody {
	return s.Body
}

func (s *GetUserConfigResponse) SetHeaders(v map[string]*string) *GetUserConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserConfigResponse) SetStatusCode(v int32) *GetUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserConfigResponse) SetBody(v *GetUserConfigResponseBody) *GetUserConfigResponse {
	s.Body = v
	return s
}

func (s *GetUserConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
