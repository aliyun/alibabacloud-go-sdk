// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaceDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFaceDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFaceDbResponse
	GetStatusCode() *int32
	SetBody(v *CreateFaceDbResponseBody) *CreateFaceDbResponse
	GetBody() *CreateFaceDbResponseBody
}

type CreateFaceDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFaceDbResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFaceDbResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFaceDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFaceDbResponse) GetBody() *CreateFaceDbResponseBody {
	return s.Body
}

func (s *CreateFaceDbResponse) SetHeaders(v map[string]*string) *CreateFaceDbResponse {
	s.Headers = v
	return s
}

func (s *CreateFaceDbResponse) SetStatusCode(v int32) *CreateFaceDbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFaceDbResponse) SetBody(v *CreateFaceDbResponseBody) *CreateFaceDbResponse {
	s.Body = v
	return s
}

func (s *CreateFaceDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
