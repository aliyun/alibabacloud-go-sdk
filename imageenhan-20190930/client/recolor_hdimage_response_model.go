// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorHDImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecolorHDImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecolorHDImageResponse
	GetStatusCode() *int32
	SetBody(v *RecolorHDImageResponseBody) *RecolorHDImageResponse
	GetBody() *RecolorHDImageResponseBody
}

type RecolorHDImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecolorHDImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecolorHDImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageResponse) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecolorHDImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecolorHDImageResponse) GetBody() *RecolorHDImageResponseBody {
	return s.Body
}

func (s *RecolorHDImageResponse) SetHeaders(v map[string]*string) *RecolorHDImageResponse {
	s.Headers = v
	return s
}

func (s *RecolorHDImageResponse) SetStatusCode(v int32) *RecolorHDImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecolorHDImageResponse) SetBody(v *RecolorHDImageResponseBody) *RecolorHDImageResponse {
	s.Body = v
	return s
}

func (s *RecolorHDImageResponse) Validate() error {
	return dara.Validate(s)
}
