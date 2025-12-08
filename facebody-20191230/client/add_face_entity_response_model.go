// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceEntityResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceEntityResponseBody) *AddFaceEntityResponse
	GetBody() *AddFaceEntityResponseBody
}

type AddFaceEntityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceEntityResponse) GetBody() *AddFaceEntityResponseBody {
	return s.Body
}

func (s *AddFaceEntityResponse) SetHeaders(v map[string]*string) *AddFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *AddFaceEntityResponse) SetStatusCode(v int32) *AddFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceEntityResponse) SetBody(v *AddFaceEntityResponseBody) *AddFaceEntityResponse {
	s.Body = v
	return s
}

func (s *AddFaceEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
