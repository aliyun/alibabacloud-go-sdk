// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodSpecificConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodSpecificConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodSpecificConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodSpecificConfigResponseBody) *DeleteVodSpecificConfigResponse
	GetBody() *DeleteVodSpecificConfigResponseBody
}

type DeleteVodSpecificConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodSpecificConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodSpecificConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodSpecificConfigResponse) GetBody() *DeleteVodSpecificConfigResponseBody {
	return s.Body
}

func (s *DeleteVodSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteVodSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodSpecificConfigResponse) SetStatusCode(v int32) *DeleteVodSpecificConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodSpecificConfigResponse) SetBody(v *DeleteVodSpecificConfigResponseBody) *DeleteVodSpecificConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteVodSpecificConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
