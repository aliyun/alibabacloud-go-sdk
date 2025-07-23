// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *SyncClusterRequest
	GetClusterId() *string
}

type SyncClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsgytet****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s SyncClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterRequest) GoString() string {
	return s.String()
}

func (s *SyncClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SyncClusterRequest) SetClusterId(v string) *SyncClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *SyncClusterRequest) Validate() error {
	return dara.Validate(s)
}
