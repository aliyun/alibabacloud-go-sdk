// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterServiceConfigResponseBody) *ModifyClusterServiceConfigResponse
	GetBody() *ModifyClusterServiceConfigResponseBody
}

type ModifyClusterServiceConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterServiceConfigResponse) GetBody() *ModifyClusterServiceConfigResponseBody {
	return s.Body
}

func (s *ModifyClusterServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterServiceConfigResponse) SetStatusCode(v int32) *ModifyClusterServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterServiceConfigResponse) SetBody(v *ModifyClusterServiceConfigResponseBody) *ModifyClusterServiceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterServiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
