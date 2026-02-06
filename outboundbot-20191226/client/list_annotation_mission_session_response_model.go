// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnnotationMissionSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnnotationMissionSessionResponse
	GetStatusCode() *int32
	SetBody(v *ListAnnotationMissionSessionResponseBody) *ListAnnotationMissionSessionResponse
	GetBody() *ListAnnotationMissionSessionResponseBody
}

type ListAnnotationMissionSessionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnnotationMissionSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnnotationMissionSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionResponse) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnnotationMissionSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnnotationMissionSessionResponse) GetBody() *ListAnnotationMissionSessionResponseBody {
	return s.Body
}

func (s *ListAnnotationMissionSessionResponse) SetHeaders(v map[string]*string) *ListAnnotationMissionSessionResponse {
	s.Headers = v
	return s
}

func (s *ListAnnotationMissionSessionResponse) SetStatusCode(v int32) *ListAnnotationMissionSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnnotationMissionSessionResponse) SetBody(v *ListAnnotationMissionSessionResponseBody) *ListAnnotationMissionSessionResponse {
	s.Body = v
	return s
}

func (s *ListAnnotationMissionSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
