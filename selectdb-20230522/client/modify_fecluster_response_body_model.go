// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFEClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyFEClusterResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *ModifyFEClusterResponseBody
	GetRequestId() *string
}

type ModifyFEClusterResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0BF61F90-ACED-55DF-A6FE-56714B2BFCF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFEClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFEClusterResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyFEClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFEClusterResponseBody) SetDBInstanceId(v string) *ModifyFEClusterResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyFEClusterResponseBody) SetRequestId(v string) *ModifyFEClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFEClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
