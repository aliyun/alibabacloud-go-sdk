// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIndexResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIndexResponseBody) *ModifyIndexResponse
	GetBody() *ModifyIndexResponseBody
}

type ModifyIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIndexResponse) GetBody() *ModifyIndexResponseBody {
	return s.Body
}

func (s *ModifyIndexResponse) SetHeaders(v map[string]*string) *ModifyIndexResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexResponse) SetStatusCode(v int32) *ModifyIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexResponse) SetBody(v *ModifyIndexResponseBody) *ModifyIndexResponse {
	s.Body = v
	return s
}

func (s *ModifyIndexResponse) Validate() error {
	return dara.Validate(s)
}
