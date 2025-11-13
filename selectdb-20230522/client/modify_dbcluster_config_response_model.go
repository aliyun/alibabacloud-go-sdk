// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterConfigResponseBody) *ModifyDBClusterConfigResponse
	GetBody() *ModifyDBClusterConfigResponseBody
}

type ModifyDBClusterConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterConfigResponse) GetBody() *ModifyDBClusterConfigResponseBody {
	return s.Body
}

func (s *ModifyDBClusterConfigResponse) SetHeaders(v map[string]*string) *ModifyDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetStatusCode(v int32) *ModifyDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetBody(v *ModifyDBClusterConfigResponseBody) *ModifyDBClusterConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
