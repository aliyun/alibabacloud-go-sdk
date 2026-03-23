// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupSsidConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApgroupSsidConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApgroupSsidConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApgroupSsidConfigResponseBody) *DeleteApgroupSsidConfigResponse
	GetBody() *DeleteApgroupSsidConfigResponseBody
}

type DeleteApgroupSsidConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApgroupSsidConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApgroupSsidConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupSsidConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteApgroupSsidConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApgroupSsidConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApgroupSsidConfigResponse) GetBody() *DeleteApgroupSsidConfigResponseBody {
	return s.Body
}

func (s *DeleteApgroupSsidConfigResponse) SetHeaders(v map[string]*string) *DeleteApgroupSsidConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteApgroupSsidConfigResponse) SetStatusCode(v int32) *DeleteApgroupSsidConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponse) SetBody(v *DeleteApgroupSsidConfigResponseBody) *DeleteApgroupSsidConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteApgroupSsidConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
