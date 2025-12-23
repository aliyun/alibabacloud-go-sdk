// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManualTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManualTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListManualTaskInstancesResponseBody) *ListManualTaskInstancesResponse
	GetBody() *ListManualTaskInstancesResponseBody
}

type ListManualTaskInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManualTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManualTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManualTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManualTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManualTaskInstancesResponse) GetBody() *ListManualTaskInstancesResponseBody {
	return s.Body
}

func (s *ListManualTaskInstancesResponse) SetHeaders(v map[string]*string) *ListManualTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListManualTaskInstancesResponse) SetStatusCode(v int32) *ListManualTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManualTaskInstancesResponse) SetBody(v *ListManualTaskInstancesResponseBody) *ListManualTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListManualTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
