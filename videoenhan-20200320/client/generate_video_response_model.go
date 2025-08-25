// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateVideoResponse
	GetStatusCode() *int32
	SetBody(v *GenerateVideoResponseBody) *GenerateVideoResponse
	GetBody() *GenerateVideoResponseBody
}

type GenerateVideoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateVideoResponse) GetBody() *GenerateVideoResponseBody {
	return s.Body
}

func (s *GenerateVideoResponse) SetHeaders(v map[string]*string) *GenerateVideoResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoResponse) SetStatusCode(v int32) *GenerateVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoResponse) SetBody(v *GenerateVideoResponseBody) *GenerateVideoResponse {
	s.Body = v
	return s
}

func (s *GenerateVideoResponse) Validate() error {
	return dara.Validate(s)
}
