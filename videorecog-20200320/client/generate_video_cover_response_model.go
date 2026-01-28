// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoCoverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateVideoCoverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateVideoCoverResponse
	GetStatusCode() *int32
	SetBody(v *GenerateVideoCoverResponseBody) *GenerateVideoCoverResponse
	GetBody() *GenerateVideoCoverResponseBody
}

type GenerateVideoCoverResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateVideoCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateVideoCoverResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateVideoCoverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateVideoCoverResponse) GetBody() *GenerateVideoCoverResponseBody {
	return s.Body
}

func (s *GenerateVideoCoverResponse) SetHeaders(v map[string]*string) *GenerateVideoCoverResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoCoverResponse) SetStatusCode(v int32) *GenerateVideoCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoCoverResponse) SetBody(v *GenerateVideoCoverResponseBody) *GenerateVideoCoverResponse {
	s.Body = v
	return s
}

func (s *GenerateVideoCoverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
