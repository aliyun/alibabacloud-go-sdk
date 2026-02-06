// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnnotationMissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnnotationMissionResponse
	GetStatusCode() *int32
	SetBody(v *ListAnnotationMissionResponseBody) *ListAnnotationMissionResponse
	GetBody() *ListAnnotationMissionResponseBody
}

type ListAnnotationMissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnnotationMissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnnotationMissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionResponse) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnnotationMissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnnotationMissionResponse) GetBody() *ListAnnotationMissionResponseBody {
	return s.Body
}

func (s *ListAnnotationMissionResponse) SetHeaders(v map[string]*string) *ListAnnotationMissionResponse {
	s.Headers = v
	return s
}

func (s *ListAnnotationMissionResponse) SetStatusCode(v int32) *ListAnnotationMissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnnotationMissionResponse) SetBody(v *ListAnnotationMissionResponseBody) *ListAnnotationMissionResponse {
	s.Body = v
	return s
}

func (s *ListAnnotationMissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
