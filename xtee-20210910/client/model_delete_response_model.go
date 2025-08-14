// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ModelDeleteResponseBody) *ModelDeleteResponse
	GetBody() *ModelDeleteResponseBody
}

type ModelDeleteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelDeleteResponse) GoString() string {
	return s.String()
}

func (s *ModelDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelDeleteResponse) GetBody() *ModelDeleteResponseBody {
	return s.Body
}

func (s *ModelDeleteResponse) SetHeaders(v map[string]*string) *ModelDeleteResponse {
	s.Headers = v
	return s
}

func (s *ModelDeleteResponse) SetStatusCode(v int32) *ModelDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelDeleteResponse) SetBody(v *ModelDeleteResponseBody) *ModelDeleteResponse {
	s.Body = v
	return s
}

func (s *ModelDeleteResponse) Validate() error {
	return dara.Validate(s)
}
