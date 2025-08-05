// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultIPSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefaultIPSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefaultIPSConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefaultIPSConfigResponseBody) *ModifyDefaultIPSConfigResponse
	GetBody() *ModifyDefaultIPSConfigResponseBody
}

type ModifyDefaultIPSConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefaultIPSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefaultIPSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefaultIPSConfigResponse) GetBody() *ModifyDefaultIPSConfigResponseBody {
	return s.Body
}

func (s *ModifyDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *ModifyDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefaultIPSConfigResponse) SetStatusCode(v int32) *ModifyDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefaultIPSConfigResponse) SetBody(v *ModifyDefaultIPSConfigResponseBody) *ModifyDefaultIPSConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDefaultIPSConfigResponse) Validate() error {
	return dara.Validate(s)
}
