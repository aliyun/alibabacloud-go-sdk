// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseTrafficControlTaskResponseBody) *ReleaseTrafficControlTaskResponse
	GetBody() *ReleaseTrafficControlTaskResponseBody
}

type ReleaseTrafficControlTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseTrafficControlTaskResponse) GetBody() *ReleaseTrafficControlTaskResponseBody {
	return s.Body
}

func (s *ReleaseTrafficControlTaskResponse) SetHeaders(v map[string]*string) *ReleaseTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *ReleaseTrafficControlTaskResponse) SetStatusCode(v int32) *ReleaseTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseTrafficControlTaskResponse) SetBody(v *ReleaseTrafficControlTaskResponseBody) *ReleaseTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *ReleaseTrafficControlTaskResponse) Validate() error {
	return dara.Validate(s)
}
