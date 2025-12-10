// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentVisualizationMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentVisualizationMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentVisualizationMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentVisualizationMetaResponseBody) *GetExperimentVisualizationMetaResponse
	GetBody() *GetExperimentVisualizationMetaResponseBody
}

type GetExperimentVisualizationMetaResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentVisualizationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentVisualizationMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentVisualizationMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentVisualizationMetaResponse) GetBody() *GetExperimentVisualizationMetaResponseBody {
	return s.Body
}

func (s *GetExperimentVisualizationMetaResponse) SetHeaders(v map[string]*string) *GetExperimentVisualizationMetaResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentVisualizationMetaResponse) SetStatusCode(v int32) *GetExperimentVisualizationMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponse) SetBody(v *GetExperimentVisualizationMetaResponseBody) *GetExperimentVisualizationMetaResponse {
	s.Body = v
	return s
}

func (s *GetExperimentVisualizationMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
