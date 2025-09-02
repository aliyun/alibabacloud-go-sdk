// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddToMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddToMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddToMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *AddToMetaCategoryResponseBody) *AddToMetaCategoryResponse
	GetBody() *AddToMetaCategoryResponseBody
}

type AddToMetaCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddToMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddToMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddToMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddToMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddToMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddToMetaCategoryResponse) GetBody() *AddToMetaCategoryResponseBody {
	return s.Body
}

func (s *AddToMetaCategoryResponse) SetHeaders(v map[string]*string) *AddToMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddToMetaCategoryResponse) SetStatusCode(v int32) *AddToMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddToMetaCategoryResponse) SetBody(v *AddToMetaCategoryResponseBody) *AddToMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *AddToMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}
