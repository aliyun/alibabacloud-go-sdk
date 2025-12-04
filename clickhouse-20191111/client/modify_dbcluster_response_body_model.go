// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBCluster(v *ModifyDBClusterResponseBodyDBCluster) *ModifyDBClusterResponseBody
	GetDBCluster() *ModifyDBClusterResponseBodyDBCluster
	SetRequestId(v string) *ModifyDBClusterResponseBody
	GetRequestId() *string
}

type ModifyDBClusterResponseBody struct {
	// The clusters.
	DBCluster *ModifyDBClusterResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BA30A000-3A4D-5B97-9420-E5D0D49F7016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) GetDBCluster() *ModifyDBClusterResponseBodyDBCluster {
	return s.DBCluster
}

func (s *ModifyDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterResponseBody) SetDBCluster(v *ModifyDBClusterResponseBodyDBCluster) *ModifyDBClusterResponseBody {
	s.DBCluster = v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) Validate() error {
	if s.DBCluster != nil {
		if err := s.DBCluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBClusterResponseBodyDBCluster struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp19lo45sy98x****
	DbClusterId *string `json:"dbClusterId,omitempty" xml:"dbClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21417210003****
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ModifyDBClusterResponseBodyDBCluster) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResponseBodyDBCluster) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBodyDBCluster) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyDBClusterResponseBodyDBCluster) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterResponseBodyDBCluster) SetDbClusterId(v string) *ModifyDBClusterResponseBodyDBCluster {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyDBCluster) SetOrderId(v string) *ModifyDBClusterResponseBodyDBCluster {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyDBCluster) Validate() error {
	return dara.Validate(s)
}
