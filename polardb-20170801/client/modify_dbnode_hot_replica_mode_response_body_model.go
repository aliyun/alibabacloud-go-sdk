// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeHotReplicaModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeHotReplicaModeResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBNodeHotReplicaModeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBNodeHotReplicaModeResponseBody
	GetRequestId() *string
}

type ModifyDBNodeHotReplicaModeResponseBody struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// example:
	//
	// pc-2vc327c2a14a3u858
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2148126708*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9F4DDFBF-205C-5A3E-8DC2-C60F3E3FDC38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeHotReplicaModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeHotReplicaModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) SetDBClusterId(v string) *ModifyDBNodeHotReplicaModeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) SetOrderId(v string) *ModifyDBNodeHotReplicaModeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) SetRequestId(v string) *ModifyDBNodeHotReplicaModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeResponseBody) Validate() error {
	return dara.Validate(s)
}
