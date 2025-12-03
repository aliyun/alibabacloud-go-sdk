// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMasterSlaveServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMasterSlaveServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMasterSlaveServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMasterSlaveServerGroupResponseBody) *CreateMasterSlaveServerGroupResponse
	GetBody() *CreateMasterSlaveServerGroupResponseBody
}

type CreateMasterSlaveServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMasterSlaveServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMasterSlaveServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMasterSlaveServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMasterSlaveServerGroupResponse) GetBody() *CreateMasterSlaveServerGroupResponseBody {
	return s.Body
}

func (s *CreateMasterSlaveServerGroupResponse) SetHeaders(v map[string]*string) *CreateMasterSlaveServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponse) SetStatusCode(v int32) *CreateMasterSlaveServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponse) SetBody(v *CreateMasterSlaveServerGroupResponseBody) *CreateMasterSlaveServerGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
