// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSuperResolutionImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateSuperResolutionImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateSuperResolutionImageResponse
	GetStatusCode() *int32
	SetBody(v *GenerateSuperResolutionImageResponseBody) *GenerateSuperResolutionImageResponse
	GetBody() *GenerateSuperResolutionImageResponseBody
}

type GenerateSuperResolutionImageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateSuperResolutionImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateSuperResolutionImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateSuperResolutionImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateSuperResolutionImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateSuperResolutionImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateSuperResolutionImageResponse) GetBody() *GenerateSuperResolutionImageResponseBody {
	return s.Body
}

func (s *GenerateSuperResolutionImageResponse) SetHeaders(v map[string]*string) *GenerateSuperResolutionImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateSuperResolutionImageResponse) SetStatusCode(v int32) *GenerateSuperResolutionImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSuperResolutionImageResponse) SetBody(v *GenerateSuperResolutionImageResponseBody) *GenerateSuperResolutionImageResponse {
	s.Body = v
	return s
}

func (s *GenerateSuperResolutionImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
