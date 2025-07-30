// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *ResetSynchronizationJobResponseBody) *ResetSynchronizationJobResponse
	GetBody() *ResetSynchronizationJobResponseBody
}

type ResetSynchronizationJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSynchronizationJobResponse) GetBody() *ResetSynchronizationJobResponseBody {
	return s.Body
}

func (s *ResetSynchronizationJobResponse) SetHeaders(v map[string]*string) *ResetSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *ResetSynchronizationJobResponse) SetStatusCode(v int32) *ResetSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSynchronizationJobResponse) SetBody(v *ResetSynchronizationJobResponseBody) *ResetSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *ResetSynchronizationJobResponse) Validate() error {
	return dara.Validate(s)
}
