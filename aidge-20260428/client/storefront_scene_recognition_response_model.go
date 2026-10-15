// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStorefrontSceneRecognitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StorefrontSceneRecognitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StorefrontSceneRecognitionResponse
	GetStatusCode() *int32
	SetBody(v *StorefrontSceneRecognitionResponseBody) *StorefrontSceneRecognitionResponse
	GetBody() *StorefrontSceneRecognitionResponseBody
}

type StorefrontSceneRecognitionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StorefrontSceneRecognitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StorefrontSceneRecognitionResponse) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionResponse) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StorefrontSceneRecognitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StorefrontSceneRecognitionResponse) GetBody() *StorefrontSceneRecognitionResponseBody {
	return s.Body
}

func (s *StorefrontSceneRecognitionResponse) SetHeaders(v map[string]*string) *StorefrontSceneRecognitionResponse {
	s.Headers = v
	return s
}

func (s *StorefrontSceneRecognitionResponse) SetStatusCode(v int32) *StorefrontSceneRecognitionResponse {
	s.StatusCode = &v
	return s
}

func (s *StorefrontSceneRecognitionResponse) SetBody(v *StorefrontSceneRecognitionResponseBody) *StorefrontSceneRecognitionResponse {
	s.Body = v
	return s
}

func (s *StorefrontSceneRecognitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
