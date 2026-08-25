// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *ProvisionAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *ProvisionAccessConfigurationRequest
	GetDirectoryId() *string
	SetOriginTargetId(v string) *ProvisionAccessConfigurationRequest
	GetOriginTargetId() *string
	SetTargetId(v string) *ProvisionAccessConfigurationRequest
	GetTargetId() *string
	SetTargetType(v string) *ProvisionAccessConfigurationRequest
	GetTargetType() *string
}

type ProvisionAccessConfigurationRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the task object.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ProvisionAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ProvisionAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ProvisionAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ProvisionAccessConfigurationRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ProvisionAccessConfigurationRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ProvisionAccessConfigurationRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ProvisionAccessConfigurationRequest) SetAccessConfigurationId(v string) *ProvisionAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetDirectoryId(v string) *ProvisionAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetOriginTargetId(v string) *ProvisionAccessConfigurationRequest {
	s.OriginTargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetTargetId(v string) *ProvisionAccessConfigurationRequest {
	s.TargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) SetTargetType(v string) *ProvisionAccessConfigurationRequest {
	s.TargetType = &v
	return s
}

func (s *ProvisionAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
