// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCnnfStatusDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *ListClusterCnnfStatusDetailRequest
	GetClusterIds() []*string
}

type ListClusterCnnfStatusDetailRequest struct {
	// An array that consists of the ID of the cluster.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
}

func (s ListClusterCnnfStatusDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCnnfStatusDetailRequest) GoString() string {
	return s.String()
}

func (s *ListClusterCnnfStatusDetailRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *ListClusterCnnfStatusDetailRequest) SetClusterIds(v []*string) *ListClusterCnnfStatusDetailRequest {
	s.ClusterIds = v
	return s
}

func (s *ListClusterCnnfStatusDetailRequest) Validate() error {
	return dara.Validate(s)
}
