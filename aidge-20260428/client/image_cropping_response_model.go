// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageCroppingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageCroppingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageCroppingResponse
	GetStatusCode() *int32
	SetBody(v *ImageCroppingResponseBody) *ImageCroppingResponse
	GetBody() *ImageCroppingResponseBody
}

type ImageCroppingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageCroppingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageCroppingResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageCroppingResponse) GoString() string {
	return s.String()
}

func (s *ImageCroppingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageCroppingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageCroppingResponse) GetBody() *ImageCroppingResponseBody {
	return s.Body
}

func (s *ImageCroppingResponse) SetHeaders(v map[string]*string) *ImageCroppingResponse {
	s.Headers = v
	return s
}

func (s *ImageCroppingResponse) SetStatusCode(v int32) *ImageCroppingResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageCroppingResponse) SetBody(v *ImageCroppingResponseBody) *ImageCroppingResponse {
	s.Body = v
	return s
}

func (s *ImageCroppingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
