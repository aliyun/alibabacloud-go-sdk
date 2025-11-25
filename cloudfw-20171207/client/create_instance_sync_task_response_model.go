// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceSyncTaskResponseBody) *CreateInstanceSyncTaskResponse
	GetBody() *CreateInstanceSyncTaskResponseBody
}

type CreateInstanceSyncTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceSyncTaskResponse) GetBody() *CreateInstanceSyncTaskResponseBody {
	return s.Body
}

func (s *CreateInstanceSyncTaskResponse) SetHeaders(v map[string]*string) *CreateInstanceSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceSyncTaskResponse) SetStatusCode(v int32) *CreateInstanceSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceSyncTaskResponse) SetBody(v *CreateInstanceSyncTaskResponseBody) *CreateInstanceSyncTaskResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceSyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
