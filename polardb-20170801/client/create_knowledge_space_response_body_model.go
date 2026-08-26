// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateKnowledgeSpaceResponseBody
	GetDBClusterId() *string
	SetKnowledgeSpaceId(v string) *CreateKnowledgeSpaceResponseBody
	GetKnowledgeSpaceId() *string
	SetOrderId(v string) *CreateKnowledgeSpaceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateKnowledgeSpaceResponseBody
	GetRequestId() *string
}

type CreateKnowledgeSpaceResponseBody struct {
	// The ID of the PolarDB instance created by automatic creation.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The unique identifier of the knowledge space.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20951253014****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F029645-FED9-4FE8-A6D3-488954******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKnowledgeSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSpaceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateKnowledgeSpaceResponseBody) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *CreateKnowledgeSpaceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateKnowledgeSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeSpaceResponseBody) SetDBClusterId(v string) *CreateKnowledgeSpaceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateKnowledgeSpaceResponseBody) SetKnowledgeSpaceId(v string) *CreateKnowledgeSpaceResponseBody {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *CreateKnowledgeSpaceResponseBody) SetOrderId(v string) *CreateKnowledgeSpaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateKnowledgeSpaceResponseBody) SetRequestId(v string) *CreateKnowledgeSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
