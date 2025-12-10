// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAINodesResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *DeleteAINodesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DeleteAINodesResponseBody
	GetRequestId() *string
}

type DeleteAINodesResponseBody struct {
	// example:
	//
	// pm-xxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAINodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAINodesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAINodesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteAINodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAINodesResponseBody) SetDBClusterId(v string) *DeleteAINodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAINodesResponseBody) SetOrderId(v string) *DeleteAINodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *DeleteAINodesResponseBody) SetRequestId(v string) *DeleteAINodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAINodesResponseBody) Validate() error {
	return dara.Validate(s)
}
