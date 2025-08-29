// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *StopSampleConsistencyJobResponseBody) *StopSampleConsistencyJobResponse
	GetBody() *StopSampleConsistencyJobResponseBody
}

type StopSampleConsistencyJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *StopSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSampleConsistencyJobResponse) GetBody() *StopSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *StopSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *StopSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *StopSampleConsistencyJobResponse) SetStatusCode(v int32) *StopSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSampleConsistencyJobResponse) SetBody(v *StopSampleConsistencyJobResponseBody) *StopSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *StopSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
