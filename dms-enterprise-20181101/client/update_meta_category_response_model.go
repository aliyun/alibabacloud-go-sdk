// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaCategoryResponseBody) *UpdateMetaCategoryResponse
	GetBody() *UpdateMetaCategoryResponseBody
}

type UpdateMetaCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaCategoryResponse) GetBody() *UpdateMetaCategoryResponseBody {
	return s.Body
}

func (s *UpdateMetaCategoryResponse) SetHeaders(v map[string]*string) *UpdateMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaCategoryResponse) SetStatusCode(v int32) *UpdateMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaCategoryResponse) SetBody(v *UpdateMetaCategoryResponseBody) *UpdateMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}
