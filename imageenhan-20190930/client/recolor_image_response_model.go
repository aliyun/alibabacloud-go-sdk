// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecolorImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecolorImageResponse
	GetStatusCode() *int32
	SetBody(v *RecolorImageResponseBody) *RecolorImageResponse
	GetBody() *RecolorImageResponseBody
}

type RecolorImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecolorImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecolorImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageResponse) GoString() string {
	return s.String()
}

func (s *RecolorImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecolorImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecolorImageResponse) GetBody() *RecolorImageResponseBody {
	return s.Body
}

func (s *RecolorImageResponse) SetHeaders(v map[string]*string) *RecolorImageResponse {
	s.Headers = v
	return s
}

func (s *RecolorImageResponse) SetStatusCode(v int32) *RecolorImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecolorImageResponse) SetBody(v *RecolorImageResponseBody) *RecolorImageResponse {
	s.Body = v
	return s
}

func (s *RecolorImageResponse) Validate() error {
	return dara.Validate(s)
}
