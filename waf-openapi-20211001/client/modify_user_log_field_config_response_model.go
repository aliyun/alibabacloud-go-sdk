// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLogFieldConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserLogFieldConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserLogFieldConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserLogFieldConfigResponseBody) *ModifyUserLogFieldConfigResponse
	GetBody() *ModifyUserLogFieldConfigResponseBody
}

type ModifyUserLogFieldConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserLogFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserLogFieldConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLogFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserLogFieldConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserLogFieldConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserLogFieldConfigResponse) GetBody() *ModifyUserLogFieldConfigResponseBody {
	return s.Body
}

func (s *ModifyUserLogFieldConfigResponse) SetHeaders(v map[string]*string) *ModifyUserLogFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserLogFieldConfigResponse) SetStatusCode(v int32) *ModifyUserLogFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserLogFieldConfigResponse) SetBody(v *ModifyUserLogFieldConfigResponseBody) *ModifyUserLogFieldConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyUserLogFieldConfigResponse) Validate() error {
	return dara.Validate(s)
}
