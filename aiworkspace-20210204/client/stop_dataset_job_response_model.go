// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatasetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDatasetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDatasetJobResponse
	GetStatusCode() *int32
	SetBody(v *StopDatasetJobResponseBody) *StopDatasetJobResponse
	GetBody() *StopDatasetJobResponseBody
}

type StopDatasetJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDatasetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDatasetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDatasetJobResponse) GoString() string {
	return s.String()
}

func (s *StopDatasetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDatasetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDatasetJobResponse) GetBody() *StopDatasetJobResponseBody {
	return s.Body
}

func (s *StopDatasetJobResponse) SetHeaders(v map[string]*string) *StopDatasetJobResponse {
	s.Headers = v
	return s
}

func (s *StopDatasetJobResponse) SetStatusCode(v int32) *StopDatasetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDatasetJobResponse) SetBody(v *StopDatasetJobResponseBody) *StopDatasetJobResponse {
	s.Body = v
	return s
}

func (s *StopDatasetJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
