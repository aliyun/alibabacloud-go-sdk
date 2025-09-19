// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoneypotProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoneypotProbeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoneypotProbeResponseBody) *DeleteHoneypotProbeResponse
	GetBody() *DeleteHoneypotProbeResponseBody
}

type DeleteHoneypotProbeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoneypotProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoneypotProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoneypotProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoneypotProbeResponse) GetBody() *DeleteHoneypotProbeResponseBody {
	return s.Body
}

func (s *DeleteHoneypotProbeResponse) SetHeaders(v map[string]*string) *DeleteHoneypotProbeResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoneypotProbeResponse) SetStatusCode(v int32) *DeleteHoneypotProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoneypotProbeResponse) SetBody(v *DeleteHoneypotProbeResponseBody) *DeleteHoneypotProbeResponse {
	s.Body = v
	return s
}

func (s *DeleteHoneypotProbeResponse) Validate() error {
	return dara.Validate(s)
}
