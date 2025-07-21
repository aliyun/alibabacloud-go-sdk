// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyMetadata(v *CreateKeyResponseBodyKeyMetadata) *CreateKeyResponseBody
	GetKeyMetadata() *CreateKeyResponseBodyKeyMetadata
	SetRequestId(v string) *CreateKeyResponseBody
	GetRequestId() *string
}

type CreateKeyResponseBody struct {
	// The metadata of the key.
	KeyMetadata *CreateKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBody) GetKeyMetadata() *CreateKeyResponseBodyKeyMetadata {
	return s.KeyMetadata
}

func (s *CreateKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKeyResponseBody) SetKeyMetadata(v *CreateKeyResponseBodyKeyMetadata) *CreateKeyResponseBody {
	s.KeyMetadata = v
	return s
}

func (s *CreateKeyResponseBody) SetRequestId(v string) *CreateKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateKeyResponseBodyKeyMetadata struct {
	// The Alibaba Cloud Resource Name (ARN) of the key.
	//
	// example:
	//
	// acs:kms:cn-qingdao:154035569884****:key/key-hzz62f1cb66fa42qo****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The status of automatic key rotation. Valid values:
	//
	// - Enabled
	//
	// - Disabled
	//
	// - Suspended
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The date and time (UTC) when the key was created.
	//
	// example:
	//
	// 2023-03-25T10:00:00Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The user who created the key.
	//
	// example:
	//
	// 154035569884****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time when the key is scheduled for deletion. For more information, see ScheduleKeyDeletion.
	//
	// This parameter is returned only when the value of KeyState is PendingDeletion.
	//
	// example:
	//
	// 2025-03-25T10:00:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The globally unique ID of the key.
	//
	// example:
	//
	// key-hzz62f1cb66fa42qo****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The specification of the key.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the key.
	//
	// For more information, see [Impacts of key status on API operations](https://help.aliyun.com/document_detail/44211.html).
	//
	// example:
	//
	// Enabled
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the key.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC.
	//
	// For a new key, this parameter value is the time when the initial version of the key was generated.
	//
	// example:
	//
	// 2023-03-25T10:00:00Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC.
	//
	// If this parameter value is empty, the key material does not expire.
	//
	// example:
	//
	// 2025-03-25T10:00:00Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the key is next rotated.
	//
	// This value is returned only when the value of AutomaticRotation is Enabled or Suspended.
	//
	// example:
	//
	// 2024-03-25T10:00:00Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The key material origin.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The current primary version identifier of the key.
	//
	// example:
	//
	// 7ce1d081-06cb-42e6-aab6-5c5de030****
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the key.
	//
	// example:
	//
	// SOFTWARE
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The interval for automatic key rotation. Unit: seconds. The format is an integer value followed by the character s. For example, if the rotation period is seven days, this parameter is set to 604800s.
	//
	// This value is returned only when the value of AutomaticRotation is Enabled or Suspended.
	//
	// example:
	//
	// 31536000s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s CreateKeyResponseBodyKeyMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseBodyKeyMetadata) GetArn() *string {
	return s.Arn
}

func (s *CreateKeyResponseBodyKeyMetadata) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *CreateKeyResponseBodyKeyMetadata) GetCreationDate() *string {
	return s.CreationDate
}

func (s *CreateKeyResponseBodyKeyMetadata) GetCreator() *string {
	return s.Creator
}

func (s *CreateKeyResponseBodyKeyMetadata) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateKeyResponseBodyKeyMetadata) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *CreateKeyResponseBodyKeyMetadata) GetDescription() *string {
	return s.Description
}

func (s *CreateKeyResponseBodyKeyMetadata) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateKeyResponseBodyKeyMetadata) GetKeySpec() *string {
	return s.KeySpec
}

func (s *CreateKeyResponseBodyKeyMetadata) GetKeyState() *string {
	return s.KeyState
}

func (s *CreateKeyResponseBodyKeyMetadata) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *CreateKeyResponseBodyKeyMetadata) GetLastRotationDate() *string {
	return s.LastRotationDate
}

func (s *CreateKeyResponseBodyKeyMetadata) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *CreateKeyResponseBodyKeyMetadata) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *CreateKeyResponseBodyKeyMetadata) GetOrigin() *string {
	return s.Origin
}

func (s *CreateKeyResponseBodyKeyMetadata) GetPrimaryKeyVersion() *string {
	return s.PrimaryKeyVersion
}

func (s *CreateKeyResponseBodyKeyMetadata) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *CreateKeyResponseBodyKeyMetadata) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *CreateKeyResponseBodyKeyMetadata) SetArn(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *CreateKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetCreator(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDKMSInstanceId(v string) *CreateKeyResponseBodyKeyMetadata {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetDescription(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyId(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeySpec(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyState(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *CreateKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *CreateKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *CreateKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetOrigin(v string) *CreateKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *CreateKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *CreateKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *CreateKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyResponseBodyKeyMetadata) Validate() error {
	return dara.Validate(s)
}
