// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupThirdAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApgroupThirdAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApgroupThirdAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApgroupThirdAppResponseBody) *DeleteApgroupThirdAppResponse
	GetBody() *DeleteApgroupThirdAppResponseBody
}

type DeleteApgroupThirdAppResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApgroupThirdAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApgroupThirdAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupThirdAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteApgroupThirdAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApgroupThirdAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApgroupThirdAppResponse) GetBody() *DeleteApgroupThirdAppResponseBody {
	return s.Body
}

func (s *DeleteApgroupThirdAppResponse) SetHeaders(v map[string]*string) *DeleteApgroupThirdAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteApgroupThirdAppResponse) SetStatusCode(v int32) *DeleteApgroupThirdAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApgroupThirdAppResponse) SetBody(v *DeleteApgroupThirdAppResponseBody) *DeleteApgroupThirdAppResponse {
	s.Body = v
	return s
}

func (s *DeleteApgroupThirdAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
