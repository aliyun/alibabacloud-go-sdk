// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizedListHeadersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomizedListHeadersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomizedListHeadersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomizedListHeadersResponseBody) *ModifyCustomizedListHeadersResponse
	GetBody() *ModifyCustomizedListHeadersResponseBody
}

type ModifyCustomizedListHeadersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomizedListHeadersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomizedListHeadersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizedListHeadersResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomizedListHeadersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomizedListHeadersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomizedListHeadersResponse) GetBody() *ModifyCustomizedListHeadersResponseBody {
	return s.Body
}

func (s *ModifyCustomizedListHeadersResponse) SetHeaders(v map[string]*string) *ModifyCustomizedListHeadersResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomizedListHeadersResponse) SetStatusCode(v int32) *ModifyCustomizedListHeadersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomizedListHeadersResponse) SetBody(v *ModifyCustomizedListHeadersResponseBody) *ModifyCustomizedListHeadersResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomizedListHeadersResponse) Validate() error {
	return dara.Validate(s)
}
