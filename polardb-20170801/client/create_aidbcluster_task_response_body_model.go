// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAIDBClusterTaskResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *CreateAIDBClusterTaskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAIDBClusterTaskResponseBody
	GetRequestId() *string
}

type CreateAIDBClusterTaskResponseBody struct {
	// example:
	//
	// pm-2zo88***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAIDBClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterTaskResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterTaskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAIDBClusterTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIDBClusterTaskResponseBody) SetDBClusterId(v string) *CreateAIDBClusterTaskResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterTaskResponseBody) SetOrderId(v string) *CreateAIDBClusterTaskResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAIDBClusterTaskResponseBody) SetRequestId(v string) *CreateAIDBClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIDBClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
