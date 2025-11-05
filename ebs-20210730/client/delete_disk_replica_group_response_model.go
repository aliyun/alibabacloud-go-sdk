// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiskReplicaGroupResponseBody) *DeleteDiskReplicaGroupResponse
	GetBody() *DeleteDiskReplicaGroupResponseBody
}

type DeleteDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiskReplicaGroupResponse) GetBody() *DeleteDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *DeleteDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetStatusCode(v int32) *DeleteDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetBody(v *DeleteDiskReplicaGroupResponseBody) *DeleteDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
