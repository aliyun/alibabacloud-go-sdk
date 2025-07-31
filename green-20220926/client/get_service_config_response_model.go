// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceConfigResponseBody) *GetServiceConfigResponse
	GetBody() *GetServiceConfigResponseBody
}

type GetServiceConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceConfigResponse) GetBody() *GetServiceConfigResponseBody {
	return s.Body
}

func (s *GetServiceConfigResponse) SetHeaders(v map[string]*string) *GetServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConfigResponse) SetStatusCode(v int32) *GetServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceConfigResponse) SetBody(v *GetServiceConfigResponseBody) *GetServiceConfigResponse {
	s.Body = v
	return s
}

func (s *GetServiceConfigResponse) Validate() error {
	return dara.Validate(s)
}
