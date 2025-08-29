// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetEngineConfigResponseBody) *GetEngineConfigResponse
	GetBody() *GetEngineConfigResponseBody
}

type GetEngineConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *GetEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEngineConfigResponse) GetBody() *GetEngineConfigResponseBody {
	return s.Body
}

func (s *GetEngineConfigResponse) SetHeaders(v map[string]*string) *GetEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *GetEngineConfigResponse) SetStatusCode(v int32) *GetEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEngineConfigResponse) SetBody(v *GetEngineConfigResponseBody) *GetEngineConfigResponse {
	s.Body = v
	return s
}

func (s *GetEngineConfigResponse) Validate() error {
	return dara.Validate(s)
}
