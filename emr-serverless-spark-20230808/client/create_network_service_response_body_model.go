// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *CreateNetworkServiceResponseBody
	GetOperationId() *string
	SetRequestId(v string) *CreateNetworkServiceResponseBody
	GetRequestId() *string
}

type CreateNetworkServiceResponseBody struct {
	// Workspace Id。
	//
	// example:
	//
	// w-******
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateNetworkServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkServiceResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateNetworkServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkServiceResponseBody) SetOperationId(v string) *CreateNetworkServiceResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateNetworkServiceResponseBody) SetRequestId(v string) *CreateNetworkServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
