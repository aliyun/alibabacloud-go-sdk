// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRecordDurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobRecordDurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobRecordDurationResponse
	GetStatusCode() *int32
	SetBody(v *GetJobRecordDurationResponseBody) *GetJobRecordDurationResponse
	GetBody() *GetJobRecordDurationResponseBody
}

type GetJobRecordDurationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobRecordDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobRecordDurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobRecordDurationResponse) GoString() string {
	return s.String()
}

func (s *GetJobRecordDurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobRecordDurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobRecordDurationResponse) GetBody() *GetJobRecordDurationResponseBody {
	return s.Body
}

func (s *GetJobRecordDurationResponse) SetHeaders(v map[string]*string) *GetJobRecordDurationResponse {
	s.Headers = v
	return s
}

func (s *GetJobRecordDurationResponse) SetStatusCode(v int32) *GetJobRecordDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobRecordDurationResponse) SetBody(v *GetJobRecordDurationResponseBody) *GetJobRecordDurationResponse {
	s.Body = v
	return s
}

func (s *GetJobRecordDurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
