// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEngineConfigResponseBody) *DeleteEngineConfigResponse
	GetBody() *DeleteEngineConfigResponseBody
}

type DeleteEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEngineConfigResponse) GetBody() *DeleteEngineConfigResponseBody {
	return s.Body
}

func (s *DeleteEngineConfigResponse) SetHeaders(v map[string]*string) *DeleteEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteEngineConfigResponse) SetStatusCode(v int32) *DeleteEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEngineConfigResponse) SetBody(v *DeleteEngineConfigResponseBody) *DeleteEngineConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteEngineConfigResponse) Validate() error {
	return dara.Validate(s)
}
