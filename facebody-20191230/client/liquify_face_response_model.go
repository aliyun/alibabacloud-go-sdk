// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiquifyFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LiquifyFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LiquifyFaceResponse
	GetStatusCode() *int32
	SetBody(v *LiquifyFaceResponseBody) *LiquifyFaceResponse
	GetBody() *LiquifyFaceResponseBody
}

type LiquifyFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiquifyFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LiquifyFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s LiquifyFaceResponse) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LiquifyFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LiquifyFaceResponse) GetBody() *LiquifyFaceResponseBody {
	return s.Body
}

func (s *LiquifyFaceResponse) SetHeaders(v map[string]*string) *LiquifyFaceResponse {
	s.Headers = v
	return s
}

func (s *LiquifyFaceResponse) SetStatusCode(v int32) *LiquifyFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *LiquifyFaceResponse) SetBody(v *LiquifyFaceResponseBody) *LiquifyFaceResponse {
	s.Body = v
	return s
}

func (s *LiquifyFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
