// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPluginAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PluginAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PluginAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *PluginAnalysisResponseBody) *PluginAnalysisResponse
	GetBody() *PluginAnalysisResponseBody
}

type PluginAnalysisResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PluginAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PluginAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s PluginAnalysisResponse) GoString() string {
	return s.String()
}

func (s *PluginAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PluginAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PluginAnalysisResponse) GetBody() *PluginAnalysisResponseBody {
	return s.Body
}

func (s *PluginAnalysisResponse) SetHeaders(v map[string]*string) *PluginAnalysisResponse {
	s.Headers = v
	return s
}

func (s *PluginAnalysisResponse) SetStatusCode(v int32) *PluginAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *PluginAnalysisResponse) SetBody(v *PluginAnalysisResponseBody) *PluginAnalysisResponse {
	s.Body = v
	return s
}

func (s *PluginAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
