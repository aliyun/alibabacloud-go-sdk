// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePluginResponse
	GetStatusCode() *int32
	SetBody(v *CreatePluginResponseBody) *CreatePluginResponse
	GetBody() *CreatePluginResponseBody
}

type CreatePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePluginResponse) GetBody() *CreatePluginResponseBody {
	return s.Body
}

func (s *CreatePluginResponse) SetHeaders(v map[string]*string) *CreatePluginResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginResponse) SetStatusCode(v int32) *CreatePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePluginResponse) SetBody(v *CreatePluginResponseBody) *CreatePluginResponse {
	s.Body = v
	return s
}

func (s *CreatePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
