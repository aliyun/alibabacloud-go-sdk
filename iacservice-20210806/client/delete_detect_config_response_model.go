// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDetectConfigResponseBody) *DeleteDetectConfigResponse
	GetBody() *DeleteDetectConfigResponseBody
}

type DeleteDetectConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDetectConfigResponse) GetBody() *DeleteDetectConfigResponseBody {
	return s.Body
}

func (s *DeleteDetectConfigResponse) SetHeaders(v map[string]*string) *DeleteDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDetectConfigResponse) SetStatusCode(v int32) *DeleteDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDetectConfigResponse) SetBody(v *DeleteDetectConfigResponseBody) *DeleteDetectConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
