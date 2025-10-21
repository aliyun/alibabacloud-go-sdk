// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotUpdateJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotUpdateJobResponse
	GetStatusCode() *int32
	SetBody(v *HotUpdateJobResponseBody) *HotUpdateJobResponse
	GetBody() *HotUpdateJobResponseBody
}

type HotUpdateJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotUpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotUpdateJobResponse) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobResponse) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotUpdateJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotUpdateJobResponse) GetBody() *HotUpdateJobResponseBody {
	return s.Body
}

func (s *HotUpdateJobResponse) SetHeaders(v map[string]*string) *HotUpdateJobResponse {
	s.Headers = v
	return s
}

func (s *HotUpdateJobResponse) SetStatusCode(v int32) *HotUpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *HotUpdateJobResponse) SetBody(v *HotUpdateJobResponseBody) *HotUpdateJobResponse {
	s.Body = v
	return s
}

func (s *HotUpdateJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
