// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaCategoryResponseBody) *GetMetaCategoryResponse
	GetBody() *GetMetaCategoryResponseBody
}

type GetMetaCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaCategoryResponse) GetBody() *GetMetaCategoryResponseBody {
	return s.Body
}

func (s *GetMetaCategoryResponse) SetHeaders(v map[string]*string) *GetMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetMetaCategoryResponse) SetStatusCode(v int32) *GetMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaCategoryResponse) SetBody(v *GetMetaCategoryResponseBody) *GetMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *GetMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}
