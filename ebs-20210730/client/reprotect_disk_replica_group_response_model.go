// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReprotectDiskReplicaGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReprotectDiskReplicaGroupResponse
	GetStatusCode() *int32
	SetBody(v *ReprotectDiskReplicaGroupResponseBody) *ReprotectDiskReplicaGroupResponse
	GetBody() *ReprotectDiskReplicaGroupResponseBody
}

type ReprotectDiskReplicaGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReprotectDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReprotectDiskReplicaGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReprotectDiskReplicaGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReprotectDiskReplicaGroupResponse) GetBody() *ReprotectDiskReplicaGroupResponseBody {
	return s.Body
}

func (s *ReprotectDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetStatusCode(v int32) *ReprotectDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetBody(v *ReprotectDiskReplicaGroupResponseBody) *ReprotectDiskReplicaGroupResponse {
	s.Body = v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
