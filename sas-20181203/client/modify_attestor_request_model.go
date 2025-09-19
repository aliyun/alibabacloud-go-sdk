// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAttestorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *ModifyAttestorRequest
	GetKeyId() *string
	SetKeyRegionId(v string) *ModifyAttestorRequest
	GetKeyRegionId() *string
	SetKeyVersionId(v string) *ModifyAttestorRequest
	GetKeyVersionId() *string
	SetName(v string) *ModifyAttestorRequest
	GetName() *string
	SetRemark(v string) *ModifyAttestorRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *ModifyAttestorRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifyAttestorRequest
	GetSourceIp() *string
}

type ModifyAttestorRequest struct {
	// The ID of the KMS key.
	//
	// example:
	//
	// key-********
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The region ID of the Key Management Service (KMS) key.
	//
	// example:
	//
	// cn-hangzhou
	KeyRegionId *string `json:"KeyRegionId,omitempty" xml:"KeyRegionId,omitempty"`
	// The version ID of the KMS key.
	//
	// example:
	//
	// key-****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The name of the witness.
	//
	// example:
	//
	// attestor-auto-ad5316
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description.
	//
	// example:
	//
	// remark
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyAttestorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAttestorRequest) GoString() string {
	return s.String()
}

func (s *ModifyAttestorRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ModifyAttestorRequest) GetKeyRegionId() *string {
	return s.KeyRegionId
}

func (s *ModifyAttestorRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *ModifyAttestorRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAttestorRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyAttestorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAttestorRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAttestorRequest) SetKeyId(v string) *ModifyAttestorRequest {
	s.KeyId = &v
	return s
}

func (s *ModifyAttestorRequest) SetKeyRegionId(v string) *ModifyAttestorRequest {
	s.KeyRegionId = &v
	return s
}

func (s *ModifyAttestorRequest) SetKeyVersionId(v string) *ModifyAttestorRequest {
	s.KeyVersionId = &v
	return s
}

func (s *ModifyAttestorRequest) SetName(v string) *ModifyAttestorRequest {
	s.Name = &v
	return s
}

func (s *ModifyAttestorRequest) SetRemark(v string) *ModifyAttestorRequest {
	s.Remark = &v
	return s
}

func (s *ModifyAttestorRequest) SetResourceOwnerId(v int64) *ModifyAttestorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAttestorRequest) SetSourceIp(v string) *ModifyAttestorRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAttestorRequest) Validate() error {
	return dara.Validate(s)
}
