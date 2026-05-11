// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRecordDurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobRecordDurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobRecordDurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobRecordDurationResponseBody) *UpdateJobRecordDurationResponse
	GetBody() *UpdateJobRecordDurationResponseBody
}

type UpdateJobRecordDurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobRecordDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobRecordDurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRecordDurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobRecordDurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobRecordDurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobRecordDurationResponse) GetBody() *UpdateJobRecordDurationResponseBody {
	return s.Body
}

func (s *UpdateJobRecordDurationResponse) SetHeaders(v map[string]*string) *UpdateJobRecordDurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobRecordDurationResponse) SetStatusCode(v int32) *UpdateJobRecordDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobRecordDurationResponse) SetBody(v *UpdateJobRecordDurationResponseBody) *UpdateJobRecordDurationResponse {
	s.Body = v
	return s
}

func (s *UpdateJobRecordDurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
