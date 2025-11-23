// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLhTaskFlowAndScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLhTaskFlowAndScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLhTaskFlowAndScenarioResponse
	GetStatusCode() *int32
	SetBody(v *ListLhTaskFlowAndScenarioResponseBody) *ListLhTaskFlowAndScenarioResponse
	GetBody() *ListLhTaskFlowAndScenarioResponseBody
}

type ListLhTaskFlowAndScenarioResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLhTaskFlowAndScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLhTaskFlowAndScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLhTaskFlowAndScenarioResponse) GoString() string {
	return s.String()
}

func (s *ListLhTaskFlowAndScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLhTaskFlowAndScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLhTaskFlowAndScenarioResponse) GetBody() *ListLhTaskFlowAndScenarioResponseBody {
	return s.Body
}

func (s *ListLhTaskFlowAndScenarioResponse) SetHeaders(v map[string]*string) *ListLhTaskFlowAndScenarioResponse {
	s.Headers = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponse) SetStatusCode(v int32) *ListLhTaskFlowAndScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponse) SetBody(v *ListLhTaskFlowAndScenarioResponseBody) *ListLhTaskFlowAndScenarioResponse {
	s.Body = v
	return s
}

func (s *ListLhTaskFlowAndScenarioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
