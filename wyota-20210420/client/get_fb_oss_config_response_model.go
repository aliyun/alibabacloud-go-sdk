// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbOssConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFbOssConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFbOssConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetFbOssConfigResponseBody) *GetFbOssConfigResponse
	GetBody() *GetFbOssConfigResponseBody
}

type GetFbOssConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFbOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFbOssConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFbOssConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFbOssConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFbOssConfigResponse) GetBody() *GetFbOssConfigResponseBody {
	return s.Body
}

func (s *GetFbOssConfigResponse) SetHeaders(v map[string]*string) *GetFbOssConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFbOssConfigResponse) SetStatusCode(v int32) *GetFbOssConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFbOssConfigResponse) SetBody(v *GetFbOssConfigResponseBody) *GetFbOssConfigResponse {
	s.Body = v
	return s
}

func (s *GetFbOssConfigResponse) Validate() error {
	return dara.Validate(s)
}
