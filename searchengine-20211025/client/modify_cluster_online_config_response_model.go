// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOnlineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterOnlineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterOnlineConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterOnlineConfigResponseBody) *ModifyClusterOnlineConfigResponse
	GetBody() *ModifyClusterOnlineConfigResponseBody
}

type ModifyClusterOnlineConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterOnlineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterOnlineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterOnlineConfigResponse) GetBody() *ModifyClusterOnlineConfigResponseBody {
	return s.Body
}

func (s *ModifyClusterOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetStatusCode(v int32) *ModifyClusterOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetBody(v *ModifyClusterOnlineConfigResponseBody) *ModifyClusterOnlineConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) Validate() error {
	return dara.Validate(s)
}
