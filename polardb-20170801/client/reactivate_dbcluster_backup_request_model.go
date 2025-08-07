// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReactivateDBClusterBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ReactivateDBClusterBackupRequest
	GetDBClusterId() *string
}

type ReactivateDBClusterBackupRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ReactivateDBClusterBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s ReactivateDBClusterBackupRequest) GoString() string {
	return s.String()
}

func (s *ReactivateDBClusterBackupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ReactivateDBClusterBackupRequest) SetDBClusterId(v string) *ReactivateDBClusterBackupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReactivateDBClusterBackupRequest) Validate() error {
	return dara.Validate(s)
}
