// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetSynchronizationJobResponseBody) *GetSynchronizationJobResponse
	GetBody() *GetSynchronizationJobResponseBody
}

type GetSynchronizationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSynchronizationJobResponse) GetBody() *GetSynchronizationJobResponseBody {
	return s.Body
}

func (s *GetSynchronizationJobResponse) SetHeaders(v map[string]*string) *GetSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *GetSynchronizationJobResponse) SetStatusCode(v int32) *GetSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSynchronizationJobResponse) SetBody(v *GetSynchronizationJobResponseBody) *GetSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *GetSynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
