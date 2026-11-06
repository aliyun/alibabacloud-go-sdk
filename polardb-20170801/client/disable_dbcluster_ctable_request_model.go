// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterCTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterCTableRequest
	GetDBClusterId() *string
}

type DisableDBClusterCTableRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DisableDBClusterCTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterCTableRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterCTableRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterCTableRequest) SetDBClusterId(v string) *DisableDBClusterCTableRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterCTableRequest) Validate() error {
	return dara.Validate(s)
}
