// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartWorkflowInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartWorkflowInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StartWorkflowInstancesResponseBody) *StartWorkflowInstancesResponse
	GetBody() *StartWorkflowInstancesResponseBody
}

type StartWorkflowInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartWorkflowInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *StartWorkflowInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartWorkflowInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartWorkflowInstancesResponse) GetBody() *StartWorkflowInstancesResponseBody {
	return s.Body
}

func (s *StartWorkflowInstancesResponse) SetHeaders(v map[string]*string) *StartWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *StartWorkflowInstancesResponse) SetStatusCode(v int32) *StartWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartWorkflowInstancesResponse) SetBody(v *StartWorkflowInstancesResponseBody) *StartWorkflowInstancesResponse {
	s.Body = v
	return s
}

func (s *StartWorkflowInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
