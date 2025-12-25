// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScenePreviewResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScenePreviewResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetScenePreviewResourceResponseBody) *GetScenePreviewResourceResponse
	GetBody() *GetScenePreviewResourceResponseBody
}

type GetScenePreviewResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScenePreviewResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScenePreviewResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScenePreviewResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScenePreviewResourceResponse) GetBody() *GetScenePreviewResourceResponseBody {
	return s.Body
}

func (s *GetScenePreviewResourceResponse) SetHeaders(v map[string]*string) *GetScenePreviewResourceResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewResourceResponse) SetStatusCode(v int32) *GetScenePreviewResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewResourceResponse) SetBody(v *GetScenePreviewResourceResponseBody) *GetScenePreviewResourceResponse {
	s.Body = v
	return s
}

func (s *GetScenePreviewResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
