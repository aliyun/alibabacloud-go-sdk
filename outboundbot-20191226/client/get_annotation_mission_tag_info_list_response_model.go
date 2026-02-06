// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionTagInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAnnotationMissionTagInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAnnotationMissionTagInfoListResponse
	GetStatusCode() *int32
	SetBody(v *GetAnnotationMissionTagInfoListResponseBody) *GetAnnotationMissionTagInfoListResponse
	GetBody() *GetAnnotationMissionTagInfoListResponseBody
}

type GetAnnotationMissionTagInfoListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnnotationMissionTagInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnnotationMissionTagInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionTagInfoListResponse) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionTagInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAnnotationMissionTagInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAnnotationMissionTagInfoListResponse) GetBody() *GetAnnotationMissionTagInfoListResponseBody {
	return s.Body
}

func (s *GetAnnotationMissionTagInfoListResponse) SetHeaders(v map[string]*string) *GetAnnotationMissionTagInfoListResponse {
	s.Headers = v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponse) SetStatusCode(v int32) *GetAnnotationMissionTagInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponse) SetBody(v *GetAnnotationMissionTagInfoListResponseBody) *GetAnnotationMissionTagInfoListResponse {
	s.Body = v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
