// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePluginResponse
	GetStatusCode() *int32
	SetBody(v *DeletePluginResponseBody) *DeletePluginResponse
	GetBody() *DeletePluginResponseBody
}

type DeletePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginResponse) GoString() string {
	return s.String()
}

func (s *DeletePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePluginResponse) GetBody() *DeletePluginResponseBody {
	return s.Body
}

func (s *DeletePluginResponse) SetHeaders(v map[string]*string) *DeletePluginResponse {
	s.Headers = v
	return s
}

func (s *DeletePluginResponse) SetStatusCode(v int32) *DeletePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePluginResponse) SetBody(v *DeletePluginResponseBody) *DeletePluginResponse {
	s.Body = v
	return s
}

func (s *DeletePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
