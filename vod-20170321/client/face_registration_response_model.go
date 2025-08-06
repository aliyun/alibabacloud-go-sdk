// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceRegistrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceRegistrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceRegistrationResponse
	GetStatusCode() *int32
	SetBody(v *FaceRegistrationResponseBody) *FaceRegistrationResponse
	GetBody() *FaceRegistrationResponseBody
}

type FaceRegistrationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceRegistrationResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceRegistrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceRegistrationResponse) GetBody() *FaceRegistrationResponseBody {
	return s.Body
}

func (s *FaceRegistrationResponse) SetHeaders(v map[string]*string) *FaceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *FaceRegistrationResponse) SetStatusCode(v int32) *FaceRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceRegistrationResponse) SetBody(v *FaceRegistrationResponseBody) *FaceRegistrationResponse {
	s.Body = v
	return s
}

func (s *FaceRegistrationResponse) Validate() error {
	return dara.Validate(s)
}
