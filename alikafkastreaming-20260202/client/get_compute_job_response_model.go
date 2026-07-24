// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeJobResponseBody) *GetComputeJobResponse
	GetBody() *GetComputeJobResponseBody
}

type GetComputeJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeJobResponse) GoString() string {
	return s.String()
}

func (s *GetComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeJobResponse) GetBody() *GetComputeJobResponseBody {
	return s.Body
}

func (s *GetComputeJobResponse) SetHeaders(v map[string]*string) *GetComputeJobResponse {
	s.Headers = v
	return s
}

func (s *GetComputeJobResponse) SetStatusCode(v int32) *GetComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeJobResponse) SetBody(v *GetComputeJobResponseBody) *GetComputeJobResponse {
	s.Body = v
	return s
}

func (s *GetComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
