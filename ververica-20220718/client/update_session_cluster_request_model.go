// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SessionCluster) *UpdateSessionClusterRequest
	GetBody() *SessionCluster
}

type UpdateSessionClusterRequest struct {
	Body *SessionCluster `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterRequest) GetBody() *SessionCluster {
	return s.Body
}

func (s *UpdateSessionClusterRequest) SetBody(v *SessionCluster) *UpdateSessionClusterRequest {
	s.Body = v
	return s
}

func (s *UpdateSessionClusterRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
