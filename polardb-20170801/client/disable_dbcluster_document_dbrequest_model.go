// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDocumentDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterDocumentDBRequest
	GetDBClusterId() *string
}

type DisableDBClusterDocumentDBRequest struct {
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DisableDBClusterDocumentDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDocumentDBRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDocumentDBRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterDocumentDBRequest) SetDBClusterId(v string) *DisableDBClusterDocumentDBRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterDocumentDBRequest) Validate() error {
	return dara.Validate(s)
}
