// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInventorySynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunInventorySynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunInventorySynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *RunInventorySynchronizationJobResponseBody) *RunInventorySynchronizationJobResponse
	GetBody() *RunInventorySynchronizationJobResponseBody
}

type RunInventorySynchronizationJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunInventorySynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunInventorySynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RunInventorySynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *RunInventorySynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunInventorySynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunInventorySynchronizationJobResponse) GetBody() *RunInventorySynchronizationJobResponseBody {
	return s.Body
}

func (s *RunInventorySynchronizationJobResponse) SetHeaders(v map[string]*string) *RunInventorySynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *RunInventorySynchronizationJobResponse) SetStatusCode(v int32) *RunInventorySynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunInventorySynchronizationJobResponse) SetBody(v *RunInventorySynchronizationJobResponseBody) *RunInventorySynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *RunInventorySynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
