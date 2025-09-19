// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoDelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoDelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoDelConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoDelConfigResponseBody) *ModifyAutoDelConfigResponse
	GetBody() *ModifyAutoDelConfigResponseBody
}

type ModifyAutoDelConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoDelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoDelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoDelConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoDelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoDelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoDelConfigResponse) GetBody() *ModifyAutoDelConfigResponseBody {
	return s.Body
}

func (s *ModifyAutoDelConfigResponse) SetHeaders(v map[string]*string) *ModifyAutoDelConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoDelConfigResponse) SetStatusCode(v int32) *ModifyAutoDelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoDelConfigResponse) SetBody(v *ModifyAutoDelConfigResponseBody) *ModifyAutoDelConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoDelConfigResponse) Validate() error {
	return dara.Validate(s)
}
