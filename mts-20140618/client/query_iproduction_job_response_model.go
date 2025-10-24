// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIProductionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryIProductionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryIProductionJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryIProductionJobResponseBody) *QueryIProductionJobResponse
	GetBody() *QueryIProductionJobResponseBody
}

type QueryIProductionJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIProductionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIProductionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponse) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryIProductionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryIProductionJobResponse) GetBody() *QueryIProductionJobResponseBody {
	return s.Body
}

func (s *QueryIProductionJobResponse) SetHeaders(v map[string]*string) *QueryIProductionJobResponse {
	s.Headers = v
	return s
}

func (s *QueryIProductionJobResponse) SetStatusCode(v int32) *QueryIProductionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIProductionJobResponse) SetBody(v *QueryIProductionJobResponseBody) *QueryIProductionJobResponse {
	s.Body = v
	return s
}

func (s *QueryIProductionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
