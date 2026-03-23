// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRegionSupportBackupEncryptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckRegionSupportBackupEncryptionResponseBody
	GetRequestId() *string
	SetSupportBackupEncryption(v bool) *CheckRegionSupportBackupEncryptionResponseBody
	GetSupportBackupEncryption() *bool
}

type CheckRegionSupportBackupEncryptionResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportBackupEncryption *bool   `json:"SupportBackupEncryption,omitempty" xml:"SupportBackupEncryption,omitempty"`
}

func (s CheckRegionSupportBackupEncryptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRegionSupportBackupEncryptionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRegionSupportBackupEncryptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRegionSupportBackupEncryptionResponseBody) GetSupportBackupEncryption() *bool {
	return s.SupportBackupEncryption
}

func (s *CheckRegionSupportBackupEncryptionResponseBody) SetRequestId(v string) *CheckRegionSupportBackupEncryptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionResponseBody) SetSupportBackupEncryption(v bool) *CheckRegionSupportBackupEncryptionResponseBody {
	s.SupportBackupEncryption = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionResponseBody) Validate() error {
	return dara.Validate(s)
}
