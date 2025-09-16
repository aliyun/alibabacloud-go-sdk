// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOfflineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterOfflineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterOfflineConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterOfflineConfigResponseBody) *ModifyClusterOfflineConfigResponse
	GetBody() *ModifyClusterOfflineConfigResponseBody
}

type ModifyClusterOfflineConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterOfflineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterOfflineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterOfflineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterOfflineConfigResponse) GetBody() *ModifyClusterOfflineConfigResponseBody {
	return s.Body
}

func (s *ModifyClusterOfflineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOfflineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetStatusCode(v int32) *ModifyClusterOfflineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetBody(v *ModifyClusterOfflineConfigResponseBody) *ModifyClusterOfflineConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) Validate() error {
	return dara.Validate(s)
}
