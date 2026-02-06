// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAnnotationMissionSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAnnotationMissionSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetAnnotationMissionSummaryResponseBody) *GetAnnotationMissionSummaryResponse
	GetBody() *GetAnnotationMissionSummaryResponseBody
}

type GetAnnotationMissionSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnnotationMissionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnnotationMissionSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAnnotationMissionSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAnnotationMissionSummaryResponse) GetBody() *GetAnnotationMissionSummaryResponseBody {
	return s.Body
}

func (s *GetAnnotationMissionSummaryResponse) SetHeaders(v map[string]*string) *GetAnnotationMissionSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAnnotationMissionSummaryResponse) SetStatusCode(v int32) *GetAnnotationMissionSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnnotationMissionSummaryResponse) SetBody(v *GetAnnotationMissionSummaryResponseBody) *GetAnnotationMissionSummaryResponse {
	s.Body = v
	return s
}

func (s *GetAnnotationMissionSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
