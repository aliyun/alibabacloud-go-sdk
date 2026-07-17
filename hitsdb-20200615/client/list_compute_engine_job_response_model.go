// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeEngineJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeEngineJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeEngineJobResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeEngineJobResponseBody) *ListComputeEngineJobResponse
	GetBody() *ListComputeEngineJobResponseBody
}

type ListComputeEngineJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeEngineJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeEngineJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeEngineJobResponse) GoString() string {
	return s.String()
}

func (s *ListComputeEngineJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeEngineJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeEngineJobResponse) GetBody() *ListComputeEngineJobResponseBody {
	return s.Body
}

func (s *ListComputeEngineJobResponse) SetHeaders(v map[string]*string) *ListComputeEngineJobResponse {
	s.Headers = v
	return s
}

func (s *ListComputeEngineJobResponse) SetStatusCode(v int32) *ListComputeEngineJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeEngineJobResponse) SetBody(v *ListComputeEngineJobResponseBody) *ListComputeEngineJobResponse {
	s.Body = v
	return s
}

func (s *ListComputeEngineJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
