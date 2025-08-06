// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetL2OssKeyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetL2OssKeyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetL2OssKeyConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetL2OssKeyConfigResponseBody) *SetL2OssKeyConfigResponse
	GetBody() *SetL2OssKeyConfigResponseBody
}

type SetL2OssKeyConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetL2OssKeyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetL2OssKeyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetL2OssKeyConfigResponse) GoString() string {
	return s.String()
}

func (s *SetL2OssKeyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetL2OssKeyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetL2OssKeyConfigResponse) GetBody() *SetL2OssKeyConfigResponseBody {
	return s.Body
}

func (s *SetL2OssKeyConfigResponse) SetHeaders(v map[string]*string) *SetL2OssKeyConfigResponse {
	s.Headers = v
	return s
}

func (s *SetL2OssKeyConfigResponse) SetStatusCode(v int32) *SetL2OssKeyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetL2OssKeyConfigResponse) SetBody(v *SetL2OssKeyConfigResponseBody) *SetL2OssKeyConfigResponse {
	s.Body = v
	return s
}

func (s *SetL2OssKeyConfigResponse) Validate() error {
	return dara.Validate(s)
}
