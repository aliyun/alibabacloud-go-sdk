// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *StartDiskReplicaGroupResponseBody) *StartDiskReplicaGroupResponse
	GetBody() *StartDiskReplicaGroupResponseBody
}

type StartDiskReplicaGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDiskReplicaGroupResponse) GetBody() *StartDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *StartDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StartDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetStatusCode(v int32) *StartDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetBody(v *StartDiskReplicaGroupResponseBody) *StartDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *StartDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
