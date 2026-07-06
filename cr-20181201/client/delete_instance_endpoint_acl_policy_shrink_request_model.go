// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceEndpointAclPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest
	GetEndpointType() *string
	SetEntriesShrink(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest
	GetEntriesShrink() *string
	SetEntry(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest
	GetEntry() *string
	SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest
	GetInstanceId() *string
	SetModuleName(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest
	GetModuleName() *string
}

type DeleteInstanceEndpointAclPolicyShrinkRequest struct {
	// The endpoint type. Only Internet is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType  *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EntriesShrink *string `json:"Entries,omitempty" xml:"Entries,omitempty"`
	// Deprecated
	//
	// The IP CIDR block.
	//
	// example:
	//
	// 127.0.0.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module for which the access policy is set. Valid values:
	//
	// - `Registry`: access to the image repository
	//
	// - `Chart`: access to Helm Chart
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) GetEntriesShrink() *string {
	return s.EntriesShrink
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) GetEntry() *string {
	return s.Entry
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest {
	s.EndpointType = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) SetEntriesShrink(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest {
	s.EntriesShrink = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) SetEntry(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest {
	s.Entry = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) SetModuleName(v string) *DeleteInstanceEndpointAclPolicyShrinkRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
