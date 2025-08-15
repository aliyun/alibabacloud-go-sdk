// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrailResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrailResponseBody) *DeleteTrailResponse
	GetBody() *DeleteTrailResponseBody
}

type DeleteTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrailResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrailResponse) GetBody() *DeleteTrailResponseBody {
	return s.Body
}

func (s *DeleteTrailResponse) SetHeaders(v map[string]*string) *DeleteTrailResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrailResponse) SetStatusCode(v int32) *DeleteTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrailResponse) SetBody(v *DeleteTrailResponseBody) *DeleteTrailResponse {
	s.Body = v
	return s
}

func (s *DeleteTrailResponse) Validate() error {
	return dara.Validate(s)
}
