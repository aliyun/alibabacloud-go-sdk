// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApgroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApgroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApgroupConfigResponseBody) *DeleteApgroupConfigResponse
	GetBody() *DeleteApgroupConfigResponseBody
}

type DeleteApgroupConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApgroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApgroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteApgroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApgroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApgroupConfigResponse) GetBody() *DeleteApgroupConfigResponseBody {
	return s.Body
}

func (s *DeleteApgroupConfigResponse) SetHeaders(v map[string]*string) *DeleteApgroupConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteApgroupConfigResponse) SetStatusCode(v int32) *DeleteApgroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApgroupConfigResponse) SetBody(v *DeleteApgroupConfigResponseBody) *DeleteApgroupConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteApgroupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
