// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttestorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *CreateAttestorRequest
	GetKeyId() *string
	SetKeyRegionId(v string) *CreateAttestorRequest
	GetKeyRegionId() *string
	SetKeyVersionId(v string) *CreateAttestorRequest
	GetKeyVersionId() *string
	SetName(v string) *CreateAttestorRequest
	GetName() *string
	SetRemark(v string) *CreateAttestorRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *CreateAttestorRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *CreateAttestorRequest
	GetSourceIp() *string
}

type CreateAttestorRequest struct {
	// The ID of the KMS key.
	//
	// example:
	//
	// 2e81355b-f8e7-4090-8082-a8f8124a****
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
	// 8d7c9c91-57ce-4cf4-a959-1e700e13****
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
	// attestor
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateAttestorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAttestorRequest) GoString() string {
	return s.String()
}

func (s *CreateAttestorRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateAttestorRequest) GetKeyRegionId() *string {
	return s.KeyRegionId
}

func (s *CreateAttestorRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *CreateAttestorRequest) GetName() *string {
	return s.Name
}

func (s *CreateAttestorRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateAttestorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAttestorRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateAttestorRequest) SetKeyId(v string) *CreateAttestorRequest {
	s.KeyId = &v
	return s
}

func (s *CreateAttestorRequest) SetKeyRegionId(v string) *CreateAttestorRequest {
	s.KeyRegionId = &v
	return s
}

func (s *CreateAttestorRequest) SetKeyVersionId(v string) *CreateAttestorRequest {
	s.KeyVersionId = &v
	return s
}

func (s *CreateAttestorRequest) SetName(v string) *CreateAttestorRequest {
	s.Name = &v
	return s
}

func (s *CreateAttestorRequest) SetRemark(v string) *CreateAttestorRequest {
	s.Remark = &v
	return s
}

func (s *CreateAttestorRequest) SetResourceOwnerId(v int64) *CreateAttestorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAttestorRequest) SetSourceIp(v string) *CreateAttestorRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAttestorRequest) Validate() error {
	return dara.Validate(s)
}
