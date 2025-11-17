// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllotDatasetAccelerationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllotDatasetAccelerationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllotDatasetAccelerationTaskResponse
	GetStatusCode() *int32
	SetBody(v *AllotDatasetAccelerationTaskResponseBody) *AllotDatasetAccelerationTaskResponse
	GetBody() *AllotDatasetAccelerationTaskResponseBody
}

type AllotDatasetAccelerationTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllotDatasetAccelerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllotDatasetAccelerationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AllotDatasetAccelerationTaskResponse) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllotDatasetAccelerationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllotDatasetAccelerationTaskResponse) GetBody() *AllotDatasetAccelerationTaskResponseBody {
	return s.Body
}

func (s *AllotDatasetAccelerationTaskResponse) SetHeaders(v map[string]*string) *AllotDatasetAccelerationTaskResponse {
	s.Headers = v
	return s
}

func (s *AllotDatasetAccelerationTaskResponse) SetStatusCode(v int32) *AllotDatasetAccelerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponse) SetBody(v *AllotDatasetAccelerationTaskResponseBody) *AllotDatasetAccelerationTaskResponse {
	s.Body = v
	return s
}

func (s *AllotDatasetAccelerationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
