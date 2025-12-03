// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMasterSlaveServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMasterSlaveServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMasterSlaveServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMasterSlaveServerGroupResponseBody) *DeleteMasterSlaveServerGroupResponse
	GetBody() *DeleteMasterSlaveServerGroupResponseBody
}

type DeleteMasterSlaveServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMasterSlaveServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMasterSlaveServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMasterSlaveServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMasterSlaveServerGroupResponse) GetBody() *DeleteMasterSlaveServerGroupResponseBody {
	return s.Body
}

func (s *DeleteMasterSlaveServerGroupResponse) SetHeaders(v map[string]*string) *DeleteMasterSlaveServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponse) SetStatusCode(v int32) *DeleteMasterSlaveServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponse) SetBody(v *DeleteMasterSlaveServerGroupResponseBody) *DeleteMasterSlaveServerGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
