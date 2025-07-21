// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowJSONAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdateFlowJSONAssetRequest
	GetCustSpaceId() *string
	SetFilePath(v string) *UpdateFlowJSONAssetRequest
	GetFilePath() *string
	SetFlowId(v string) *UpdateFlowJSONAssetRequest
	GetFlowId() *string
	SetOwnerId(v int64) *UpdateFlowJSONAssetRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateFlowJSONAssetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateFlowJSONAssetRequest
	GetResourceOwnerId() *int64
}

type UpdateFlowJSONAssetRequest struct {
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateFlowJSONAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowJSONAssetRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateFlowJSONAssetRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateFlowJSONAssetRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateFlowJSONAssetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateFlowJSONAssetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateFlowJSONAssetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateFlowJSONAssetRequest) SetCustSpaceId(v string) *UpdateFlowJSONAssetRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetFilePath(v string) *UpdateFlowJSONAssetRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetFlowId(v string) *UpdateFlowJSONAssetRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetOwnerId(v int64) *UpdateFlowJSONAssetRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetResourceOwnerAccount(v string) *UpdateFlowJSONAssetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetResourceOwnerId(v int64) *UpdateFlowJSONAssetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) Validate() error {
	return dara.Validate(s)
}
