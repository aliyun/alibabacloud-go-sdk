// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStoragePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterStoragePerformanceResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBClusterStoragePerformanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBClusterStoragePerformanceResponseBody
	GetRequestId() *string
}

type ModifyDBClusterStoragePerformanceResponseBody struct {
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterStoragePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStoragePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) SetDBClusterId(v string) *ModifyDBClusterStoragePerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) SetOrderId(v string) *ModifyDBClusterStoragePerformanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) SetRequestId(v string) *ModifyDBClusterStoragePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}
