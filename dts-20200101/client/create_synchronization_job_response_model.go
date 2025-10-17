// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateSynchronizationJobResponseBody) *CreateSynchronizationJobResponse
	GetBody() *CreateSynchronizationJobResponseBody
}

type CreateSynchronizationJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSynchronizationJobResponse) GetBody() *CreateSynchronizationJobResponseBody {
	return s.Body
}

func (s *CreateSynchronizationJobResponse) SetHeaders(v map[string]*string) *CreateSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSynchronizationJobResponse) SetStatusCode(v int32) *CreateSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSynchronizationJobResponse) SetBody(v *CreateSynchronizationJobResponseBody) *CreateSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *CreateSynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
