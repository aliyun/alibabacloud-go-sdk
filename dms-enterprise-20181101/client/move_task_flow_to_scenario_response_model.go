// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveTaskFlowToScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveTaskFlowToScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveTaskFlowToScenarioResponse
	GetStatusCode() *int32
	SetBody(v *MoveTaskFlowToScenarioResponseBody) *MoveTaskFlowToScenarioResponse
	GetBody() *MoveTaskFlowToScenarioResponseBody
}

type MoveTaskFlowToScenarioResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveTaskFlowToScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveTaskFlowToScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveTaskFlowToScenarioResponse) GoString() string {
	return s.String()
}

func (s *MoveTaskFlowToScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveTaskFlowToScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveTaskFlowToScenarioResponse) GetBody() *MoveTaskFlowToScenarioResponseBody {
	return s.Body
}

func (s *MoveTaskFlowToScenarioResponse) SetHeaders(v map[string]*string) *MoveTaskFlowToScenarioResponse {
	s.Headers = v
	return s
}

func (s *MoveTaskFlowToScenarioResponse) SetStatusCode(v int32) *MoveTaskFlowToScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveTaskFlowToScenarioResponse) SetBody(v *MoveTaskFlowToScenarioResponseBody) *MoveTaskFlowToScenarioResponse {
	s.Body = v
	return s
}

func (s *MoveTaskFlowToScenarioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
