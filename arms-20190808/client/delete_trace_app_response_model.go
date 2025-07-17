// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTraceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTraceAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTraceAppResponseBody) *DeleteTraceAppResponse
	GetBody() *DeleteTraceAppResponseBody
}

type DeleteTraceAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTraceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTraceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTraceAppResponse) GetBody() *DeleteTraceAppResponseBody {
	return s.Body
}

func (s *DeleteTraceAppResponse) SetHeaders(v map[string]*string) *DeleteTraceAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteTraceAppResponse) SetStatusCode(v int32) *DeleteTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTraceAppResponse) SetBody(v *DeleteTraceAppResponseBody) *DeleteTraceAppResponse {
	s.Body = v
	return s
}

func (s *DeleteTraceAppResponse) Validate() error {
	return dara.Validate(s)
}
