// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyMetadata(v *DescribeKeyResponseBodyKeyMetadata) *DescribeKeyResponseBody
	GetKeyMetadata() *DescribeKeyResponseBodyKeyMetadata
	SetRequestId(v string) *DescribeKeyResponseBody
	GetRequestId() *string
}

type DescribeKeyResponseBody struct {
	// The metadata of the CMK.
	KeyMetadata *DescribeKeyResponseBodyKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f1fdfa9d-bd49-418b-942f-8f3e3ec00a4f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBody) GetKeyMetadata() *DescribeKeyResponseBodyKeyMetadata {
	return s.KeyMetadata
}

func (s *DescribeKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKeyResponseBody) SetKeyMetadata(v *DescribeKeyResponseBodyKeyMetadata) *DescribeKeyResponseBody {
	s.KeyMetadata = v
	return s
}

func (s *DescribeKeyResponseBody) SetRequestId(v string) *DescribeKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyResponseBody) Validate() error {
	if s.KeyMetadata != nil {
		if err := s.KeyMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKeyResponseBodyKeyMetadata struct {
	// The Alibaba Cloud Resource Name (ARN) of the CMK.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:key/05754286-3ba2-4fa6-8d41-4323aca6****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// Indicates whether automatic key rotation is enabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// 	- Suspended
	//
	// For more information, see [Automatic key rotation](https://help.aliyun.com/document_detail/134270.html).
	//
	// >  Only symmetric CMKs support automatic key rotation.
	//
	// example:
	//
	// Disabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the CMK was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-20T06:34:21Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The Alibaba Cloud account that is used to create the CMK.
	//
	// example:
	//
	// 154035569884****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The time at which the CMK is scheduled for deletion. The time is displayed in UTC.
	//
	// For more information, see [ScheduleKeyDeletion](https://help.aliyun.com/document_detail/44196.html).
	//
	// >  This parameter is returned only when the value of the KeyState parameter is PendingDeletion.
	//
	// example:
	//
	// 2021-05-26T18:22:03Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of deletion protection.
	//
	// example:
	//
	// The CMK is being used by XXX. Deletion protection is set.
	DeletionProtectionDescription *string `json:"DeletionProtectionDescription,omitempty" xml:"DeletionProtectionDescription,omitempty"`
	// The description of the CMK.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 05754286-3ba2-4fa6-8d41-4323aca6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the CMK.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The status of the CMK.
	//
	// For more information, see [Impact of CMK status on API operations](https://help.aliyun.com/document_detail/44211.html).
	//
	// example:
	//
	// Enabled
	KeyState *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	// The usage of the CMK.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the last rotation was performed. The time is displayed in UTC. For a new CMK, the value of this parameter is the time when the initial version of the CMK was generated.
	//
	// example:
	//
	// 2021-05-20T06:34:21Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the key material expires. The time is displayed in UTC. If this parameter value is empty, the key material does not expire.
	//
	// example:
	//
	// 2021-07-06T18:22:03Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	//
	// example:
	//
	// 2021-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The source of the key material for the CMK.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the current primary key version for the symmetric CMK.
	//
	// example:
	//
	// 515e0b0a-624f-45ab-92b5-54f9b551****
	PrimaryKeyVersion *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty"`
	// The protection level of the CMK.
	//
	// example:
	//
	// HSM
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The interval for automatic key rotation.
	//
	// Unit: seconds.
	//
	// For example, if the value is 604800s, automatic key rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned only when the value of the AutomaticRotation parameter is Enabled or Suspended.
	//
	// example:
	//
	// 31536000s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s DescribeKeyResponseBodyKeyMetadata) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyResponseBodyKeyMetadata) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetArn() *string {
	return s.Arn
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetCreationDate() *string {
	return s.CreationDate
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetCreator() *string {
	return s.Creator
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetDeletionProtectionDescription() *string {
	return s.DeletionProtectionDescription
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetDescription() *string {
	return s.Description
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetKeySpec() *string {
	return s.KeySpec
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetKeyState() *string {
	return s.KeyState
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetLastRotationDate() *string {
	return s.LastRotationDate
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetPrimaryKeyVersion() *string {
	return s.PrimaryKeyVersion
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *DescribeKeyResponseBodyKeyMetadata) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetArn(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Arn = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetAutomaticRotation(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetCreator(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Creator = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDKMSInstanceId(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DKMSInstanceId = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeleteDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeletionProtection(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDeletionProtectionDescription(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.DeletionProtectionDescription = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetDescription(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Description = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyId(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeySpec(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeySpec = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyState(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetKeyUsage(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetLastRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetMaterialExpireTime(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetNextRotationDate(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetOrigin(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.Origin = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetPrimaryKeyVersion(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetProtectionLevel(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) SetRotationInterval(v string) *DescribeKeyResponseBodyKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *DescribeKeyResponseBodyKeyMetadata) Validate() error {
	return dara.Validate(s)
}
