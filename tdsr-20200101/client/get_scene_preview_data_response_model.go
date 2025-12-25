// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScenePreviewDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScenePreviewDataResponse
	GetStatusCode() *int32
	SetBody(v *GetScenePreviewDataResponseBody) *GetScenePreviewDataResponse
	GetBody() *GetScenePreviewDataResponseBody
}

type GetScenePreviewDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScenePreviewDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScenePreviewDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScenePreviewDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScenePreviewDataResponse) GetBody() *GetScenePreviewDataResponseBody {
	return s.Body
}

func (s *GetScenePreviewDataResponse) SetHeaders(v map[string]*string) *GetScenePreviewDataResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewDataResponse) SetStatusCode(v int32) *GetScenePreviewDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewDataResponse) SetBody(v *GetScenePreviewDataResponseBody) *GetScenePreviewDataResponse {
	s.Body = v
	return s
}

func (s *GetScenePreviewDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
