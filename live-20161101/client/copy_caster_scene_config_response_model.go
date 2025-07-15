// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyCasterSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyCasterSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyCasterSceneConfigResponseBody) *CopyCasterSceneConfigResponse
	GetBody() *CopyCasterSceneConfigResponseBody
}

type CopyCasterSceneConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyCasterSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyCasterSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyCasterSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyCasterSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyCasterSceneConfigResponse) GetBody() *CopyCasterSceneConfigResponseBody {
	return s.Body
}

func (s *CopyCasterSceneConfigResponse) SetHeaders(v map[string]*string) *CopyCasterSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyCasterSceneConfigResponse) SetStatusCode(v int32) *CopyCasterSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyCasterSceneConfigResponse) SetBody(v *CopyCasterSceneConfigResponseBody) *CopyCasterSceneConfigResponse {
	s.Body = v
	return s
}

func (s *CopyCasterSceneConfigResponse) Validate() error {
	return dara.Validate(s)
}
