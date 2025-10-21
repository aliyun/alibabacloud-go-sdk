// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SessionCluster) *CreateSessionClusterRequest
	GetBody() *SessionCluster
}

type CreateSessionClusterRequest struct {
	Body *SessionCluster `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequest) GetBody() *SessionCluster {
	return s.Body
}

func (s *CreateSessionClusterRequest) SetBody(v *SessionCluster) *CreateSessionClusterRequest {
	s.Body = v
	return s
}

func (s *CreateSessionClusterRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
