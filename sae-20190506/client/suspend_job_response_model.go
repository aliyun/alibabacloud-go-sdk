// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendJobResponseBody) *SuspendJobResponse
	GetBody() *SuspendJobResponseBody
}

type SuspendJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendJobResponse) GetBody() *SuspendJobResponseBody {
	return s.Body
}

func (s *SuspendJobResponse) SetHeaders(v map[string]*string) *SuspendJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendJobResponse) SetStatusCode(v int32) *SuspendJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendJobResponse) SetBody(v *SuspendJobResponseBody) *SuspendJobResponse {
	s.Body = v
	return s
}

func (s *SuspendJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
