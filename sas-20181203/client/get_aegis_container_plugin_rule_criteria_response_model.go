// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAegisContainerPluginRuleCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAegisContainerPluginRuleCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *GetAegisContainerPluginRuleCriteriaResponseBody) *GetAegisContainerPluginRuleCriteriaResponse
	GetBody() *GetAegisContainerPluginRuleCriteriaResponseBody
}

type GetAegisContainerPluginRuleCriteriaResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAegisContainerPluginRuleCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAegisContainerPluginRuleCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleCriteriaResponse) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) GetBody() *GetAegisContainerPluginRuleCriteriaResponseBody {
	return s.Body
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) SetHeaders(v map[string]*string) *GetAegisContainerPluginRuleCriteriaResponse {
	s.Headers = v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) SetStatusCode(v int32) *GetAegisContainerPluginRuleCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) SetBody(v *GetAegisContainerPluginRuleCriteriaResponseBody) *GetAegisContainerPluginRuleCriteriaResponse {
	s.Body = v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
