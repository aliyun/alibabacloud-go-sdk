// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInventorySynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInventorySynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetInventorySynchronizationJobResponseBody) *GetInventorySynchronizationJobResponse
	GetBody() *GetInventorySynchronizationJobResponseBody
}

type GetInventorySynchronizationJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInventorySynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInventorySynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInventorySynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInventorySynchronizationJobResponse) GetBody() *GetInventorySynchronizationJobResponseBody {
	return s.Body
}

func (s *GetInventorySynchronizationJobResponse) SetHeaders(v map[string]*string) *GetInventorySynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *GetInventorySynchronizationJobResponse) SetStatusCode(v int32) *GetInventorySynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventorySynchronizationJobResponse) SetBody(v *GetInventorySynchronizationJobResponseBody) *GetInventorySynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *GetInventorySynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
