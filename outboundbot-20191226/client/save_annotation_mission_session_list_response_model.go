// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionSessionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAnnotationMissionSessionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAnnotationMissionSessionListResponse
	GetStatusCode() *int32
	SetBody(v *SaveAnnotationMissionSessionListResponseBody) *SaveAnnotationMissionSessionListResponse
	GetBody() *SaveAnnotationMissionSessionListResponseBody
}

type SaveAnnotationMissionSessionListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAnnotationMissionSessionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAnnotationMissionSessionListResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListResponse) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAnnotationMissionSessionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAnnotationMissionSessionListResponse) GetBody() *SaveAnnotationMissionSessionListResponseBody {
	return s.Body
}

func (s *SaveAnnotationMissionSessionListResponse) SetHeaders(v map[string]*string) *SaveAnnotationMissionSessionListResponse {
	s.Headers = v
	return s
}

func (s *SaveAnnotationMissionSessionListResponse) SetStatusCode(v int32) *SaveAnnotationMissionSessionListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponse) SetBody(v *SaveAnnotationMissionSessionListResponseBody) *SaveAnnotationMissionSessionListResponse {
	s.Body = v
	return s
}

func (s *SaveAnnotationMissionSessionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
