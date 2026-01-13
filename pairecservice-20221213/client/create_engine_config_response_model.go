// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateEngineConfigResponseBody) *CreateEngineConfigResponse
	GetBody() *CreateEngineConfigResponseBody
}

type CreateEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEngineConfigResponse) GetBody() *CreateEngineConfigResponseBody {
	return s.Body
}

func (s *CreateEngineConfigResponse) SetHeaders(v map[string]*string) *CreateEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateEngineConfigResponse) SetStatusCode(v int32) *CreateEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEngineConfigResponse) SetBody(v *CreateEngineConfigResponseBody) *CreateEngineConfigResponse {
	s.Body = v
	return s
}

func (s *CreateEngineConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
