// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSaveFlowConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterSaveFlowConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterSaveFlowConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterSaveFlowConfigResponseBody) *ModelRouterSaveFlowConfigResponse
	GetBody() *ModelRouterSaveFlowConfigResponseBody
}

type ModelRouterSaveFlowConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterSaveFlowConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterSaveFlowConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSaveFlowConfigResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterSaveFlowConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterSaveFlowConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterSaveFlowConfigResponse) GetBody() *ModelRouterSaveFlowConfigResponseBody {
	return s.Body
}

func (s *ModelRouterSaveFlowConfigResponse) SetHeaders(v map[string]*string) *ModelRouterSaveFlowConfigResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterSaveFlowConfigResponse) SetStatusCode(v int32) *ModelRouterSaveFlowConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponse) SetBody(v *ModelRouterSaveFlowConfigResponseBody) *ModelRouterSaveFlowConfigResponse {
	s.Body = v
	return s
}

func (s *ModelRouterSaveFlowConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
