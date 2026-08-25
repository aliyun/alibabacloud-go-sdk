// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *DeprovisionAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *DeprovisionAccessConfigurationRequest
	GetDirectoryId() *string
	SetOriginTargetId(v string) *DeprovisionAccessConfigurationRequest
	GetOriginTargetId() *string
	SetTargetId(v string) *DeprovisionAccessConfigurationRequest
	GetTargetId() *string
	SetTargetType(v string) *DeprovisionAccessConfigurationRequest
	GetTargetType() *string
}

type DeprovisionAccessConfigurationRequest struct {
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

func (s DeprovisionAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *DeprovisionAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeprovisionAccessConfigurationRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *DeprovisionAccessConfigurationRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DeprovisionAccessConfigurationRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DeprovisionAccessConfigurationRequest) SetAccessConfigurationId(v string) *DeprovisionAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetDirectoryId(v string) *DeprovisionAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetOriginTargetId(v string) *DeprovisionAccessConfigurationRequest {
	s.OriginTargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetTargetId(v string) *DeprovisionAccessConfigurationRequest {
	s.TargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) SetTargetType(v string) *DeprovisionAccessConfigurationRequest {
	s.TargetType = &v
	return s
}

func (s *DeprovisionAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
