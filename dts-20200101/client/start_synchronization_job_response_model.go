// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *StartSynchronizationJobResponseBody) *StartSynchronizationJobResponse
	GetBody() *StartSynchronizationJobResponseBody
}

type StartSynchronizationJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSynchronizationJobResponse) GetBody() *StartSynchronizationJobResponseBody {
	return s.Body
}

func (s *StartSynchronizationJobResponse) SetHeaders(v map[string]*string) *StartSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *StartSynchronizationJobResponse) SetStatusCode(v int32) *StartSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSynchronizationJobResponse) SetBody(v *StartSynchronizationJobResponseBody) *StartSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *StartSynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
