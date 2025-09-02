// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCycleDagNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCycleDagNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCycleDagNodesResponse
	GetStatusCode() *int32
	SetBody(v *RunCycleDagNodesResponseBody) *RunCycleDagNodesResponse
	GetBody() *RunCycleDagNodesResponseBody
}

type RunCycleDagNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCycleDagNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCycleDagNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCycleDagNodesResponse) GoString() string {
	return s.String()
}

func (s *RunCycleDagNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCycleDagNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCycleDagNodesResponse) GetBody() *RunCycleDagNodesResponseBody {
	return s.Body
}

func (s *RunCycleDagNodesResponse) SetHeaders(v map[string]*string) *RunCycleDagNodesResponse {
	s.Headers = v
	return s
}

func (s *RunCycleDagNodesResponse) SetStatusCode(v int32) *RunCycleDagNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCycleDagNodesResponse) SetBody(v *RunCycleDagNodesResponseBody) *RunCycleDagNodesResponse {
	s.Body = v
	return s
}

func (s *RunCycleDagNodesResponse) Validate() error {
	return dara.Validate(s)
}
