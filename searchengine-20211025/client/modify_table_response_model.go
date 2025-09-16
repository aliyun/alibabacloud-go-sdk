// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTableResponseBody) *ModifyTableResponse
	GetBody() *ModifyTableResponseBody
}

type ModifyTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableResponse) GoString() string {
	return s.String()
}

func (s *ModifyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTableResponse) GetBody() *ModifyTableResponseBody {
	return s.Body
}

func (s *ModifyTableResponse) SetHeaders(v map[string]*string) *ModifyTableResponse {
	s.Headers = v
	return s
}

func (s *ModifyTableResponse) SetStatusCode(v int32) *ModifyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTableResponse) SetBody(v *ModifyTableResponseBody) *ModifyTableResponse {
	s.Body = v
	return s
}

func (s *ModifyTableResponse) Validate() error {
	return dara.Validate(s)
}
