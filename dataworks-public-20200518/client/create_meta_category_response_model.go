// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetaCategoryResponseBody) *CreateMetaCategoryResponse
	GetBody() *CreateMetaCategoryResponseBody
}

type CreateMetaCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetaCategoryResponse) GetBody() *CreateMetaCategoryResponseBody {
	return s.Body
}

func (s *CreateMetaCategoryResponse) SetHeaders(v map[string]*string) *CreateMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateMetaCategoryResponse) SetStatusCode(v int32) *CreateMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetaCategoryResponse) SetBody(v *CreateMetaCategoryResponseBody) *CreateMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}
