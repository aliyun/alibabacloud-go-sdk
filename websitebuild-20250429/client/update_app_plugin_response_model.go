// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppPluginResponseBody) *UpdateAppPluginResponse
	GetBody() *UpdateAppPluginResponseBody
}

type UpdateAppPluginResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPluginResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppPluginResponse) GetBody() *UpdateAppPluginResponseBody {
	return s.Body
}

func (s *UpdateAppPluginResponse) SetHeaders(v map[string]*string) *UpdateAppPluginResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppPluginResponse) SetStatusCode(v int32) *UpdateAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppPluginResponse) SetBody(v *UpdateAppPluginResponseBody) *UpdateAppPluginResponse {
	s.Body = v
	return s
}

func (s *UpdateAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
