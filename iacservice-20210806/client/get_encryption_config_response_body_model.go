// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEncryptionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *GetEncryptionConfigResponseBodyConfig) *GetEncryptionConfigResponseBody
	GetConfig() *GetEncryptionConfigResponseBodyConfig
	SetRequestId(v string) *GetEncryptionConfigResponseBody
	GetRequestId() *string
}

type GetEncryptionConfigResponseBody struct {
	// The object key.
	Config *GetEncryptionConfigResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1A662F56-CA76-55F6-869D-7F26293B8E67
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEncryptionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEncryptionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetEncryptionConfigResponseBody) GetConfig() *GetEncryptionConfigResponseBodyConfig {
	return s.Config
}

func (s *GetEncryptionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEncryptionConfigResponseBody) SetConfig(v *GetEncryptionConfigResponseBodyConfig) *GetEncryptionConfigResponseBody {
	s.Config = v
	return s
}

func (s *GetEncryptionConfigResponseBody) SetRequestId(v string) *GetEncryptionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEncryptionConfigResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEncryptionConfigResponseBodyConfig struct {
	// The key alias.
	//
	// example:
	//
	// alias/default
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 561786482014xxxx
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The key ARN.
	//
	// example:
	//
	// acs:kms:ap-southeast-6:561786482014xxxx:key/5f2fe70a-d110-44cc-990f-706eb20fxxxx
	KeyArn *string `json:"keyArn,omitempty" xml:"keyArn,omitempty"`
	// The key ID.
	//
	// example:
	//
	// 5f2fe70a-d110-44cc-990f-706eb20fxxxx
	KeyId *string `json:"keyId,omitempty" xml:"keyId,omitempty"`
	// The key status. Valid values:
	//
	// - Enabled
	//
	// - Disabled
	//
	// - PendingDeletion
	//
	// - PendingImport
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetEncryptionConfigResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetEncryptionConfigResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetEncryptionConfigResponseBodyConfig) GetAlias() *string {
	return s.Alias
}

func (s *GetEncryptionConfigResponseBodyConfig) GetCreator() *string {
	return s.Creator
}

func (s *GetEncryptionConfigResponseBodyConfig) GetKeyArn() *string {
	return s.KeyArn
}

func (s *GetEncryptionConfigResponseBodyConfig) GetKeyId() *string {
	return s.KeyId
}

func (s *GetEncryptionConfigResponseBodyConfig) GetStatus() *string {
	return s.Status
}

func (s *GetEncryptionConfigResponseBodyConfig) SetAlias(v string) *GetEncryptionConfigResponseBodyConfig {
	s.Alias = &v
	return s
}

func (s *GetEncryptionConfigResponseBodyConfig) SetCreator(v string) *GetEncryptionConfigResponseBodyConfig {
	s.Creator = &v
	return s
}

func (s *GetEncryptionConfigResponseBodyConfig) SetKeyArn(v string) *GetEncryptionConfigResponseBodyConfig {
	s.KeyArn = &v
	return s
}

func (s *GetEncryptionConfigResponseBodyConfig) SetKeyId(v string) *GetEncryptionConfigResponseBodyConfig {
	s.KeyId = &v
	return s
}

func (s *GetEncryptionConfigResponseBodyConfig) SetStatus(v string) *GetEncryptionConfigResponseBodyConfig {
	s.Status = &v
	return s
}

func (s *GetEncryptionConfigResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
