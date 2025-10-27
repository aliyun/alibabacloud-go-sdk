// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRestoreJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRestoreJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateRestoreJobResponseBody) *CreateRestoreJobResponse
	GetBody() *CreateRestoreJobResponseBody
}

type CreateRestoreJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRestoreJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRestoreJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRestoreJobResponse) GetBody() *CreateRestoreJobResponseBody {
	return s.Body
}

func (s *CreateRestoreJobResponse) SetHeaders(v map[string]*string) *CreateRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRestoreJobResponse) SetStatusCode(v int32) *CreateRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRestoreJobResponse) SetBody(v *CreateRestoreJobResponseBody) *CreateRestoreJobResponse {
	s.Body = v
	return s
}

func (s *CreateRestoreJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
