// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKMSSecretsForHostShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v int32) *ImportKMSSecretsForHostShrinkRequest
	GetHostId() *int32
	SetInstanceId(v string) *ImportKMSSecretsForHostShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ImportKMSSecretsForHostShrinkRequest
	GetRegionId() *string
	SetSecretsShrink(v string) *ImportKMSSecretsForHostShrinkRequest
	GetSecretsShrink() *string
}

type ImportKMSSecretsForHostShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *int32 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecretsShrink *string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
}

func (s ImportKMSSecretsForHostShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostShrinkRequest) GetHostId() *int32 {
	return s.HostId
}

func (s *ImportKMSSecretsForHostShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportKMSSecretsForHostShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportKMSSecretsForHostShrinkRequest) GetSecretsShrink() *string {
	return s.SecretsShrink
}

func (s *ImportKMSSecretsForHostShrinkRequest) SetHostId(v int32) *ImportKMSSecretsForHostShrinkRequest {
	s.HostId = &v
	return s
}

func (s *ImportKMSSecretsForHostShrinkRequest) SetInstanceId(v string) *ImportKMSSecretsForHostShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportKMSSecretsForHostShrinkRequest) SetRegionId(v string) *ImportKMSSecretsForHostShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ImportKMSSecretsForHostShrinkRequest) SetSecretsShrink(v string) *ImportKMSSecretsForHostShrinkRequest {
	s.SecretsShrink = &v
	return s
}

func (s *ImportKMSSecretsForHostShrinkRequest) Validate() error {
	return dara.Validate(s)
}
