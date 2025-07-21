// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDKMSInstanceId(v string) *CreateKeyRequest
	GetDKMSInstanceId() *string
	SetDescription(v string) *CreateKeyRequest
	GetDescription() *string
	SetEnableAutomaticRotation(v bool) *CreateKeyRequest
	GetEnableAutomaticRotation() *bool
	SetKeySpec(v string) *CreateKeyRequest
	GetKeySpec() *string
	SetKeyStorageMechanism(v string) *CreateKeyRequest
	GetKeyStorageMechanism() *string
	SetKeyUsage(v string) *CreateKeyRequest
	GetKeyUsage() *string
	SetOrigin(v string) *CreateKeyRequest
	GetOrigin() *string
	SetPolicy(v string) *CreateKeyRequest
	GetPolicy() *string
	SetProtectionLevel(v string) *CreateKeyRequest
	GetProtectionLevel() *string
	SetRotationInterval(v string) *CreateKeyRequest
	GetRotationInterval() *string
	SetTags(v string) *CreateKeyRequest
	GetTags() *string
}

type CreateKeyRequest struct {
	// The ID of the KMS instance.
	//
	// > You must specify this parameter if you need to create a key for a KMS instance. If you need to create a default key of the CMK type, you do not need to specify this parameter.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the key.
	//
	// The description can be 0 to 8,192 characters in length.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable automatic key rotation. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// This parameter is valid only when the key belongs to an instance type that supports automatic rotation. For more information, see [Key rotation](https://help.aliyun.com/document_detail/2358146.html).
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The key specification. The valid values vary based on the KMS instance type. For more information, see [Overview](https://help.aliyun.com/document_detail/480159.html).
	//
	// > If you do not specify a value for this parameter, the default key specification is Aliyun_AES_256.
	//
	// example:
	//
	// Aliyun_AES_256
	KeySpec             *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	KeyStorageMechanism *string `json:"KeyStorageMechanism,omitempty" xml:"KeyStorageMechanism,omitempty"`
	// The usage of the key. Valid values:
	//
	// - ENCRYPT/DECRYPT
	//
	// - SIGN/VERIFY
	//
	// If the key supports signing and verification, the default value is SIGN/VERIFY. If the key does not support signing and verification, the default value is ENCRYPT/DECRYPT.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The key material origin. Valid values:
	//
	// - Aliyun_KMS (default): KMS generates key material.
	//
	// - EXTERNAL: You import key material.
	//
	//
	// > - The value of this parameter is case-sensitive.
	//
	// > - Default keys of the customer master key (CMK) type support Aliyun_KMS and EXTERNAL. Keys in instances of the software key management type support only Aliyun_KMS. Keys in instances of the hardware key management type support Aliyun_KMS and EXTERNAL.
	//
	// > - If you set Origin to EXTERNAL, you must import key material. For more information, see [Import key material into a symmetric key](https://help.aliyun.com/document_detail/607841.html) or [Import key material into an asymmetric key](https://help.aliyun.com/document_detail/608827.html).
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// You do not need to specify this parameter. KMS sets a protection level for your key.
	//
	// The protection level of the key. Valid values:
	//
	// - SOFTWARE
	//
	// - HSM
	//
	//
	// > - If DKMSInstanceId is specified, this parameter does not take effect. If your instance is an instance of the software key management type, set the value to SOFTWARE. If your instance is an instance of the hardware key management type, set the value to HSM.
	//
	// > - If you do not specify DKMSInstanceId, we recommend that you do not specify this parameter. KMS sets a protection level for your key. If managed hardware security modules (HSMs) exist in the region of your KMS instance, set the value to HSM. If managed HSMs do not exist in the region of your KMS instance, set the value to SOFTWARE. For more information, see Managed HSM overview.
	//
	// example:
	//
	// SOFTWARE
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The period of automatic key rotation. Format: integer[unit]. Unit: d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s represent a seven-day interval.
	//
	// - For a default key, set the value to 365 days.
	//
	// - For a software-protected key, set a value that ranges from 7 to 365 days.
	//
	// - A hardware-protected key does not support automatic rotation.
	//
	// > If EnableAutomaticRotation is set to true, this parameter is required.
	//
	// example:
	//
	// 365d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The tag that is added to the key. A tag consists of a key-value pair.
	//
	// You can enter up to 20 tags. Enter multiple tags in the [{"TagKey":"key1","TagValue":"value1"},{"TagKey":"key2","TagValue":"value2"},..] format.
	//
	// Each tag key or tag value can be up to 128 characters in length and can contain letters, digits, forward slashes (/), backslashes (\\), underscores (_), hyphens (-), periods (.), plus signs (+), equal signs (=), colons (:), and at signs (@).
	//
	// > The tag key cannot start with aliyun or acs:.
	//
	// example:
	//
	// [{"TagKey":"disk-encryption","TagValue":"true"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyRequest) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKeyRequest) GetEnableAutomaticRotation() *bool {
	return s.EnableAutomaticRotation
}

func (s *CreateKeyRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *CreateKeyRequest) GetKeyStorageMechanism() *string {
	return s.KeyStorageMechanism
}

func (s *CreateKeyRequest) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *CreateKeyRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateKeyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateKeyRequest) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *CreateKeyRequest) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *CreateKeyRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateKeyRequest) SetDKMSInstanceId(v string) *CreateKeyRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateKeyRequest) SetDescription(v string) *CreateKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateKeyRequest) SetEnableAutomaticRotation(v bool) *CreateKeyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateKeyRequest) SetKeySpec(v string) *CreateKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateKeyRequest) SetKeyStorageMechanism(v string) *CreateKeyRequest {
	s.KeyStorageMechanism = &v
	return s
}

func (s *CreateKeyRequest) SetKeyUsage(v string) *CreateKeyRequest {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyRequest) SetOrigin(v string) *CreateKeyRequest {
	s.Origin = &v
	return s
}

func (s *CreateKeyRequest) SetPolicy(v string) *CreateKeyRequest {
	s.Policy = &v
	return s
}

func (s *CreateKeyRequest) SetProtectionLevel(v string) *CreateKeyRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyRequest) SetRotationInterval(v string) *CreateKeyRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyRequest) SetTags(v string) *CreateKeyRequest {
	s.Tags = &v
	return s
}

func (s *CreateKeyRequest) Validate() error {
	return dara.Validate(s)
}
