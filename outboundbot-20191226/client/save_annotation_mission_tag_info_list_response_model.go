// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionTagInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAnnotationMissionTagInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAnnotationMissionTagInfoListResponse
	GetStatusCode() *int32
	SetBody(v *SaveAnnotationMissionTagInfoListResponseBody) *SaveAnnotationMissionTagInfoListResponse
	GetBody() *SaveAnnotationMissionTagInfoListResponseBody
}

type SaveAnnotationMissionTagInfoListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAnnotationMissionTagInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAnnotationMissionTagInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionTagInfoListResponse) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionTagInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAnnotationMissionTagInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAnnotationMissionTagInfoListResponse) GetBody() *SaveAnnotationMissionTagInfoListResponseBody {
	return s.Body
}

func (s *SaveAnnotationMissionTagInfoListResponse) SetHeaders(v map[string]*string) *SaveAnnotationMissionTagInfoListResponse {
	s.Headers = v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponse) SetStatusCode(v int32) *SaveAnnotationMissionTagInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponse) SetBody(v *SaveAnnotationMissionTagInfoListResponseBody) *SaveAnnotationMissionTagInfoListResponse {
	s.Body = v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
