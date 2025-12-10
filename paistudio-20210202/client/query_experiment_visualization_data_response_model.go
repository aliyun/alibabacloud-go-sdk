// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExperimentVisualizationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryExperimentVisualizationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryExperimentVisualizationDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryExperimentVisualizationDataResponseBody) *QueryExperimentVisualizationDataResponse
	GetBody() *QueryExperimentVisualizationDataResponseBody
}

type QueryExperimentVisualizationDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryExperimentVisualizationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryExperimentVisualizationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponse) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryExperimentVisualizationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryExperimentVisualizationDataResponse) GetBody() *QueryExperimentVisualizationDataResponseBody {
	return s.Body
}

func (s *QueryExperimentVisualizationDataResponse) SetHeaders(v map[string]*string) *QueryExperimentVisualizationDataResponse {
	s.Headers = v
	return s
}

func (s *QueryExperimentVisualizationDataResponse) SetStatusCode(v int32) *QueryExperimentVisualizationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponse) SetBody(v *QueryExperimentVisualizationDataResponseBody) *QueryExperimentVisualizationDataResponse {
	s.Body = v
	return s
}

func (s *QueryExperimentVisualizationDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
