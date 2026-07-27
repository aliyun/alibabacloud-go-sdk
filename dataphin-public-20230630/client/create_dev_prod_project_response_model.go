// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDevProdProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDevProdProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDevProdProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateDevProdProjectResponseBody) *CreateDevProdProjectResponse
	GetBody() *CreateDevProdProjectResponseBody
}

type CreateDevProdProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDevProdProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDevProdProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDevProdProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDevProdProjectResponse) GetBody() *CreateDevProdProjectResponseBody {
	return s.Body
}

func (s *CreateDevProdProjectResponse) SetHeaders(v map[string]*string) *CreateDevProdProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateDevProdProjectResponse) SetStatusCode(v int32) *CreateDevProdProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDevProdProjectResponse) SetBody(v *CreateDevProdProjectResponseBody) *CreateDevProdProjectResponse {
	s.Body = v
	return s
}

func (s *CreateDevProdProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
