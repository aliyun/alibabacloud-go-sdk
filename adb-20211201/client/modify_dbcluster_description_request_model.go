// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest
	GetDBClusterId() *string
}

type ModifyDBClusterDescriptionRequest struct {
	// The description of the cluster.
	//
	// 	- The description cannot start with `http://` or `https`.
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *ModifyDBClusterDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
