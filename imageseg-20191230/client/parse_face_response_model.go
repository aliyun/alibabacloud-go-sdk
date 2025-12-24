// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ParseFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ParseFaceResponse
	GetStatusCode() *int32
	SetBody(v *ParseFaceResponseBody) *ParseFaceResponse
	GetBody() *ParseFaceResponseBody
}

type ParseFaceResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ParseFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ParseFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ParseFaceResponse) GoString() string {
	return s.String()
}

func (s *ParseFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ParseFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ParseFaceResponse) GetBody() *ParseFaceResponseBody {
	return s.Body
}

func (s *ParseFaceResponse) SetHeaders(v map[string]*string) *ParseFaceResponse {
	s.Headers = v
	return s
}

func (s *ParseFaceResponse) SetStatusCode(v int32) *ParseFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseFaceResponse) SetBody(v *ParseFaceResponseBody) *ParseFaceResponse {
	s.Body = v
	return s
}

func (s *ParseFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
