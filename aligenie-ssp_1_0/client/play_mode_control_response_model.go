// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PlayModeControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PlayModeControlResponse
	GetStatusCode() *int32
	SetBody(v *PlayModeControlResponseBody) *PlayModeControlResponse
	GetBody() *PlayModeControlResponseBody
}

type PlayModeControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlayModeControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlayModeControlResponse) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlResponse) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PlayModeControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PlayModeControlResponse) GetBody() *PlayModeControlResponseBody {
	return s.Body
}

func (s *PlayModeControlResponse) SetHeaders(v map[string]*string) *PlayModeControlResponse {
	s.Headers = v
	return s
}

func (s *PlayModeControlResponse) SetStatusCode(v int32) *PlayModeControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayModeControlResponse) SetBody(v *PlayModeControlResponseBody) *PlayModeControlResponse {
	s.Body = v
	return s
}

func (s *PlayModeControlResponse) Validate() error {
	return dara.Validate(s)
}
