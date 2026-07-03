// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentAccuracyTestInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentAccuracyTestInstancesResponseBody) *ListDataAgentAccuracyTestInstancesResponse
	GetBody() *ListDataAgentAccuracyTestInstancesResponseBody
}

type ListDataAgentAccuracyTestInstancesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentAccuracyTestInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentAccuracyTestInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentAccuracyTestInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentAccuracyTestInstancesResponse) GetBody() *ListDataAgentAccuracyTestInstancesResponseBody {
	return s.Body
}

func (s *ListDataAgentAccuracyTestInstancesResponse) SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponse) SetStatusCode(v int32) *ListDataAgentAccuracyTestInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponse) SetBody(v *ListDataAgentAccuracyTestInstancesResponseBody) *ListDataAgentAccuracyTestInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
