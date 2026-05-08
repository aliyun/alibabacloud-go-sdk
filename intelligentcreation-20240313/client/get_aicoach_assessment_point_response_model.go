// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachAssessmentPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachAssessmentPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachAssessmentPointResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachAssessmentPointResponseBody) *GetAICoachAssessmentPointResponse
	GetBody() *GetAICoachAssessmentPointResponseBody
}

type GetAICoachAssessmentPointResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachAssessmentPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachAssessmentPointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachAssessmentPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachAssessmentPointResponse) GetBody() *GetAICoachAssessmentPointResponseBody {
	return s.Body
}

func (s *GetAICoachAssessmentPointResponse) SetHeaders(v map[string]*string) *GetAICoachAssessmentPointResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachAssessmentPointResponse) SetStatusCode(v int32) *GetAICoachAssessmentPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachAssessmentPointResponse) SetBody(v *GetAICoachAssessmentPointResponseBody) *GetAICoachAssessmentPointResponse {
	s.Body = v
	return s
}

func (s *GetAICoachAssessmentPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
