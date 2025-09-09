// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOperationTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApproveComment(v string) *CreateOperationTicketRequest
	GetApproveComment() *string
	SetAssetAccountName(v string) *CreateOperationTicketRequest
	GetAssetAccountName() *string
	SetAssetId(v string) *CreateOperationTicketRequest
	GetAssetId() *string
	SetEffectEndTime(v int64) *CreateOperationTicketRequest
	GetEffectEndTime() *int64
	SetEffectStartTime(v int64) *CreateOperationTicketRequest
	GetEffectStartTime() *int64
	SetInstanceId(v string) *CreateOperationTicketRequest
	GetInstanceId() *string
	SetIsOneTimeEffect(v bool) *CreateOperationTicketRequest
	GetIsOneTimeEffect() *bool
	SetProtocolName(v string) *CreateOperationTicketRequest
	GetProtocolName() *string
	SetRegionId(v string) *CreateOperationTicketRequest
	GetRegionId() *string
}

type CreateOperationTicketRequest struct {
	// This parameter is required.
	ApproveComment *string `json:"ApproveComment,omitempty" xml:"ApproveComment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root
	AssetAccountName *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// example:
	//
	// 1679393152
	EffectEndTime *int64 `json:"EffectEndTime,omitempty" xml:"EffectEndTime,omitempty"`
	// example:
	//
	// 1685600242
	EffectStartTime *int64 `json:"EffectStartTime,omitempty" xml:"EffectStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IsOneTimeEffect *bool `json:"IsOneTimeEffect,omitempty" xml:"IsOneTimeEffect,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOperationTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateOperationTicketRequest) GetApproveComment() *string {
	return s.ApproveComment
}

func (s *CreateOperationTicketRequest) GetAssetAccountName() *string {
	return s.AssetAccountName
}

func (s *CreateOperationTicketRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *CreateOperationTicketRequest) GetEffectEndTime() *int64 {
	return s.EffectEndTime
}

func (s *CreateOperationTicketRequest) GetEffectStartTime() *int64 {
	return s.EffectStartTime
}

func (s *CreateOperationTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOperationTicketRequest) GetIsOneTimeEffect() *bool {
	return s.IsOneTimeEffect
}

func (s *CreateOperationTicketRequest) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *CreateOperationTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOperationTicketRequest) SetApproveComment(v string) *CreateOperationTicketRequest {
	s.ApproveComment = &v
	return s
}

func (s *CreateOperationTicketRequest) SetAssetAccountName(v string) *CreateOperationTicketRequest {
	s.AssetAccountName = &v
	return s
}

func (s *CreateOperationTicketRequest) SetAssetId(v string) *CreateOperationTicketRequest {
	s.AssetId = &v
	return s
}

func (s *CreateOperationTicketRequest) SetEffectEndTime(v int64) *CreateOperationTicketRequest {
	s.EffectEndTime = &v
	return s
}

func (s *CreateOperationTicketRequest) SetEffectStartTime(v int64) *CreateOperationTicketRequest {
	s.EffectStartTime = &v
	return s
}

func (s *CreateOperationTicketRequest) SetInstanceId(v string) *CreateOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOperationTicketRequest) SetIsOneTimeEffect(v bool) *CreateOperationTicketRequest {
	s.IsOneTimeEffect = &v
	return s
}

func (s *CreateOperationTicketRequest) SetProtocolName(v string) *CreateOperationTicketRequest {
	s.ProtocolName = &v
	return s
}

func (s *CreateOperationTicketRequest) SetRegionId(v string) *CreateOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
