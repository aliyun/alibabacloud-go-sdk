// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateBroadcastNewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateBroadcastNewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateBroadcastNewsResponse
	GetStatusCode() *int32
	SetBody(v *GenerateBroadcastNewsResponseBody) *GenerateBroadcastNewsResponse
	GetBody() *GenerateBroadcastNewsResponseBody
}

type GenerateBroadcastNewsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateBroadcastNewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateBroadcastNewsResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponse) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateBroadcastNewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateBroadcastNewsResponse) GetBody() *GenerateBroadcastNewsResponseBody {
	return s.Body
}

func (s *GenerateBroadcastNewsResponse) SetHeaders(v map[string]*string) *GenerateBroadcastNewsResponse {
	s.Headers = v
	return s
}

func (s *GenerateBroadcastNewsResponse) SetStatusCode(v int32) *GenerateBroadcastNewsResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateBroadcastNewsResponse) SetBody(v *GenerateBroadcastNewsResponseBody) *GenerateBroadcastNewsResponse {
	s.Body = v
	return s
}

func (s *GenerateBroadcastNewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
