// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendApsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendApsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendApsJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendApsJobResponseBody) *SuspendApsJobResponse
	GetBody() *SuspendApsJobResponseBody
}

type SuspendApsJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendApsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendApsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendApsJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendApsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendApsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendApsJobResponse) GetBody() *SuspendApsJobResponseBody {
	return s.Body
}

func (s *SuspendApsJobResponse) SetHeaders(v map[string]*string) *SuspendApsJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendApsJobResponse) SetStatusCode(v int32) *SuspendApsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendApsJobResponse) SetBody(v *SuspendApsJobResponseBody) *SuspendApsJobResponse {
	s.Body = v
	return s
}

func (s *SuspendApsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
