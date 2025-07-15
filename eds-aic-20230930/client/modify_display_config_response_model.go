// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDisplayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDisplayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDisplayConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDisplayConfigResponseBody) *ModifyDisplayConfigResponse
	GetBody() *ModifyDisplayConfigResponseBody
}

type ModifyDisplayConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDisplayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDisplayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDisplayConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDisplayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDisplayConfigResponse) GetBody() *ModifyDisplayConfigResponseBody {
	return s.Body
}

func (s *ModifyDisplayConfigResponse) SetHeaders(v map[string]*string) *ModifyDisplayConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDisplayConfigResponse) SetStatusCode(v int32) *ModifyDisplayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDisplayConfigResponse) SetBody(v *ModifyDisplayConfigResponseBody) *ModifyDisplayConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDisplayConfigResponse) Validate() error {
	return dara.Validate(s)
}
