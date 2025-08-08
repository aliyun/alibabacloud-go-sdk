// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcubeWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcubeWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcubeWhitelistResponseBody) *UpdateMcubeWhitelistResponse
	GetBody() *UpdateMcubeWhitelistResponseBody
}

type UpdateMcubeWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcubeWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcubeWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeWhitelistResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcubeWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcubeWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcubeWhitelistResponse) GetBody() *UpdateMcubeWhitelistResponseBody {
	return s.Body
}

func (s *UpdateMcubeWhitelistResponse) SetHeaders(v map[string]*string) *UpdateMcubeWhitelistResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcubeWhitelistResponse) SetStatusCode(v int32) *UpdateMcubeWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcubeWhitelistResponse) SetBody(v *UpdateMcubeWhitelistResponseBody) *UpdateMcubeWhitelistResponse {
	s.Body = v
	return s
}

func (s *UpdateMcubeWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
