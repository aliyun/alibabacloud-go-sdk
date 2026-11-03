// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageUpscaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageUpscaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageUpscaleResponse
	GetStatusCode() *int32
	SetBody(v *ImageUpscaleResponseBody) *ImageUpscaleResponse
	GetBody() *ImageUpscaleResponseBody
}

type ImageUpscaleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageUpscaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageUpscaleResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleResponse) GoString() string {
	return s.String()
}

func (s *ImageUpscaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageUpscaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageUpscaleResponse) GetBody() *ImageUpscaleResponseBody {
	return s.Body
}

func (s *ImageUpscaleResponse) SetHeaders(v map[string]*string) *ImageUpscaleResponse {
	s.Headers = v
	return s
}

func (s *ImageUpscaleResponse) SetStatusCode(v int32) *ImageUpscaleResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageUpscaleResponse) SetBody(v *ImageUpscaleResponseBody) *ImageUpscaleResponse {
	s.Body = v
	return s
}

func (s *ImageUpscaleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
