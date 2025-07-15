// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PollUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PollUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *PollUserStatusResponseBody) *PollUserStatusResponse
	GetBody() *PollUserStatusResponseBody
}

type PollUserStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PollUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PollUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusResponse) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PollUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PollUserStatusResponse) GetBody() *PollUserStatusResponseBody {
	return s.Body
}

func (s *PollUserStatusResponse) SetHeaders(v map[string]*string) *PollUserStatusResponse {
	s.Headers = v
	return s
}

func (s *PollUserStatusResponse) SetStatusCode(v int32) *PollUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PollUserStatusResponse) SetBody(v *PollUserStatusResponseBody) *PollUserStatusResponse {
	s.Body = v
	return s
}

func (s *PollUserStatusResponse) Validate() error {
	return dara.Validate(s)
}
