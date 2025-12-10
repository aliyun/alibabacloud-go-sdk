// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAINodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAINodesResponseBody
	GetDBClusterId() *string
	SetDBNodeIds(v []*string) *CreateAINodesResponseBody
	GetDBNodeIds() []*string
	SetOrderId(v string) *CreateAINodesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAINodesResponseBody
	GetRequestId() *string
}

type CreateAINodesResponseBody struct {
	// example:
	//
	// pm-xxxxxx
	DBClusterId *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeIds   []*string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2148126708*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAINodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAINodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAINodesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAINodesResponseBody) GetDBNodeIds() []*string {
	return s.DBNodeIds
}

func (s *CreateAINodesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAINodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAINodesResponseBody) SetDBClusterId(v string) *CreateAINodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAINodesResponseBody) SetDBNodeIds(v []*string) *CreateAINodesResponseBody {
	s.DBNodeIds = v
	return s
}

func (s *CreateAINodesResponseBody) SetOrderId(v string) *CreateAINodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAINodesResponseBody) SetRequestId(v string) *CreateAINodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAINodesResponseBody) Validate() error {
	return dara.Validate(s)
}
