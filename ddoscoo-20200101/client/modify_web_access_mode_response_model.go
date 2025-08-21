// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAccessModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebAccessModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebAccessModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebAccessModeResponseBody) *ModifyWebAccessModeResponse
	GetBody() *ModifyWebAccessModeResponseBody
}

type ModifyWebAccessModeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebAccessModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAccessModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebAccessModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebAccessModeResponse) GetBody() *ModifyWebAccessModeResponseBody {
	return s.Body
}

func (s *ModifyWebAccessModeResponse) SetHeaders(v map[string]*string) *ModifyWebAccessModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAccessModeResponse) SetStatusCode(v int32) *ModifyWebAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebAccessModeResponse) SetBody(v *ModifyWebAccessModeResponseBody) *ModifyWebAccessModeResponse {
	s.Body = v
	return s
}

func (s *ModifyWebAccessModeResponse) Validate() error {
	return dara.Validate(s)
}
