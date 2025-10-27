// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaClusterPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpaClusterPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpaClusterPluginResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpaClusterPluginResponseBody) *CreateOpaClusterPluginResponse
	GetBody() *CreateOpaClusterPluginResponseBody
}

type CreateOpaClusterPluginResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpaClusterPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpaClusterPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaClusterPluginResponse) GoString() string {
	return s.String()
}

func (s *CreateOpaClusterPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpaClusterPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpaClusterPluginResponse) GetBody() *CreateOpaClusterPluginResponseBody {
	return s.Body
}

func (s *CreateOpaClusterPluginResponse) SetHeaders(v map[string]*string) *CreateOpaClusterPluginResponse {
	s.Headers = v
	return s
}

func (s *CreateOpaClusterPluginResponse) SetStatusCode(v int32) *CreateOpaClusterPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpaClusterPluginResponse) SetBody(v *CreateOpaClusterPluginResponseBody) *CreateOpaClusterPluginResponse {
	s.Body = v
	return s
}

func (s *CreateOpaClusterPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
