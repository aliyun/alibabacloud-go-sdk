// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FailoverDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FailoverDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *FailoverDiskReplicaGroupResponseBody) *FailoverDiskReplicaGroupResponse
	GetBody() *FailoverDiskReplicaGroupResponseBody
}

type FailoverDiskReplicaGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FailoverDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FailoverDiskReplicaGroupResponse) GetBody() *FailoverDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *FailoverDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetStatusCode(v int32) *FailoverDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetBody(v *FailoverDiskReplicaGroupResponseBody) *FailoverDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
