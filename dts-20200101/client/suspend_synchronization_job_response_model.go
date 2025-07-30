// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendSynchronizationJobResponseBody) *SuspendSynchronizationJobResponse
	GetBody() *SuspendSynchronizationJobResponseBody
}

type SuspendSynchronizationJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendSynchronizationJobResponse) GetBody() *SuspendSynchronizationJobResponseBody {
	return s.Body
}

func (s *SuspendSynchronizationJobResponse) SetHeaders(v map[string]*string) *SuspendSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendSynchronizationJobResponse) SetStatusCode(v int32) *SuspendSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendSynchronizationJobResponse) SetBody(v *SuspendSynchronizationJobResponseBody) *SuspendSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *SuspendSynchronizationJobResponse) Validate() error {
	return dara.Validate(s)
}
