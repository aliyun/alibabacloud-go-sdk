// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultHttpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefaultHttpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefaultHttpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefaultHttpsResponseBody) *ModifyDefaultHttpsResponse
	GetBody() *ModifyDefaultHttpsResponseBody
}

type ModifyDefaultHttpsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefaultHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefaultHttpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultHttpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefaultHttpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefaultHttpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefaultHttpsResponse) GetBody() *ModifyDefaultHttpsResponseBody {
	return s.Body
}

func (s *ModifyDefaultHttpsResponse) SetHeaders(v map[string]*string) *ModifyDefaultHttpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefaultHttpsResponse) SetStatusCode(v int32) *ModifyDefaultHttpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefaultHttpsResponse) SetBody(v *ModifyDefaultHttpsResponseBody) *ModifyDefaultHttpsResponse {
	s.Body = v
	return s
}

func (s *ModifyDefaultHttpsResponse) Validate() error {
	return dara.Validate(s)
}
