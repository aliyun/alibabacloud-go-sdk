// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigResponseBody) *GetConfigResponse
	GetBody() *GetConfigResponseBody
}

type GetConfigResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigResponse) GoString() string {
	return s.String()
}

func (s *GetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigResponse) GetBody() *GetConfigResponseBody {
	return s.Body
}

func (s *GetConfigResponse) SetHeaders(v map[string]*string) *GetConfigResponse {
	s.Headers = v
	return s
}

func (s *GetConfigResponse) SetStatusCode(v int32) *GetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigResponse) SetBody(v *GetConfigResponseBody) *GetConfigResponse {
	s.Body = v
	return s
}

func (s *GetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
