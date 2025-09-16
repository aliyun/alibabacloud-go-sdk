// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIndexVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIndexVersionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIndexVersionResponseBody) *ModifyIndexVersionResponse
	GetBody() *ModifyIndexVersionResponseBody
}

type ModifyIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIndexVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIndexVersionResponse) GetBody() *ModifyIndexVersionResponseBody {
	return s.Body
}

func (s *ModifyIndexVersionResponse) SetHeaders(v map[string]*string) *ModifyIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexVersionResponse) SetStatusCode(v int32) *ModifyIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexVersionResponse) SetBody(v *ModifyIndexVersionResponseBody) *ModifyIndexVersionResponse {
	s.Body = v
	return s
}

func (s *ModifyIndexVersionResponse) Validate() error {
	return dara.Validate(s)
}
