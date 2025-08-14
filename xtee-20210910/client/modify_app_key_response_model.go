// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppKeyResponseBody) *ModifyAppKeyResponse
	GetBody() *ModifyAppKeyResponseBody
}

type ModifyAppKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppKeyResponse) GetBody() *ModifyAppKeyResponseBody {
	return s.Body
}

func (s *ModifyAppKeyResponse) SetHeaders(v map[string]*string) *ModifyAppKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppKeyResponse) SetStatusCode(v int32) *ModifyAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppKeyResponse) SetBody(v *ModifyAppKeyResponseBody) *ModifyAppKeyResponse {
	s.Body = v
	return s
}

func (s *ModifyAppKeyResponse) Validate() error {
	return dara.Validate(s)
}
