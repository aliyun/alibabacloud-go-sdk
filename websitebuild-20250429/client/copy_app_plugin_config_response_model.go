// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppPluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyAppPluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyAppPluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyAppPluginConfigResponseBody) *CopyAppPluginConfigResponse
	GetBody() *CopyAppPluginConfigResponseBody
}

type CopyAppPluginConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyAppPluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyAppPluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyAppPluginConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyAppPluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyAppPluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyAppPluginConfigResponse) GetBody() *CopyAppPluginConfigResponseBody {
	return s.Body
}

func (s *CopyAppPluginConfigResponse) SetHeaders(v map[string]*string) *CopyAppPluginConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyAppPluginConfigResponse) SetStatusCode(v int32) *CopyAppPluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyAppPluginConfigResponse) SetBody(v *CopyAppPluginConfigResponseBody) *CopyAppPluginConfigResponse {
	s.Body = v
	return s
}

func (s *CopyAppPluginConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
