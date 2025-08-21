// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHeadersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHeadersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHeadersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHeadersResponseBody) *ModifyHeadersResponse
	GetBody() *ModifyHeadersResponseBody
}

type ModifyHeadersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHeadersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHeadersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHeadersResponse) GoString() string {
	return s.String()
}

func (s *ModifyHeadersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHeadersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHeadersResponse) GetBody() *ModifyHeadersResponseBody {
	return s.Body
}

func (s *ModifyHeadersResponse) SetHeaders(v map[string]*string) *ModifyHeadersResponse {
	s.Headers = v
	return s
}

func (s *ModifyHeadersResponse) SetStatusCode(v int32) *ModifyHeadersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHeadersResponse) SetBody(v *ModifyHeadersResponseBody) *ModifyHeadersResponse {
	s.Body = v
	return s
}

func (s *ModifyHeadersResponse) Validate() error {
	return dara.Validate(s)
}
