// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageFixCycleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageFixCycleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageFixCycleConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageFixCycleConfigResponseBody) *ModifyImageFixCycleConfigResponse
	GetBody() *ModifyImageFixCycleConfigResponseBody
}

type ModifyImageFixCycleConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageFixCycleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageFixCycleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageFixCycleConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageFixCycleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageFixCycleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageFixCycleConfigResponse) GetBody() *ModifyImageFixCycleConfigResponseBody {
	return s.Body
}

func (s *ModifyImageFixCycleConfigResponse) SetHeaders(v map[string]*string) *ModifyImageFixCycleConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageFixCycleConfigResponse) SetStatusCode(v int32) *ModifyImageFixCycleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageFixCycleConfigResponse) SetBody(v *ModifyImageFixCycleConfigResponseBody) *ModifyImageFixCycleConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyImageFixCycleConfigResponse) Validate() error {
	return dara.Validate(s)
}
