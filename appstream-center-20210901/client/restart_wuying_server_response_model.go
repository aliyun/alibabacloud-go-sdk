// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *RestartWuyingServerResponseBody) *RestartWuyingServerResponse
	GetBody() *RestartWuyingServerResponseBody
}

type RestartWuyingServerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *RestartWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartWuyingServerResponse) GetBody() *RestartWuyingServerResponseBody {
	return s.Body
}

func (s *RestartWuyingServerResponse) SetHeaders(v map[string]*string) *RestartWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *RestartWuyingServerResponse) SetStatusCode(v int32) *RestartWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartWuyingServerResponse) SetBody(v *RestartWuyingServerResponseBody) *RestartWuyingServerResponse {
	s.Body = v
	return s
}

func (s *RestartWuyingServerResponse) Validate() error {
	return dara.Validate(s)
}
