// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserPoolSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserPoolSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *GetUserPoolSyncJobResponseBody) *GetUserPoolSyncJobResponse
	GetBody() *GetUserPoolSyncJobResponseBody
}

type GetUserPoolSyncJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPoolSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPoolSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolSyncJobResponse) GoString() string {
	return s.String()
}

func (s *GetUserPoolSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserPoolSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserPoolSyncJobResponse) GetBody() *GetUserPoolSyncJobResponseBody {
	return s.Body
}

func (s *GetUserPoolSyncJobResponse) SetHeaders(v map[string]*string) *GetUserPoolSyncJobResponse {
	s.Headers = v
	return s
}

func (s *GetUserPoolSyncJobResponse) SetStatusCode(v int32) *GetUserPoolSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPoolSyncJobResponse) SetBody(v *GetUserPoolSyncJobResponseBody) *GetUserPoolSyncJobResponse {
	s.Body = v
	return s
}

func (s *GetUserPoolSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
