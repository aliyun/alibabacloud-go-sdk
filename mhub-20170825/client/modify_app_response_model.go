// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppResponseBody) *ModifyAppResponse
	GetBody() *ModifyAppResponseBody
}

type ModifyAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppResponse) GetBody() *ModifyAppResponseBody {
	return s.Body
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

func (s *ModifyAppResponse) Validate() error {
	return dara.Validate(s)
}
