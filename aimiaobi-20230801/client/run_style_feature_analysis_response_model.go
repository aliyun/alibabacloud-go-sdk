// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleFeatureAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunStyleFeatureAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunStyleFeatureAnalysisResponse
	GetStatusCode() *int32
	SetId(v string) *RunStyleFeatureAnalysisResponse
	GetId() *string
	SetEvent(v string) *RunStyleFeatureAnalysisResponse
	GetEvent() *string
	SetBody(v *RunStyleFeatureAnalysisResponseBody) *RunStyleFeatureAnalysisResponse
	GetBody() *RunStyleFeatureAnalysisResponseBody
}

type RunStyleFeatureAnalysisResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                              `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunStyleFeatureAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStyleFeatureAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunStyleFeatureAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunStyleFeatureAnalysisResponse) GetId() *string {
	return s.Id
}

func (s *RunStyleFeatureAnalysisResponse) GetEvent() *string {
	return s.Event
}

func (s *RunStyleFeatureAnalysisResponse) GetBody() *RunStyleFeatureAnalysisResponseBody {
	return s.Body
}

func (s *RunStyleFeatureAnalysisResponse) SetHeaders(v map[string]*string) *RunStyleFeatureAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetStatusCode(v int32) *RunStyleFeatureAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetId(v string) *RunStyleFeatureAnalysisResponse {
	s.Id = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetEvent(v string) *RunStyleFeatureAnalysisResponse {
	s.Event = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetBody(v *RunStyleFeatureAnalysisResponseBody) *RunStyleFeatureAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
