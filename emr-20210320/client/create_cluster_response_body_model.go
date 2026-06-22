// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterResponseBody
	GetClusterId() *string
	SetOperationId(v string) *CreateClusterResponseBody
	GetOperationId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
}

type CreateClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac7f7***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOperationId(v string) *CreateClusterResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
