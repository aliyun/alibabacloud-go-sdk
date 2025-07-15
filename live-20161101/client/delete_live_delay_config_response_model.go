// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveDelayConfigResponseBody) *DeleteLiveDelayConfigResponse
	GetBody() *DeleteLiveDelayConfigResponseBody
}

type DeleteLiveDelayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveDelayConfigResponse) GetBody() *DeleteLiveDelayConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveDelayConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveDelayConfigResponse) SetStatusCode(v int32) *DeleteLiveDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveDelayConfigResponse) SetBody(v *DeleteLiveDelayConfigResponseBody) *DeleteLiveDelayConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveDelayConfigResponse) Validate() error {
	return dara.Validate(s)
}
