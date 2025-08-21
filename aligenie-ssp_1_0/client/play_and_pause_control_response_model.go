// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayAndPauseControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PlayAndPauseControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PlayAndPauseControlResponse
	GetStatusCode() *int32
	SetBody(v *PlayAndPauseControlResponseBody) *PlayAndPauseControlResponse
	GetBody() *PlayAndPauseControlResponseBody
}

type PlayAndPauseControlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlayAndPauseControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlayAndPauseControlResponse) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlResponse) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PlayAndPauseControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PlayAndPauseControlResponse) GetBody() *PlayAndPauseControlResponseBody {
	return s.Body
}

func (s *PlayAndPauseControlResponse) SetHeaders(v map[string]*string) *PlayAndPauseControlResponse {
	s.Headers = v
	return s
}

func (s *PlayAndPauseControlResponse) SetStatusCode(v int32) *PlayAndPauseControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayAndPauseControlResponse) SetBody(v *PlayAndPauseControlResponseBody) *PlayAndPauseControlResponse {
	s.Body = v
	return s
}

func (s *PlayAndPauseControlResponse) Validate() error {
	return dara.Validate(s)
}
