// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendDtsJobResponseBody) *SuspendDtsJobResponse
	GetBody() *SuspendDtsJobResponseBody
}

type SuspendDtsJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendDtsJobResponse) GetBody() *SuspendDtsJobResponseBody {
	return s.Body
}

func (s *SuspendDtsJobResponse) SetHeaders(v map[string]*string) *SuspendDtsJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendDtsJobResponse) SetStatusCode(v int32) *SuspendDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendDtsJobResponse) SetBody(v *SuspendDtsJobResponseBody) *SuspendDtsJobResponse {
	s.Body = v
	return s
}

func (s *SuspendDtsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
