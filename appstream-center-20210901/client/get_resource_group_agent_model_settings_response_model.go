// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupAgentModelSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupAgentModelSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupAgentModelSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupAgentModelSettingsResponseBody) *GetResourceGroupAgentModelSettingsResponse
	GetBody() *GetResourceGroupAgentModelSettingsResponseBody
}

type GetResourceGroupAgentModelSettingsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupAgentModelSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupAgentModelSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAgentModelSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAgentModelSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupAgentModelSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupAgentModelSettingsResponse) GetBody() *GetResourceGroupAgentModelSettingsResponseBody {
	return s.Body
}

func (s *GetResourceGroupAgentModelSettingsResponse) SetHeaders(v map[string]*string) *GetResourceGroupAgentModelSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponse) SetStatusCode(v int32) *GetResourceGroupAgentModelSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponse) SetBody(v *GetResourceGroupAgentModelSettingsResponseBody) *GetResourceGroupAgentModelSettingsResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
