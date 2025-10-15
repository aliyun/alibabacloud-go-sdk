// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationSsoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationSsoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationSsoConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationSsoConfigResponseBody) *GetApplicationSsoConfigResponse
	GetBody() *GetApplicationSsoConfigResponseBody
}

type GetApplicationSsoConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationSsoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationSsoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationSsoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationSsoConfigResponse) GetBody() *GetApplicationSsoConfigResponseBody {
	return s.Body
}

func (s *GetApplicationSsoConfigResponse) SetHeaders(v map[string]*string) *GetApplicationSsoConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationSsoConfigResponse) SetStatusCode(v int32) *GetApplicationSsoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationSsoConfigResponse) SetBody(v *GetApplicationSsoConfigResponseBody) *GetApplicationSsoConfigResponse {
	s.Body = v
	return s
}

func (s *GetApplicationSsoConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
