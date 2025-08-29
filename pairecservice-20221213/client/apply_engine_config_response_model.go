// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *ApplyEngineConfigResponseBody) *ApplyEngineConfigResponse
	GetBody() *ApplyEngineConfigResponseBody
}

type ApplyEngineConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyEngineConfigResponse) GetBody() *ApplyEngineConfigResponseBody {
	return s.Body
}

func (s *ApplyEngineConfigResponse) SetHeaders(v map[string]*string) *ApplyEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *ApplyEngineConfigResponse) SetStatusCode(v int32) *ApplyEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyEngineConfigResponse) SetBody(v *ApplyEngineConfigResponseBody) *ApplyEngineConfigResponse {
	s.Body = v
	return s
}

func (s *ApplyEngineConfigResponse) Validate() error {
	return dara.Validate(s)
}
