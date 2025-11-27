// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVsPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVsPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVsPullStreamInfoConfigResponseBody) *DeleteVsPullStreamInfoConfigResponse
	GetBody() *DeleteVsPullStreamInfoConfigResponseBody
}

type DeleteVsPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVsPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVsPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVsPullStreamInfoConfigResponse) GetBody() *DeleteVsPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *DeleteVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DeleteVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVsPullStreamInfoConfigResponse) SetStatusCode(v int32) *DeleteVsPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigResponse) SetBody(v *DeleteVsPullStreamInfoConfigResponseBody) *DeleteVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteVsPullStreamInfoConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
