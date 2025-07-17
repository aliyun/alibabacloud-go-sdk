// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableToCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTableToCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTableToCategoryResponse
	GetStatusCode() *int32
	SetBody(v *AddTableToCategoryResponseBody) *AddTableToCategoryResponse
	GetBody() *AddTableToCategoryResponseBody
}

type AddTableToCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTableToCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTableToCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTableToCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddTableToCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTableToCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTableToCategoryResponse) GetBody() *AddTableToCategoryResponseBody {
	return s.Body
}

func (s *AddTableToCategoryResponse) SetHeaders(v map[string]*string) *AddTableToCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddTableToCategoryResponse) SetStatusCode(v int32) *AddTableToCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTableToCategoryResponse) SetBody(v *AddTableToCategoryResponseBody) *AddTableToCategoryResponse {
	s.Body = v
	return s
}

func (s *AddTableToCategoryResponse) Validate() error {
	return dara.Validate(s)
}
