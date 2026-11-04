// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceUserResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceUserResponseBody) *AddFaceUserResponse
	GetBody() *AddFaceUserResponseBody
}

type AddFaceUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceUserResponse) GetBody() *AddFaceUserResponseBody {
	return s.Body
}

func (s *AddFaceUserResponse) SetHeaders(v map[string]*string) *AddFaceUserResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserResponse) SetStatusCode(v int32) *AddFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserResponse) SetBody(v *AddFaceUserResponseBody) *AddFaceUserResponse {
	s.Body = v
	return s
}

func (s *AddFaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
