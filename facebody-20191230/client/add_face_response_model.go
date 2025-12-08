// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceResponseBody) *AddFaceResponse
	GetBody() *AddFaceResponseBody
}

type AddFaceResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceResponse) GoString() string {
	return s.String()
}

func (s *AddFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceResponse) GetBody() *AddFaceResponseBody {
	return s.Body
}

func (s *AddFaceResponse) SetHeaders(v map[string]*string) *AddFaceResponse {
	s.Headers = v
	return s
}

func (s *AddFaceResponse) SetStatusCode(v int32) *AddFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceResponse) SetBody(v *AddFaceResponseBody) *AddFaceResponse {
	s.Body = v
	return s
}

func (s *AddFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
