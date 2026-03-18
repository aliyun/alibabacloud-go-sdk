// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantGradeVolumesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantGradeVolumesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantGradeVolumesResponseBody) *ListTextbookAssistantGradeVolumesResponse
	GetBody() *ListTextbookAssistantGradeVolumesResponseBody
}

type ListTextbookAssistantGradeVolumesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantGradeVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantGradeVolumesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantGradeVolumesResponse) GetBody() *ListTextbookAssistantGradeVolumesResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantGradeVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetBody(v *ListTextbookAssistantGradeVolumesResponseBody) *ListTextbookAssistantGradeVolumesResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
