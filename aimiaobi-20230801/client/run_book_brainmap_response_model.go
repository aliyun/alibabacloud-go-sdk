// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookBrainmapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunBookBrainmapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunBookBrainmapResponse
	GetStatusCode() *int32
	SetBody(v *RunBookBrainmapResponseBody) *RunBookBrainmapResponse
	GetBody() *RunBookBrainmapResponseBody
}

type RunBookBrainmapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunBookBrainmapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunBookBrainmapResponse) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponse) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunBookBrainmapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunBookBrainmapResponse) GetBody() *RunBookBrainmapResponseBody {
	return s.Body
}

func (s *RunBookBrainmapResponse) SetHeaders(v map[string]*string) *RunBookBrainmapResponse {
	s.Headers = v
	return s
}

func (s *RunBookBrainmapResponse) SetStatusCode(v int32) *RunBookBrainmapResponse {
	s.StatusCode = &v
	return s
}

func (s *RunBookBrainmapResponse) SetBody(v *RunBookBrainmapResponseBody) *RunBookBrainmapResponse {
	s.Body = v
	return s
}

func (s *RunBookBrainmapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
