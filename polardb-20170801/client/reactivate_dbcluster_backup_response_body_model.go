// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReactivateDBClusterBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ReactivateDBClusterBackupResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ReactivateDBClusterBackupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ReactivateDBClusterBackupResponseBody
	GetRequestId() *string
}

type ReactivateDBClusterBackupResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReactivateDBClusterBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReactivateDBClusterBackupResponseBody) GoString() string {
	return s.String()
}

func (s *ReactivateDBClusterBackupResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ReactivateDBClusterBackupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ReactivateDBClusterBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReactivateDBClusterBackupResponseBody) SetDBClusterId(v string) *ReactivateDBClusterBackupResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ReactivateDBClusterBackupResponseBody) SetOrderId(v string) *ReactivateDBClusterBackupResponseBody {
	s.OrderId = &v
	return s
}

func (s *ReactivateDBClusterBackupResponseBody) SetRequestId(v string) *ReactivateDBClusterBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReactivateDBClusterBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
