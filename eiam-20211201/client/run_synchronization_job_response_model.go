// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *RunSynchronizationJobResponseBody) *RunSynchronizationJobResponse
	GetBody() *RunSynchronizationJobResponseBody
}

type RunSynchronizationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSynchronizationJobResponse) GetBody() *RunSynchronizationJobResponseBody {
	return s.Body
}

func (s *RunSynchronizationJobResponse) SetHeaders(v map[string]*string) *RunSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *RunSynchronizationJobResponse) SetStatusCode(v int32) *RunSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSynchronizationJobResponse) SetBody(v *RunSynchronizationJobResponseBody) *RunSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *RunSynchronizationJobResponse) Validate() error {
	return dara.Validate(s)
}
