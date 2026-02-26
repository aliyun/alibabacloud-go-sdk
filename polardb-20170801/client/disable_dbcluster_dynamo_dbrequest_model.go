// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDynamoDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterDynamoDBRequest
	GetDBClusterId() *string
}

type DisableDBClusterDynamoDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DisableDBClusterDynamoDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDynamoDBRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDynamoDBRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterDynamoDBRequest) SetDBClusterId(v string) *DisableDBClusterDynamoDBRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterDynamoDBRequest) Validate() error {
	return dara.Validate(s)
}
