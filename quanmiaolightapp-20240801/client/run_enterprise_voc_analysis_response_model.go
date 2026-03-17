// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEnterpriseVocAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunEnterpriseVocAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunEnterpriseVocAnalysisResponse
	GetStatusCode() *int32
	SetId(v string) *RunEnterpriseVocAnalysisResponse
	GetId() *string
	SetEvent(v string) *RunEnterpriseVocAnalysisResponse
	GetEvent() *string
	SetBody(v *RunEnterpriseVocAnalysisResponseBody) *RunEnterpriseVocAnalysisResponse
	GetBody() *RunEnterpriseVocAnalysisResponseBody
}

type RunEnterpriseVocAnalysisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                               `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                               `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunEnterpriseVocAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunEnterpriseVocAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunEnterpriseVocAnalysisResponse) GetId() *string {
	return s.Id
}

func (s *RunEnterpriseVocAnalysisResponse) GetEvent() *string {
	return s.Event
}

func (s *RunEnterpriseVocAnalysisResponse) GetBody() *RunEnterpriseVocAnalysisResponseBody {
	return s.Body
}

func (s *RunEnterpriseVocAnalysisResponse) SetHeaders(v map[string]*string) *RunEnterpriseVocAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetStatusCode(v int32) *RunEnterpriseVocAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetId(v string) *RunEnterpriseVocAnalysisResponse {
	s.Id = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetEvent(v string) *RunEnterpriseVocAnalysisResponse {
	s.Event = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) SetBody(v *RunEnterpriseVocAnalysisResponseBody) *RunEnterpriseVocAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
