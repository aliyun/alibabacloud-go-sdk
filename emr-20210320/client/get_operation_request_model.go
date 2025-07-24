// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetOperationRequest
	GetClusterId() *string
	SetOperationId(v string) *GetOperationRequest
	GetOperationId() *string
	SetRegionId(v string) *GetOperationRequest
	GetRegionId() *string
}

type GetOperationRequest struct {
	// The ID of the cluster that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation ID.
	//
	// References:
	//
	// 	- [CreateCluster](https://help.aliyun.com/document_detail/454393.html)
	//
	// 	- [IncreaseNodes](https://help.aliyun.com/document_detail/454397.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The district ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOperationRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetOperationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOperationRequest) SetClusterId(v string) *GetOperationRequest {
	s.ClusterId = &v
	return s
}

func (s *GetOperationRequest) SetOperationId(v string) *GetOperationRequest {
	s.OperationId = &v
	return s
}

func (s *GetOperationRequest) SetRegionId(v string) *GetOperationRequest {
	s.RegionId = &v
	return s
}

func (s *GetOperationRequest) Validate() error {
	return dara.Validate(s)
}
