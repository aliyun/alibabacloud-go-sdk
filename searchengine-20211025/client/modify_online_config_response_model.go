// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOnlineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOnlineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOnlineConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOnlineConfigResponseBody) *ModifyOnlineConfigResponse
	GetBody() *ModifyOnlineConfigResponseBody
}

type ModifyOnlineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOnlineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOnlineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOnlineConfigResponse) GetBody() *ModifyOnlineConfigResponseBody {
	return s.Body
}

func (s *ModifyOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyOnlineConfigResponse) SetStatusCode(v int32) *ModifyOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOnlineConfigResponse) SetBody(v *ModifyOnlineConfigResponseBody) *ModifyOnlineConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyOnlineConfigResponse) Validate() error {
	return dara.Validate(s)
}
