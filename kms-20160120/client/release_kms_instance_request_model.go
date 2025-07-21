// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseKmsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDeleteWithoutBackup(v string) *ReleaseKmsInstanceRequest
	GetForceDeleteWithoutBackup() *string
	SetKmsInstanceId(v string) *ReleaseKmsInstanceRequest
	GetKmsInstanceId() *string
}

type ReleaseKmsInstanceRequest struct {
	// example:
	//
	// false
	ForceDeleteWithoutBackup *string `json:"ForceDeleteWithoutBackup,omitempty" xml:"ForceDeleteWithoutBackup,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kst-hzz6****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s ReleaseKmsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseKmsInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseKmsInstanceRequest) GetForceDeleteWithoutBackup() *string {
	return s.ForceDeleteWithoutBackup
}

func (s *ReleaseKmsInstanceRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *ReleaseKmsInstanceRequest) SetForceDeleteWithoutBackup(v string) *ReleaseKmsInstanceRequest {
	s.ForceDeleteWithoutBackup = &v
	return s
}

func (s *ReleaseKmsInstanceRequest) SetKmsInstanceId(v string) *ReleaseKmsInstanceRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *ReleaseKmsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
