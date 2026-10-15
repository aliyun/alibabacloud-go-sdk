// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppPluginResponseBody) *CreateAppPluginResponse
	GetBody() *CreateAppPluginResponseBody
}

type CreateAppPluginResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPluginResponse) GoString() string {
	return s.String()
}

func (s *CreateAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppPluginResponse) GetBody() *CreateAppPluginResponseBody {
	return s.Body
}

func (s *CreateAppPluginResponse) SetHeaders(v map[string]*string) *CreateAppPluginResponse {
	s.Headers = v
	return s
}

func (s *CreateAppPluginResponse) SetStatusCode(v int32) *CreateAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppPluginResponse) SetBody(v *CreateAppPluginResponseBody) *CreateAppPluginResponse {
	s.Body = v
	return s
}

func (s *CreateAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
