// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetClientConfigResponseBody) *GetClientConfigResponse
	GetBody() *GetClientConfigResponseBody
}

type GetClientConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientConfigResponse) GoString() string {
	return s.String()
}

func (s *GetClientConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientConfigResponse) GetBody() *GetClientConfigResponseBody {
	return s.Body
}

func (s *GetClientConfigResponse) SetHeaders(v map[string]*string) *GetClientConfigResponse {
	s.Headers = v
	return s
}

func (s *GetClientConfigResponse) SetStatusCode(v int32) *GetClientConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientConfigResponse) SetBody(v *GetClientConfigResponseBody) *GetClientConfigResponse {
	s.Body = v
	return s
}

func (s *GetClientConfigResponse) Validate() error {
	return dara.Validate(s)
}
