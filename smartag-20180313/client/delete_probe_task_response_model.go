// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProbeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProbeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProbeTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProbeTaskResponseBody) *DeleteProbeTaskResponse
	GetBody() *DeleteProbeTaskResponseBody
}

type DeleteProbeTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProbeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProbeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProbeTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteProbeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProbeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProbeTaskResponse) GetBody() *DeleteProbeTaskResponseBody {
	return s.Body
}

func (s *DeleteProbeTaskResponse) SetHeaders(v map[string]*string) *DeleteProbeTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteProbeTaskResponse) SetStatusCode(v int32) *DeleteProbeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProbeTaskResponse) SetBody(v *DeleteProbeTaskResponseBody) *DeleteProbeTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteProbeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
