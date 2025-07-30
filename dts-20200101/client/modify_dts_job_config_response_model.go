// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobConfigResponseBody) *ModifyDtsJobConfigResponse
	GetBody() *ModifyDtsJobConfigResponseBody
}

type ModifyDtsJobConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobConfigResponse) GetBody() *ModifyDtsJobConfigResponseBody {
	return s.Body
}

func (s *ModifyDtsJobConfigResponse) SetHeaders(v map[string]*string) *ModifyDtsJobConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobConfigResponse) SetStatusCode(v int32) *ModifyDtsJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobConfigResponse) SetBody(v *ModifyDtsJobConfigResponseBody) *ModifyDtsJobConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobConfigResponse) Validate() error {
	return dara.Validate(s)
}
