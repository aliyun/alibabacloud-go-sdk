// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetOssConfigResponseBody) *GetOssConfigResponse
	GetBody() *GetOssConfigResponseBody
}

type GetOssConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssConfigResponse) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssConfigResponse) GetBody() *GetOssConfigResponseBody {
	return s.Body
}

func (s *GetOssConfigResponse) SetHeaders(v map[string]*string) *GetOssConfigResponse {
	s.Headers = v
	return s
}

func (s *GetOssConfigResponse) SetStatusCode(v int32) *GetOssConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssConfigResponse) SetBody(v *GetOssConfigResponseBody) *GetOssConfigResponse {
	s.Body = v
	return s
}

func (s *GetOssConfigResponse) Validate() error {
	return dara.Validate(s)
}
