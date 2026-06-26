// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunUserPoolSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunUserPoolSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunUserPoolSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *RunUserPoolSyncJobResponseBody) *RunUserPoolSyncJobResponse
	GetBody() *RunUserPoolSyncJobResponseBody
}

type RunUserPoolSyncJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunUserPoolSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunUserPoolSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RunUserPoolSyncJobResponse) GoString() string {
	return s.String()
}

func (s *RunUserPoolSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunUserPoolSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunUserPoolSyncJobResponse) GetBody() *RunUserPoolSyncJobResponseBody {
	return s.Body
}

func (s *RunUserPoolSyncJobResponse) SetHeaders(v map[string]*string) *RunUserPoolSyncJobResponse {
	s.Headers = v
	return s
}

func (s *RunUserPoolSyncJobResponse) SetStatusCode(v int32) *RunUserPoolSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunUserPoolSyncJobResponse) SetBody(v *RunUserPoolSyncJobResponseBody) *RunUserPoolSyncJobResponse {
	s.Body = v
	return s
}

func (s *RunUserPoolSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
