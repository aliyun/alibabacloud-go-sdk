// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdcProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIdcProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIdcProbeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIdcProbeResponseBody) *DeleteIdcProbeResponse
	GetBody() *DeleteIdcProbeResponseBody
}

type DeleteIdcProbeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIdcProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIdcProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdcProbeResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdcProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIdcProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIdcProbeResponse) GetBody() *DeleteIdcProbeResponseBody {
	return s.Body
}

func (s *DeleteIdcProbeResponse) SetHeaders(v map[string]*string) *DeleteIdcProbeResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdcProbeResponse) SetStatusCode(v int32) *DeleteIdcProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdcProbeResponse) SetBody(v *DeleteIdcProbeResponseBody) *DeleteIdcProbeResponse {
	s.Body = v
	return s
}

func (s *DeleteIdcProbeResponse) Validate() error {
	return dara.Validate(s)
}
