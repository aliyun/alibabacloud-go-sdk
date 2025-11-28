// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceIPArrayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstanceIPArrayRequest
	GetDBInstanceId() *string
	SetIPArrayName(v string) *DeleteDBInstanceIPArrayRequest
	GetIPArrayName() *string
}

type DeleteDBInstanceIPArrayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testarray
	IPArrayName *string `json:"IPArrayName,omitempty" xml:"IPArrayName,omitempty"`
}

func (s DeleteDBInstanceIPArrayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceIPArrayRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceIPArrayRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceIPArrayRequest) GetIPArrayName() *string {
	return s.IPArrayName
}

func (s *DeleteDBInstanceIPArrayRequest) SetDBInstanceId(v string) *DeleteDBInstanceIPArrayRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceIPArrayRequest) SetIPArrayName(v string) *DeleteDBInstanceIPArrayRequest {
	s.IPArrayName = &v
	return s
}

func (s *DeleteDBInstanceIPArrayRequest) Validate() error {
	return dara.Validate(s)
}
