// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationWhitelistResponseBody) *ModifyApplicationWhitelistResponse
	GetBody() *ModifyApplicationWhitelistResponseBody
}

type ModifyApplicationWhitelistResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationWhitelistResponse) GetBody() *ModifyApplicationWhitelistResponseBody {
	return s.Body
}

func (s *ModifyApplicationWhitelistResponse) SetHeaders(v map[string]*string) *ModifyApplicationWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationWhitelistResponse) SetStatusCode(v int32) *ModifyApplicationWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationWhitelistResponse) SetBody(v *ModifyApplicationWhitelistResponseBody) *ModifyApplicationWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
