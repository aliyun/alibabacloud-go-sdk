// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelIProductionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelIProductionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelIProductionJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelIProductionJobResponseBody) *CancelIProductionJobResponse
	GetBody() *CancelIProductionJobResponseBody
}

type CancelIProductionJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelIProductionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelIProductionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobResponse) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelIProductionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelIProductionJobResponse) GetBody() *CancelIProductionJobResponseBody {
	return s.Body
}

func (s *CancelIProductionJobResponse) SetHeaders(v map[string]*string) *CancelIProductionJobResponse {
	s.Headers = v
	return s
}

func (s *CancelIProductionJobResponse) SetStatusCode(v int32) *CancelIProductionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelIProductionJobResponse) SetBody(v *CancelIProductionJobResponseBody) *CancelIProductionJobResponse {
	s.Body = v
	return s
}

func (s *CancelIProductionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
