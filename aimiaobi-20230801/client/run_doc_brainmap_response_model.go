// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocBrainmapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocBrainmapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocBrainmapResponse
	GetStatusCode() *int32
	SetBody(v *RunDocBrainmapResponseBody) *RunDocBrainmapResponse
	GetBody() *RunDocBrainmapResponseBody
}

type RunDocBrainmapResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDocBrainmapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocBrainmapResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponse) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocBrainmapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocBrainmapResponse) GetBody() *RunDocBrainmapResponseBody {
	return s.Body
}

func (s *RunDocBrainmapResponse) SetHeaders(v map[string]*string) *RunDocBrainmapResponse {
	s.Headers = v
	return s
}

func (s *RunDocBrainmapResponse) SetStatusCode(v int32) *RunDocBrainmapResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocBrainmapResponse) SetBody(v *RunDocBrainmapResponseBody) *RunDocBrainmapResponse {
	s.Body = v
	return s
}

func (s *RunDocBrainmapResponse) Validate() error {
	return dara.Validate(s)
}
