// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskReplicaGroupResponseBody) *ModifyDiskReplicaGroupResponse
	GetBody() *ModifyDiskReplicaGroupResponseBody
}

type ModifyDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskReplicaGroupResponse) GetBody() *ModifyDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *ModifyDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetStatusCode(v int32) *ModifyDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetBody(v *ModifyDiskReplicaGroupResponseBody) *ModifyDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
