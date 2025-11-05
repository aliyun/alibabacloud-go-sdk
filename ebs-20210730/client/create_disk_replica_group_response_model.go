// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiskReplicaGroupResponseBody) *CreateDiskReplicaGroupResponse
	GetBody() *CreateDiskReplicaGroupResponseBody
}

type CreateDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiskReplicaGroupResponse) GetBody() *CreateDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *CreateDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetStatusCode(v int32) *CreateDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetBody(v *CreateDiskReplicaGroupResponseBody) *CreateDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
