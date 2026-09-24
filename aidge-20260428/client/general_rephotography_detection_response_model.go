// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralRephotographyDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GeneralRephotographyDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GeneralRephotographyDetectionResponse
	GetStatusCode() *int32
	SetBody(v *GeneralRephotographyDetectionResponseBody) *GeneralRephotographyDetectionResponse
	GetBody() *GeneralRephotographyDetectionResponseBody
}

type GeneralRephotographyDetectionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GeneralRephotographyDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GeneralRephotographyDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponse) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GeneralRephotographyDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GeneralRephotographyDetectionResponse) GetBody() *GeneralRephotographyDetectionResponseBody {
	return s.Body
}

func (s *GeneralRephotographyDetectionResponse) SetHeaders(v map[string]*string) *GeneralRephotographyDetectionResponse {
	s.Headers = v
	return s
}

func (s *GeneralRephotographyDetectionResponse) SetStatusCode(v int32) *GeneralRephotographyDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GeneralRephotographyDetectionResponse) SetBody(v *GeneralRephotographyDetectionResponseBody) *GeneralRephotographyDetectionResponse {
	s.Body = v
	return s
}

func (s *GeneralRephotographyDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
