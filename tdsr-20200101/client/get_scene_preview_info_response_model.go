// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScenePreviewInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScenePreviewInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetScenePreviewInfoResponseBody) *GetScenePreviewInfoResponse
	GetBody() *GetScenePreviewInfoResponseBody
}

type GetScenePreviewInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScenePreviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScenePreviewInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewInfoResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScenePreviewInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScenePreviewInfoResponse) GetBody() *GetScenePreviewInfoResponseBody {
	return s.Body
}

func (s *GetScenePreviewInfoResponse) SetHeaders(v map[string]*string) *GetScenePreviewInfoResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewInfoResponse) SetStatusCode(v int32) *GetScenePreviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewInfoResponse) SetBody(v *GetScenePreviewInfoResponseBody) *GetScenePreviewInfoResponse {
	s.Body = v
	return s
}

func (s *GetScenePreviewInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
