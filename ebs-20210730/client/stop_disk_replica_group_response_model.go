// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *StopDiskReplicaGroupResponseBody) *StopDiskReplicaGroupResponse
	GetBody() *StopDiskReplicaGroupResponseBody
}

type StopDiskReplicaGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDiskReplicaGroupResponse) GetBody() *StopDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *StopDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StopDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetStatusCode(v int32) *StopDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetBody(v *StopDiskReplicaGroupResponseBody) *StopDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *StopDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
