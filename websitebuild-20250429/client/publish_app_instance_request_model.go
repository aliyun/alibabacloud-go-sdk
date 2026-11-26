// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactDigest(v string) *PublishAppInstanceRequest
	GetArtifactDigest() *string
	SetBizId(v string) *PublishAppInstanceRequest
	GetBizId() *string
	SetCandidateId(v string) *PublishAppInstanceRequest
	GetCandidateId() *string
	SetDeployChannel(v string) *PublishAppInstanceRequest
	GetDeployChannel() *string
	SetDescription(v string) *PublishAppInstanceRequest
	GetDescription() *string
	SetExt(v string) *PublishAppInstanceRequest
	GetExt() *string
	SetIdempotencyKey(v string) *PublishAppInstanceRequest
	GetIdempotencyKey() *string
	SetLogicalNumber(v int32) *PublishAppInstanceRequest
	GetLogicalNumber() *int32
	SetPlatformAppid(v string) *PublishAppInstanceRequest
	GetPlatformAppid() *string
	SetPluginDeclarationHash(v string) *PublishAppInstanceRequest
	GetPluginDeclarationHash() *string
	SetPrivacyDraftId(v string) *PublishAppInstanceRequest
	GetPrivacyDraftId() *string
	SetPrivacyRevision(v string) *PublishAppInstanceRequest
	GetPrivacyRevision() *string
	SetPublishNumber(v string) *PublishAppInstanceRequest
	GetPublishNumber() *string
	SetPublishOrderId(v int64) *PublishAppInstanceRequest
	GetPublishOrderId() *int64
	SetUploadSequence(v int64) *PublishAppInstanceRequest
	GetUploadSequence() *int64
	SetWeappAction(v string) *PublishAppInstanceRequest
	GetWeappAction() *string
}

type PublishAppInstanceRequest struct {
	ArtifactDigest *string `json:"ArtifactDigest,omitempty" xml:"ArtifactDigest,omitempty"`
	// The business ID.
	//
	// example:
	//
	// WD20250821161210000001
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The deployment channel.
	//
	// example:
	//
	// PC
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// The publish description.
	//
	// example:
	//
	// CREATE_BY_NLB.nlb-jcvs5sm9l8um84zbfa
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Ext            *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	IdempotencyKey *string `json:"IdempotencyKey,omitempty" xml:"IdempotencyKey,omitempty"`
	// The logical version number.
	//
	// example:
	//
	// 111
	LogicalNumber         *int32  `json:"LogicalNumber,omitempty" xml:"LogicalNumber,omitempty"`
	PlatformAppid         *string `json:"PlatformAppid,omitempty" xml:"PlatformAppid,omitempty"`
	PluginDeclarationHash *string `json:"PluginDeclarationHash,omitempty" xml:"PluginDeclarationHash,omitempty"`
	PrivacyDraftId        *string `json:"PrivacyDraftId,omitempty" xml:"PrivacyDraftId,omitempty"`
	PrivacyRevision       *string `json:"PrivacyRevision,omitempty" xml:"PrivacyRevision,omitempty"`
	// The publish number.
	//
	// example:
	//
	// 123
	PublishNumber  *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	PublishOrderId *int64  `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
	UploadSequence *int64  `json:"UploadSequence,omitempty" xml:"UploadSequence,omitempty"`
	// The action.
	//
	// example:
	//
	// BUILD
	WeappAction *string `json:"WeappAction,omitempty" xml:"WeappAction,omitempty"`
}

func (s PublishAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *PublishAppInstanceRequest) GetArtifactDigest() *string {
	return s.ArtifactDigest
}

func (s *PublishAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *PublishAppInstanceRequest) GetCandidateId() *string {
	return s.CandidateId
}

func (s *PublishAppInstanceRequest) GetDeployChannel() *string {
	return s.DeployChannel
}

func (s *PublishAppInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishAppInstanceRequest) GetExt() *string {
	return s.Ext
}

func (s *PublishAppInstanceRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *PublishAppInstanceRequest) GetLogicalNumber() *int32 {
	return s.LogicalNumber
}

func (s *PublishAppInstanceRequest) GetPlatformAppid() *string {
	return s.PlatformAppid
}

func (s *PublishAppInstanceRequest) GetPluginDeclarationHash() *string {
	return s.PluginDeclarationHash
}

func (s *PublishAppInstanceRequest) GetPrivacyDraftId() *string {
	return s.PrivacyDraftId
}

func (s *PublishAppInstanceRequest) GetPrivacyRevision() *string {
	return s.PrivacyRevision
}

func (s *PublishAppInstanceRequest) GetPublishNumber() *string {
	return s.PublishNumber
}

func (s *PublishAppInstanceRequest) GetPublishOrderId() *int64 {
	return s.PublishOrderId
}

func (s *PublishAppInstanceRequest) GetUploadSequence() *int64 {
	return s.UploadSequence
}

func (s *PublishAppInstanceRequest) GetWeappAction() *string {
	return s.WeappAction
}

func (s *PublishAppInstanceRequest) SetArtifactDigest(v string) *PublishAppInstanceRequest {
	s.ArtifactDigest = &v
	return s
}

func (s *PublishAppInstanceRequest) SetBizId(v string) *PublishAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *PublishAppInstanceRequest) SetCandidateId(v string) *PublishAppInstanceRequest {
	s.CandidateId = &v
	return s
}

func (s *PublishAppInstanceRequest) SetDeployChannel(v string) *PublishAppInstanceRequest {
	s.DeployChannel = &v
	return s
}

func (s *PublishAppInstanceRequest) SetDescription(v string) *PublishAppInstanceRequest {
	s.Description = &v
	return s
}

func (s *PublishAppInstanceRequest) SetExt(v string) *PublishAppInstanceRequest {
	s.Ext = &v
	return s
}

func (s *PublishAppInstanceRequest) SetIdempotencyKey(v string) *PublishAppInstanceRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *PublishAppInstanceRequest) SetLogicalNumber(v int32) *PublishAppInstanceRequest {
	s.LogicalNumber = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPlatformAppid(v string) *PublishAppInstanceRequest {
	s.PlatformAppid = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPluginDeclarationHash(v string) *PublishAppInstanceRequest {
	s.PluginDeclarationHash = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPrivacyDraftId(v string) *PublishAppInstanceRequest {
	s.PrivacyDraftId = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPrivacyRevision(v string) *PublishAppInstanceRequest {
	s.PrivacyRevision = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPublishNumber(v string) *PublishAppInstanceRequest {
	s.PublishNumber = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPublishOrderId(v int64) *PublishAppInstanceRequest {
	s.PublishOrderId = &v
	return s
}

func (s *PublishAppInstanceRequest) SetUploadSequence(v int64) *PublishAppInstanceRequest {
	s.UploadSequence = &v
	return s
}

func (s *PublishAppInstanceRequest) SetWeappAction(v string) *PublishAppInstanceRequest {
	s.WeappAction = &v
	return s
}

func (s *PublishAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
