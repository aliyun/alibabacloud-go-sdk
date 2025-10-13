// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMachineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMachineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMachineGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetMachineGroupResponseBody) *GetMachineGroupResponse
	GetBody() *GetMachineGroupResponseBody
}

type GetMachineGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMachineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMachineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMachineGroupResponse) GetBody() *GetMachineGroupResponseBody {
	return s.Body
}

func (s *GetMachineGroupResponse) SetHeaders(v map[string]*string) *GetMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMachineGroupResponse) SetStatusCode(v int32) *GetMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMachineGroupResponse) SetBody(v *GetMachineGroupResponseBody) *GetMachineGroupResponse {
	s.Body = v
	return s
}

func (s *GetMachineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
