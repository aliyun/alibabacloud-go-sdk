// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMachineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMachineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMachineGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMachineGroupResponseBody) *DeleteMachineGroupResponse
	GetBody() *DeleteMachineGroupResponseBody
}

type DeleteMachineGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMachineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMachineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMachineGroupResponse) GetBody() *DeleteMachineGroupResponseBody {
	return s.Body
}

func (s *DeleteMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMachineGroupResponse) SetStatusCode(v int32) *DeleteMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMachineGroupResponse) SetBody(v *DeleteMachineGroupResponseBody) *DeleteMachineGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteMachineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
