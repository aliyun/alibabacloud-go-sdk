// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoScalingRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoScalingRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoScalingRecordsResponseBody) *ListAutoScalingRecordsResponse
	GetBody() *ListAutoScalingRecordsResponseBody
}

type ListAutoScalingRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoScalingRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoScalingRecordsResponse) GetBody() *ListAutoScalingRecordsResponseBody {
	return s.Body
}

func (s *ListAutoScalingRecordsResponse) SetHeaders(v map[string]*string) *ListAutoScalingRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingRecordsResponse) SetStatusCode(v int32) *ListAutoScalingRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponse) SetBody(v *ListAutoScalingRecordsResponseBody) *ListAutoScalingRecordsResponse {
	s.Body = v
	return s
}

func (s *ListAutoScalingRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
