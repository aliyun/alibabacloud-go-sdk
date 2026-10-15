// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppPluginResponseBody) *DeleteAppPluginResponse
	GetBody() *DeleteAppPluginResponseBody
}

type DeleteAppPluginResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPluginResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppPluginResponse) GetBody() *DeleteAppPluginResponseBody {
	return s.Body
}

func (s *DeleteAppPluginResponse) SetHeaders(v map[string]*string) *DeleteAppPluginResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppPluginResponse) SetStatusCode(v int32) *DeleteAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppPluginResponse) SetBody(v *DeleteAppPluginResponseBody) *DeleteAppPluginResponse {
	s.Body = v
	return s
}

func (s *DeleteAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
