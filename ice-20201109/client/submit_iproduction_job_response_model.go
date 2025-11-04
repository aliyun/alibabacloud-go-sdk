// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIProductionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIProductionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIProductionJobResponseBody) *SubmitIProductionJobResponse
	GetBody() *SubmitIProductionJobResponseBody
}

type SubmitIProductionJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIProductionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIProductionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIProductionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIProductionJobResponse) GetBody() *SubmitIProductionJobResponseBody {
	return s.Body
}

func (s *SubmitIProductionJobResponse) SetHeaders(v map[string]*string) *SubmitIProductionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIProductionJobResponse) SetStatusCode(v int32) *SubmitIProductionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIProductionJobResponse) SetBody(v *SubmitIProductionJobResponseBody) *SubmitIProductionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitIProductionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
