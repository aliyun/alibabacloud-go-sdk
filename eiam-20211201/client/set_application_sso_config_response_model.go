// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationSsoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationSsoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationSsoConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationSsoConfigResponseBody) *SetApplicationSsoConfigResponse
	GetBody() *SetApplicationSsoConfigResponseBody
}

type SetApplicationSsoConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationSsoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationSsoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationSsoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationSsoConfigResponse) GetBody() *SetApplicationSsoConfigResponseBody {
	return s.Body
}

func (s *SetApplicationSsoConfigResponse) SetHeaders(v map[string]*string) *SetApplicationSsoConfigResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationSsoConfigResponse) SetStatusCode(v int32) *SetApplicationSsoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationSsoConfigResponse) SetBody(v *SetApplicationSsoConfigResponseBody) *SetApplicationSsoConfigResponse {
	s.Body = v
	return s
}

func (s *SetApplicationSsoConfigResponse) Validate() error {
	return dara.Validate(s)
}
