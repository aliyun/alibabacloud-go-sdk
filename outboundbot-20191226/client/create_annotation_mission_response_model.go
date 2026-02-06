// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnotationMissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnnotationMissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnnotationMissionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnnotationMissionResponseBody) *CreateAnnotationMissionResponse
	GetBody() *CreateAnnotationMissionResponseBody
}

type CreateAnnotationMissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnnotationMissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnnotationMissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnotationMissionResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnotationMissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnnotationMissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnnotationMissionResponse) GetBody() *CreateAnnotationMissionResponseBody {
	return s.Body
}

func (s *CreateAnnotationMissionResponse) SetHeaders(v map[string]*string) *CreateAnnotationMissionResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnotationMissionResponse) SetStatusCode(v int32) *CreateAnnotationMissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnnotationMissionResponse) SetBody(v *CreateAnnotationMissionResponseBody) *CreateAnnotationMissionResponse {
	s.Body = v
	return s
}

func (s *CreateAnnotationMissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
