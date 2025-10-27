// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDynamicDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDynamicDictResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDynamicDictResponseBody) *GenerateDynamicDictResponse
	GetBody() *GenerateDynamicDictResponseBody
}

type GenerateDynamicDictResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDynamicDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDynamicDictResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicDictResponse) GoString() string {
	return s.String()
}

func (s *GenerateDynamicDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDynamicDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDynamicDictResponse) GetBody() *GenerateDynamicDictResponseBody {
	return s.Body
}

func (s *GenerateDynamicDictResponse) SetHeaders(v map[string]*string) *GenerateDynamicDictResponse {
	s.Headers = v
	return s
}

func (s *GenerateDynamicDictResponse) SetStatusCode(v int32) *GenerateDynamicDictResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDynamicDictResponse) SetBody(v *GenerateDynamicDictResponseBody) *GenerateDynamicDictResponse {
	s.Body = v
	return s
}

func (s *GenerateDynamicDictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
