// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStorageSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterStorageSpaceResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBClusterStorageSpaceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBClusterStorageSpaceResponseBody
	GetRequestId() *string
}

type ModifyDBClusterStorageSpaceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 685F028C-4FCD-407D-A559-072D63******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterStorageSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStorageSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStorageSpaceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterStorageSpaceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterStorageSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterStorageSpaceResponseBody) SetDBClusterId(v string) *ModifyDBClusterStorageSpaceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponseBody) SetOrderId(v string) *ModifyDBClusterStorageSpaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponseBody) SetRequestId(v string) *ModifyDBClusterStorageSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
