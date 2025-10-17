// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMaintainTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterMaintainTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse
	GetBody() *ModifyDBClusterMaintainTimeResponseBody
}

type ModifyDBClusterMaintainTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterMaintainTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterMaintainTimeResponse) GetBody() *ModifyDBClusterMaintainTimeResponseBody {
	return s.Body
}

func (s *ModifyDBClusterMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBClusterMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
