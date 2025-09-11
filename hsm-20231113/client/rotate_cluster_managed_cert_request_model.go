// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateClusterManagedCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RotateClusterManagedCertRequest
	GetClusterId() *string
}

type RotateClusterManagedCertRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-hfau****gkaud
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s RotateClusterManagedCertRequest) String() string {
	return dara.Prettify(s)
}

func (s RotateClusterManagedCertRequest) GoString() string {
	return s.String()
}

func (s *RotateClusterManagedCertRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RotateClusterManagedCertRequest) SetClusterId(v string) *RotateClusterManagedCertRequest {
	s.ClusterId = &v
	return s
}

func (s *RotateClusterManagedCertRequest) Validate() error {
	return dara.Validate(s)
}
