// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficControlResponseBody) *DeleteTrafficControlResponse
	GetBody() *DeleteTrafficControlResponseBody
}

type DeleteTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficControlResponse) GetBody() *DeleteTrafficControlResponseBody {
	return s.Body
}

func (s *DeleteTrafficControlResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlResponse) SetStatusCode(v int32) *DeleteTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlResponse) SetBody(v *DeleteTrafficControlResponseBody) *DeleteTrafficControlResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficControlResponse) Validate() error {
	return dara.Validate(s)
}
