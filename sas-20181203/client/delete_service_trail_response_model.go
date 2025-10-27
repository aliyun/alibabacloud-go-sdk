// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceTrailResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceTrailResponseBody) *DeleteServiceTrailResponse
	GetBody() *DeleteServiceTrailResponseBody
}

type DeleteServiceTrailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTrailResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceTrailResponse) GetBody() *DeleteServiceTrailResponseBody {
	return s.Body
}

func (s *DeleteServiceTrailResponse) SetHeaders(v map[string]*string) *DeleteServiceTrailResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceTrailResponse) SetStatusCode(v int32) *DeleteServiceTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceTrailResponse) SetBody(v *DeleteServiceTrailResponseBody) *DeleteServiceTrailResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceTrailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
