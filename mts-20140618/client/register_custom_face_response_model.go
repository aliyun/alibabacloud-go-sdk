// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterCustomFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterCustomFaceResponse
	GetStatusCode() *int32
	SetBody(v *RegisterCustomFaceResponseBody) *RegisterCustomFaceResponse
	GetBody() *RegisterCustomFaceResponseBody
}

type RegisterCustomFaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCustomFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCustomFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomFaceResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterCustomFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterCustomFaceResponse) GetBody() *RegisterCustomFaceResponseBody {
	return s.Body
}

func (s *RegisterCustomFaceResponse) SetHeaders(v map[string]*string) *RegisterCustomFaceResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomFaceResponse) SetStatusCode(v int32) *RegisterCustomFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCustomFaceResponse) SetBody(v *RegisterCustomFaceResponseBody) *RegisterCustomFaceResponse {
	s.Body = v
	return s
}

func (s *RegisterCustomFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
